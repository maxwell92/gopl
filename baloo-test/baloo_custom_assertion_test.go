package custom_assertion

import (
	"errors"
	"net/http"
	"testing"

	"gopkg.in/h2non/baloo.v0"
)

var test = baloo.New("http://httpbin.org")

func assert(res *http.Response, req *http.Request) error {
	if res.StatusCode >= 400 {
		return errors.New("Invalid server response (>400)")
	}

	return nil
}

func TestBalooClient(t *testing.T) {
	test.Post("/post").
		SetHeader("Foo", "Bar").
		JSON(map[string]string{"foo": "bar"}).
		Expect(t).
		Status(200).
		Type("json").
		AssertFunc(assert).
		Done()
}
