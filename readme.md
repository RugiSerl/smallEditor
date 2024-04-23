# smallEditor

Small project to code an golang ide in golang

## Compiling
### Dependencies
- Go 1.22.1 or later
- C compiler (I recommand [TDM-GCC](https://jmeubank.github.io/tdm-gcc/))
### Building from source
For a debug build:

Get into the repository folder and type:
```bash
go run main.go
```

For a release build:
```bash
go build -ldflags -H=windowsgui
```