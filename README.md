# go-playground

### Setup

### Environment variables
There are 2 essential environment variables to setup before running any Go program:
- `GOPATH`
- `GOROOT`

`GOPATH` is the path to your workspace. It is required and has no default. 
I'm setting it up at `$HOME/go` because it works for me.

`GOROOT` is that path to where Go standard library is located on your local filesystem.
Is set automatically setup when you install Go.

`GOBIN` is the path to where your Go binaries are installed from running `go install`.

All of the environment variables configurations can be set in `.bashrc`. For example:
```
export GOPATH=$HOME/go
export GOROOT=/usr/local/opt/go/libexec
export GOBIN=$HOME/go/bin
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOBIN
```

Execute the command `go env` to see all the environment variable.
