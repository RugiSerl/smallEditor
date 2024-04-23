# smallEditor

Small project to code an golang ide in golang

## Compiling
### Dependencies
- Go 1.22.1 or later
- C compiler (I recommand [TDM-GCC](https://jmeubank.github.io/tdm-gcc/))
### Building from source
For a debug build:

Get into the repository folder and type:
> go run main.go

For a release build:
> go build -ldflags -H=windowsgui