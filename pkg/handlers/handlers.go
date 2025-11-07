package handlers

import (
	"log"
	"github.com/SiddhantSShende/bookings-app/pkg/config"
	"github.com/SiddhantSShende/bookings-app/pkg/models"
	"github.com/SiddhantSShende/bookings-app/pkg/render"
	"net/http"
	"time"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	log.Println("remoteIP:", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About page handler that uses addValues function
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	m.App.Session.Lifetime = 24 * time.Hour // This page will have 24 hr lifetime
	// m.App.Session.Put(r.Context(), "test", "Hello, again.")

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
