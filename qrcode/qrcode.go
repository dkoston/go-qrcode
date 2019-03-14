package qrcode

// #cgo linux LDFLAGS: -lzbar -lpng -ljpeg -lz -lrt -lm -pthread
// #cgo darwin LDFLAGS: -lzbar -lpng -ljpeg -lz -lm -pthread
// #include <stdio.h>
// #include <stdlib.h>
// #include <png.h>
// #include <zbar.h>
// #include "get_data.h"
// typedef void (*zbar_image_set_data_callback)(zbar_image_t *  image);
import "C"
import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"unsafe"
)

type Result struct {
	SymbolType string
	Data       string
}

var client = http.Client{
	Timeout: 5 * time.Second,
}

func GetDataFromPNG(imgUrl string) (results []Result, err error) {
	pngPath, err := GetPNGFromWeb(imgUrl)

	pth := C.CString(pngPath)
	scanner := C.zbar_image_scanner_create()
	C.zbar_image_scanner_set_config(scanner, 0, C.ZBAR_CFG_ENABLE, 1)

	defer C.zbar_image_scanner_destroy(scanner)

	var width, height C.int = 0, 0
	var raw unsafe.Pointer = nil
	errorCode := C.get_data(pth, &width, &height, &raw)
	if int(errorCode) != 0 {
		err = errors.New(fmt.Sprintf("Error reading from png file. Error code %d", errorCode))
		return
	}

	image := C.zbar_image_create()

	defer C.zbar_image_destroy(image)

	C.zbar_image_set_format(image, C.ulong(808466521))
	C.zbar_image_set_size(image, C.uint(width), C.uint(height))


	C.zbar_image_set_data(image, raw, C.ulong(width*height), nil)

	C.zbar_scan_image(scanner, image)

	symbol := C.zbar_image_first_symbol(image)

	for ; symbol != nil; symbol = C.zbar_symbol_next(symbol) {
		typ := C.zbar_symbol_get_type(symbol)
		data := C.zbar_symbol_get_data(symbol)
		symbolType := C.GoString(C.zbar_get_symbol_name(typ))
		dataString := C.GoString(data)
		results = append(results, Result{symbolType, dataString})
	}

	err = os.Remove(pngPath)

	return
}

func GetPNGFromWeb(imgUrl string) (imgPath string, err error){
	imgResponse, err := client.Get(imgUrl)
	defer imgResponse.Body.Close()
	if err != nil {
		return imgPath, err
	}

	u, err := uuid.NewV4()
	if err != nil {
		return imgPath, err
	}

	imgPath = fmt.Sprintf("%s.png",u.String())

	f, err := os.Create(imgPath)
	defer f.Close()
	if err != nil {
		return imgPath, err
	}

	body, err := ioutil.ReadAll(imgResponse.Body)
	if err != nil {
		return imgPath, err
	}

	err = ioutil.WriteFile(imgPath, body, 0644)

	return imgPath, err
}
