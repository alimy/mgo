package utils

import "testing"

func TestBytesToString(t *testing.T) {
	testStr := "Hello Gopher"
	originBytes := []byte(testStr)
	convertedStr := BytesToString(originBytes)
	if convertedStr != testStr {
		t.Errorf("want result: %s but get %s", testStr, convertedStr)
	}
}

func TestStringToBytes(t *testing.T) {
	testStr := "Hello Gopher"
	wantedBytes := []byte(testStr)
	convertedBytes := StringToBytes(testStr)
	if len(convertedBytes) == len(wantedBytes) {
		for i, v := range convertedBytes {
			if wantedBytes[i] != v {
				goto ErrEdge
			}
		}
		return
	}
ErrEdge:
	t.Errorf("want result: %v but get %v", wantedBytes, convertedBytes)
}
