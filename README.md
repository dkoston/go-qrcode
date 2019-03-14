go-remoteqr
=========

[![CircleCI](https://circleci.com/gh/dkoston/go-remoteqr.svg?style=svg)](https://circleci.com/gh/dkoston/go-remoteqr)

A golang convenience wrapper around [zbar](http://zbar.sourceforge.net/), which will get a remote QR code image and process it.

## Requirements 

To compile this package requires the zbar header files which can be installed on debian/ubuntu with
```
sudo apt-get install libzbar-dev
```

or on OS X with:

```
brew install zbar
```

Go get the library:
```
go get github.com/dkoston/go-remoteqr/qrcode
```

## Usage (Currently under development)

It currently only supports extracting data from a PNG Image. Example Usage:

```go
package main

import (
    "fmt"
    "github.com/dkoston/go-remoteqr/qrcode"
)

func main() {
    results, err := qrcode.GetDataFromPNG("https://domain.com/path/to/image.png")
    if err != nil {
        panic(err)
    }

    for _, result := range results{
        fmt.Printf("Symbol Type: %s, Data %s", result.SymbolType, result.Data )
    }

}
```

## Building

Building a staticlly linked binary with cgo dependencies can be a little fragile, by default cgo libraries are dynamically linked so require the libzbar-dev to be present on machine running your binary. However the following command may work if you want to statically link the zbar libs into your go binary.
```
go build -ldflags "-linkmode external -extldflags -static"
```

## Thanks


Thanks to https://github.com/shezadkhan137/go-qrcode as I forked your wrapper.
