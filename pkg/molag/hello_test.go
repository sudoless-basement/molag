package molag

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestHello(t *testing.T) {
	Hello()

	_, _ = http.Get("http://example.net")
	fmt.Println(io.EOF.Error())
}
