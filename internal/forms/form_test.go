package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Validate(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)

	isValid := form.Validate()

	if !isValid {
		t.Errorf(("got invalid when should have been valid"))
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)

	form.Required("a", "b", "c")

	if form.Validate() {
		t.Errorf("forms show valid when required show missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if !form.Validate() {
		t.Errorf("show doest have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)

	has := form.Has("whatever")

	if has {
		t.Errorf(("form shows the field when it doest not"))
	}

	postedData := url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)

	has = form.Has("a")

	if !has {
		t.Errorf(("show forms doest have field when it should"))
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)

	form.MinLength("x", 10)

	if form.Validate() {
		t.Errorf("form shows min length for show non-exist field")
	}

	isError := form.Errors.Get("x")

	if isError == "" {
		t.Errorf("should have an err but not get one")
	}

	postedData := url.Values{}
	postedData.Add("some_field", "some value")

	form.MinLength("some_field", 100)

	if form.Validate() {
		t.Errorf("show min length 100 met when data shorter")
	}

	postedData = url.Values{}
	postedData.Add("another_field", "abc123")

	form.MinLength("another_field", 1)

	if form.Validate() {
		t.Errorf("show min length of 1 not met when it is")
	}

	isError = form.Errors.Get("another_field")

	if isError == "" {
		t.Errorf("should not have an err but got one")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}

	form := New(postedData)

	form.IsEmail("x")

	if form.Validate() {
		t.Errorf("form show valid email for non-existent field")
	}

	postedData = url.Values{}
	postedData.Add("email", "s@gmail.com")

	form = New(postedData)

	form.IsEmail("email")

	if !form.Validate() {
		t.Errorf("got an invalid email but it should have")
	}

	postedData = url.Values{}
	postedData.Add("email", "x")

	form = New(postedData)

	form.IsEmail("email")

	if form.Validate() {
		t.Errorf("got an valid for invalid email adress")
	}
}
