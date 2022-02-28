package handlers

import (
	"github.com/fareskato/kabarda"
	"myapp/data"
	"net/http"
)

// Handlers type: wraps Kabarda type
type Handlers struct {
	App    *kabarda.Kabarda
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering template:", err)
	}

}
