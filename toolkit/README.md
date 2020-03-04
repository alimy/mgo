## mgo
mgo's help toolkit

### Usage
```bash
% git clone https://github.com/alimy/mgo.git $GOPATH/src/github.com/alimy/mgo
% cd $GOPATH/src/github.com/alimy/mgo/toolkit
% make build
% ./mgo new -d example -p github.com/alimy/mgo/toolkit/example
% tree example
example
├── Makefile
├── README.md
├── assets
│   ├── README.md
│   ├── certs
│   └── config
├── cmd
│   ├── README.md
│   ├── core
│   ├── zimctl
│   ├── zimctl.go
│   ├── zimlet
│   └── zimlet.go
├── core
│   ├── core.go
│   ├── models
│   ├── proto
│   └── servant
├── docs
│   └── README.md
├── go.mod
├── hack
│   ├── README.md
│   ├── goprotoc.sh
│   └── systemd
├── internal
│   ├── assets
│   ├── config
│   ├── insecure
│   ├── logus
│   └── rpcx
├── servants
│   ├── agent
│   ├── business
│   ├── dataware
│   ├── servants.go
│   └── share
└── version
    └── version.go

% cd example
% make build
% ls
Makefile  assets    core      go.mod    hack      servants  zimctl
README.md cmd       docs      go.sum    internal  version   zimlet
```

### Release
```bash
%  hub release create -m "toolkit/{tag eg:v0.1.0} release" toolkit/{tag eg:v0.1.0}
```
