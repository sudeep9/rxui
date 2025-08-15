package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/starfederation/datastar-go/datastar"
	"github.com/sudeep9/rxui/cmd/showcase/pages"
)

type Handlers struct {
}

func (h *Handlers) Register(r *chi.Mux) {
	r.Get("/", h.homePage)
	r.Get("/patients", h.patientsPage)

	r.Post("/ctrl/addtab", h.addTab)
}

func (h *Handlers) homePage(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(r.Context(), w)
}

func (h *Handlers) patientsPage(w http.ResponseWriter, r *http.Request) {
	pages.PatientsPageInst.Page().Render(r.Context(), w)
}

func (h *Handlers) addTab(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	sse.PatchElementTempl(pages.PatientsPageInst.NewTab(),
		datastar.WithSelectorID("patientsTab"),
		datastar.WithModeAppend())
}
