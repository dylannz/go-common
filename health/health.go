package health

import (
	"fmt"
	"net/http"
)

var okMessage = []byte("ok\n")

// Handler runs the check func param, if an error is returned from checks, a 500 status response will be written with
// checks error, otherwise a 200 status response will be written.
// a 500 status is returned.
func Handler(checks func() error) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := checks(); err != nil {
			http.Error(
				w,
				fmt.Sprintf("500 internal server error: service not healthy: %v", err),
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Length", fmt.Sprintf("%v", len(okMessage)))
		w.WriteHeader(http.StatusOK)
		w.Write(okMessage)
	})
}
