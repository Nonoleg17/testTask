package integration_test

import (
	. "github.com/Eun/go-hit"
	"net/http"
	"testing"
)

const (
	// Attempts connection
	host = "app:8000"
	// HTTP REST
	basePath = "http://" + host + "/"
)

func TestHTTPCreateUser(t *testing.T) {
	body := `{
	"firstname":"Test",
    "surname":"Tester",
    "middleName":"Robot",
    "sex": "M",
     "age":16
		
	}`
	Test(t,
		Description("CreateUser Success"),
		Post(basePath+"user/create-user"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Equal("Success create"),
	)
	body = `{
  	"middleName":"Robot",
    "sex": "W",
     "age":20
	}`
	Test(t,
		Description("CreateUser Error"),
		Post(basePath+"user/create-user"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".error").Equal("invalid request body"),
	)
}

func TestHTTPCreateProduct(t *testing.T) {
	body := `{
	"description":"Test",
    "price":203,
    "currency":"dollars",
    "left_in_stock": 10
		
	}`
	Test(t,
		Description("CreateProduct Success"),
		Post(basePath+"product/create-product"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Equal("Success create"),
	)
	body = `{
  	"description":"Test_prod",
    "price":2032,
    "currency":"dollars",
    "left_in_stock": -1
	}`
	Test(t,
		Description("CreateProduct Error"),
		Post(basePath+"product/create-product"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusInternalServerError),
		Expect().Body().JSON().JQ(".error").Equal("create Product problem"),
	)
}
