package utils

import (
	"strings"
	"testing"
)

func TestGetRequest(t *testing.T) {
	http := HttpRequest{}
	hr, _ := http.Get("http://localhost:8080/ok", nil).Raw()
	println(hr)
	hrp, _ := http.Post("http://localhost:8080/check?aaa=bbb", map[string]string{"bbb": "123"}, strings.NewReader("reader-123")).Raw()
	println(hrp)
}
