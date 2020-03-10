package weberror_test

import (
	"net/http"
	"strings"
	"testing"
	"weberror"
)

func TestToJSON(t *testing.T) {
	testCases := []struct {
		name           string
		in             weberror.Error
		expectedString string
	}{
		{
			name:           "Test Code and Message",
			in:             weberror.Error{Code: http.StatusBadRequest, Message: "Bad Request test message."},
			expectedString: "{\"code\":400,\"message\":\"Bad Request test message.\"}",
		}, {
			name:           "Test Code only",
			in:             weberror.Error{Code: http.StatusNotFound},
			expectedString: "{\"code\":404}",
		}, {
			name:           "Test Message only",
			in:             weberror.Error{Message: "Test Message"},
			expectedString: "{\"message\":\"Test Message\"}",
		}, {
			name:           "Test empty",
			in:             weberror.Error{},
			expectedString: "{}",
		},
	}

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {
			response := test.in.ToJSON()

			if strings.Compare(string(response), test.expectedString) != 0 {
				t.Log(test.name, "failed")
				t.Fail()
			}
		})
	}
}
