package util

import (
	"fmt"
	"io"
	"net/http"
)

// ContentType reads the first 512 bytes of a file to determine it's mime content type. If the first 512 bytes cannot be
// read or the content type cannot be determined, it returns "application/octet-stream"
func ContentType(r io.ReadSeeker) string {
	// we should only be peeking here rather than reading as we miss this data later.
	buff := make([]byte, 512)
	_, err := r.Read(buff)
	// Make sure we rewind the reader back to the start.
	r.Seek(0, 0)
	if err != nil {
		fmt.Println("Error determining content type: ", err)
		return "application/octet-stream"
	}

	return http.DetectContentType(buff)
}
