package integration_test

import (
	. "github.com/Eun/go-hit"
	"net/http"
	"testing"
)

const (
	// Attempts connection
	host = "127.0.0.1:8000"
	// HTTP REST
	basePath = "http://" + host + "/"
)

func TestHTTPEncode(t *testing.T) {
	firstTestBody := `{
	"text": ["aaaa bb cc", "tqwer", "ttt qwe qweyyy"]	
	}`
	firstTestRes := []interface{}{"4a 2b 2c", "tqwer", "3t qwe qwe3y"}
	Test(t,
		Description("Encode First Test"),
		Post(basePath+"rle/encode"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(firstTestBody),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Equal(firstTestRes),
	)
	secondTestBody := `{
	"text": ["a b c", "t w q", "oooooo"]	
	}`
	secondTestRes := []interface{}{"a b c", "t w q", "6o"}
	Test(t,
		Description("Encode Second Test"),
		Post(basePath+"rle/encode"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(secondTestBody),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Equal(secondTestRes),
	)

}

func TestHTTPDecode(t *testing.T) {
	firstTestBody := `{
	"text": ["4a 2b 2c", "tqwer", "3t qwe qwe3y"]	
	}`
	firstTestRes := []interface{}{"aaaa bb cc", "tqwer", "ttt qwe qweyyy"}
	Test(t,
		Description("Decode First Test"),
		Post(basePath+"rle/decode"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(firstTestBody),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Equal(firstTestRes),
	)
	secondTestBody := `{
	"text": ["a b c", "t w q", "6o"]	
	}`
	secondTestRes := []interface{}{"a b c", "t w q", "oooooo"}
	Test(t,
		Description("Decode Second Test"),
		Post(basePath+"rle/decode"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(secondTestBody),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Equal(secondTestRes),
	)

}
