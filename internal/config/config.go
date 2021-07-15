package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCatche      bool
	TemplateCatche map[string]*template.Template
	Info           *log.Logger
	InProduction   bool
	Session        *scs.SessionManager
}
