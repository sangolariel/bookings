package render

import (
	"net/http"
	"testing"

	"github.com/sangolariel/bookings/internal/models"
)

func TestDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()

	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("Failed")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "../../templates"

	tc, err := CreateTemplateCatche()

	if err != nil {
		t.Error(err)
	}

	app.TemplateCatche = tc

	r, err := getSession()

	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = Template(&ww, r, "home.page.tmpl", &models.TemplateData{})

	if err != nil {
		t.Error("error writing template to browser")
	}

	err = Template(&ww, r, "non-existent.page.tmpl", &models.TemplateData{})

	if err == nil {
		t.Error("render template that not exist")
	}

}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))

	r = r.WithContext(ctx)

	return r, nil
}

func TestNewTemplates(t *testing.T) {
	NewRenderer(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "../../templates"
	_, err := CreateTemplateCatche()
	if err != nil {
		t.Error(err)
	}
}
