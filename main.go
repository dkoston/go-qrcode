package main

import (
	"flag"
	"fmt"
	"github.com/dkoston/go-remoteqr/qrcode"
)

var image *string

//go build -o test -ldflags "-linkmode external -extldflags -static"
func init() {
	image = flag.String("i", "", "image url")
}

func main() {

	flag.Parse()

	results, err := qrcode.GetDataFromPNG(*image)
	if err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("Symbol Type: %s, Data %s", result.SymbolType, result.Data)
	}
}
