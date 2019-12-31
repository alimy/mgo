package grace

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// SignalFun os.Signal handler
type SignalFun func(os.Signal) (isContinue bool)

var (
	// SignalDefault specify exception signals for application.
	// Program will want exit when catch SIGINT/SIGTERM and SIGHUP is ignored
	// since that is usually used for triggering a reload of configuration which
	// isn't  supported but shouldn't kill the process either.
	SignalDefault = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGHUP}

	// quit channel is closed to cleanup exit listeners.
	quit = make(chan struct{})
	// stubbed out for testing
	osExit = os.Exit

	wg sync.WaitGroup
)

// Listen registers an signal handler which is called on
// os.Signal catch or when Exit/Fatal/Fatalf is called.
// if signals is empty will use SignalExceptions as signals.
func Listen(fn SignalFun, signals ...os.Signal) {
	if fn == nil {
		log.Print("func of SignalFun is null")
		return
	}
	if len(signals) == 0 { // use SignalDefault when signals is empty
		signals = SignalDefault
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		sigChan := make(chan os.Signal, 1)
		for {
			signal.Notify(sigChan, signals...)
			select {
			case sig := <-sigChan:
				// specify handler SIGQUIT
				if fn(sig) && sig != syscall.SIGQUIT {
					continue
				}
			case <-quit:
				// just to quit when Exit/Fatal/Fatalf is called whether handler's
				// return true or false.
				fn(syscall.SIGQUIT)
			}
			return
		}
	}()
}

// Exit terminates the application via os.Exit but
// waits for all exit handlers to complete before
// calling os.Exit.
func Exit(code int) {
	quit <- struct{}{}
	wg.Wait()
	osExit(code)
}

// Fatal is a replacement for log.Fatal which will trigger
// the closure of all registered exit handlers and waits
// for their completion and then call os.Exit(1).
func Fatal(v ...interface{}) {
	log.Print(v...)
	Exit(1)
}

// Fatalf is a replacement for log.Fatalf and behaves like Fatal.
func Fatalf(format string, v ...interface{}) {
	log.Printf(format, v...)
	Exit(1)
}

// Wait waits for all exit handlers to complete.
func Wait() {
	wg.Wait()
}
