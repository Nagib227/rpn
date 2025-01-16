package rpn_test

import (
	"net/http"
	"net/http/httptest"
	// "strconv"
	"testing"
	// "time"
	"github.com/Nagib227/rpn/pkg/rpn"
)

func TestFibonacciHandler(t *testing.T) {
	// Clear the request count before testing
	// requestCount = 0	

	tt := []struct {
		name           string
		expectedStatus int
		expectedBody   string
	}{
		{"First Fibonacci Number", http.StatusOK, "0"},
		{"Second Fibonacci Number", http.StatusOK, "1"},
		{"Third Fibonacci Number", http.StatusOK, "1"},
		{"Fourth Fibonacci Number", http.StatusOK, "2"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(rpn.FibonacciHandler)

			handler.ServeHTTP(rr, req)

			if tc.expectedStatus != rr.Code {
				t.Errorf("failed error code")
			}
			// t.Errorf("'%s', '%s',", rr.Body.String(), tc.expectedBody)
			if tc.expectedBody != rr.Body.String() {
				t.Errorf("failed body")
			}
		})
	}
}