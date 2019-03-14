package qrcode

import "testing"

func Test_GetDataFromPNG(t *testing.T) {
	testImage := "https://raw.githubusercontent.com/dkoston/go-remoteqr/master/test_images/qr-code.png"

	results, err := GetDataFromPNG(testImage)

	if err != nil {
		t.Error(err)
	}

	t.Logf("Results: %v", results)
}
