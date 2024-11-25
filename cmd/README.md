# Building

Install libxls with `brew install libxls`

Then build the go binary as follows:

```shell
GO_CFLAGS=-I/opt/homebrew/Cellar/libxls/1.6.2/include CGO_LDFLAGS="-L/opt/homebrew/Cellar/libxls/1.6.2/lib/ -l xlsreader" go build main.go
```