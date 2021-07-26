package render

import (
	"encoding/gob"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sangolariel/bookings/internal/config"
	"github.com/sangolariel/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	gob.Register(models.Reservation{})
	//Env
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = false

	app.Session = session

	app = &testApp

	os.Exit(m.Run())
}
