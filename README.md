# go-playground

[![Go Report Card](https://goreportcard.com/badge/github.com/hoanhan101/go-playground)](https://goreportcard.com/report/github.com/hoanhan101/go-playground)

**go-playground** is a collection of Go's programs and notes for learning
Go, which I mostly use for reference.

## Topic
- Basic: scope, callback, function expression, defer, slice, pointer
- Concurrency
- Networking: http, tcp, mux, session, cookie
- Error Handling

## Note

### How to setup `GOPATH`

There are 2 essential environment variables to setup before running any Go program:
- `GOPATH`
- `GOROOT`

`GOPATH` is the path to our workspace. It is required and has no default. 
I'm setting it up at `$HOME/go` because it works for me.

`GOROOT` is that path to where Go standard library is located on our local filesystem.
Is set automatically setup when we install Go.

Additionally, `GOBIN` is the path to where our Go binaries are installed from running `go install`.

All of the environment variables configurations can be set in `.bashrc` or `.zshrc`. For example:
```
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

Execute the command `go env` to see all the environment variable.

### How to structure project

#### Exported/unexported, aka visible/not visible
In go we don't say public or or private, we say exported or unexported.
Depend on capialization:
- capitalize: exported, visible outside the package
- lowercase: unxported, not visible outside the package

#### go commands
Assume that our folder directory looks like this:
```
package
├── main
│  └── main.go
├── utils
│  └── utilOne.go
│  └── utilTwo.go
```

If we run these commands in our executable, `main`:
- `go run main.go` will run the file the file
- `go build` will make an executable file and put that inside
- `go clean` will clean that executable file
- `go install` will make an executable file inside the workspace `bin` folder and archive file in `pkg`

If we run these commands in our packages, `utils`:
- they cannot be run because they are non-main packages
- they cannot be built because there are no executable

### `go fmt`

`go fmt ./...`: Go fmt all the files in this directory and anything below it.
