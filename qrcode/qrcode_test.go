package qrcode

import (
	"testing"
)

func Test_GetDataFromPNG(t *testing.T) {
	type testCase struct {
		ImageUrl     string
		ExpectedData string
	}

	testCases := []testCase{
		{
			ImageUrl:     "https://raw.githubusercontent.com/dkoston/go-remoteqr/master/test_images/go-remoteqr.png",
			ExpectedData: "https://github.com/dkoston/go-remoteqr",
		},
		{
			ImageUrl:     "https://raw.githubusercontent.com/dkoston/go-remoteqr/master/test_images/test-url-1.png",
			ExpectedData: "https://test-url-1",
		},
		{
			ImageUrl:     "https://raw.githubusercontent.com/dkoston/go-remoteqr/master/test_images/ftp.png",
			ExpectedData: "ftp://this_is_a_test",
		},
	}

	for i := 0; i < len(testCases); i++ {
		results, err := GetDataFromPNG(testCases[i].ImageUrl)

		if err != nil {
			t.Error(err)
		}

		if results[0].Data != testCases[i].ExpectedData {
			t.Errorf("Expected to get (%s). Got (%s). Test case %d", testCases[i].ExpectedData, results[0].Data, i)
		}
	}
}
