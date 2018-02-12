# go-playground

### Setup

### Environment variables
There are 2 essential environment variables to setup before running any Go program:
- `GOPATH`
- `GOROOT`

`GOPATH` is the path to our workspace. It is required and has no default. 
I'm setting it up at `$HOME/go` because it works for me.

`GOROOT` is that path to where Go standard library is located on our local filesystem.
Is set automatically setup when we install Go.

Additionally, `GOBIN` is the path to where our Go binaries are installed from running `go install`.

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

### Packages
#### Exported/unexported, aka visible/not visible
In go we don't say public or or private, we say exported or unexported.
Depend on capialization:
- capitalize: exported, visible outside the package
- lowercase: unxported, not visible outside the package

#### go commands
Assume that our folder directory looks like this:
```
package
    main
        main.go
    utils
        utilOne.go
        utilTwo.go
```

If we run these commands in our executable, `main`:
- `go run main.go` will run the file the file
- `go build` will make an executable file and put that inside
- `go clean` will clean that executable file
- `go install` will make an executable file inside the workspace `bin` folder and archive file in `pkg`

If we run these commands in our packages, `utils`:
- they cannot be run because they are non-main packages
- they cannot be built because there are no executable

### Format

`go fmt ./...`: In this folder, go to every file and format it
