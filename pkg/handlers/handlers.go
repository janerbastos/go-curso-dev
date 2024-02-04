package handlers

import (
	"net/http"

	"github.com/janerbastos/go-curso/pkg/config"
	"github.com/janerbastos/go-curso/pkg/hender"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repository create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers seta the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	hender.HenderTemplate(w, "home.page.html")
}

func (m *Repository) Abaut(w http.ResponseWriter, r *http.Request) {
	hender.HenderTemplate(w, "abaut.page.html")
}
