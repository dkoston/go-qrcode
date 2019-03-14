package qrcode

import (
	"testing"
)

func Test_GetDataFromPNG(t *testing.T) {
	testImage := "https://raw.githubusercontent.com/dkoston/go-remoteqr/master/test_images/qr-code.png"
	expectedData := "https://github.com/dkoston/go-remoteqr"


	results, err := GetDataFromPNG(testImage)

	if err != nil {
		t.Error(err)
	}

	if results[0].Data != expectedData {
		t.Errorf("Expected to get (%s). Got (%s)", expectedData, results[0].Data)
	}
}
