package utils_test

import (
	"github.com/aasumitro/gowa/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHttpRespond(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		data     interface{}
		expected interface{}
	}{
		{
			name:     "success with no pagination",
			code:     http.StatusOK,
			data:     []string{"foo", "bar"},
			expected: utils.SuccessRespond{Code: http.StatusOK, Status: "OK", Data: []string{"foo", "bar"}},
		},
		{
			name:     "error with data",
			code:     http.StatusBadRequest,
			data:     "invalid request",
			expected: utils.ErrorRespond{Code: http.StatusBadRequest, Status: "Bad Request", Data: "invalid request"},
		},
		{
			name:     "error with no data",
			code:     http.StatusBadRequest,
			expected: utils.ErrorRespond{Code: http.StatusBadRequest, Status: "Bad Request", Data: "something went wrong with the request"},
		},
		{
			name:     "error with no data and server error code",
			code:     http.StatusInternalServerError,
			expected: utils.ErrorRespond{Code: http.StatusInternalServerError, Status: "Internal Server Error", Data: "something went wrong with the server"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(writer)
			utils.NewHTTPRespond(c, test.code, test.data)
		})
	}
}
