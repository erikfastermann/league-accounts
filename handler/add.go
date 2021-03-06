package handler

import (
	"fmt"
	"net/http"

	"github.com/erikfastermann/lam/db"
)

func (h *Handler) add(username string, w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		acc := db.Account{Region: "euw", User: username}
		data := editPage{Title: "Add new account", Users: h.usernames(), Username: username, Account: acc}
		return h.Templates.ExecuteTemplate(w, templateEdit, data)
	}

	acc, err := accFromForm(r)
	if err != nil {
		return badRequestf("failed validating form input, %v", err)
	}

	if err := h.DB.AddAccount(acc); err != nil {
		return fmt.Errorf("writing to database failed, %v", err)
	}

	http.Redirect(w, r, routeOverview, http.StatusSeeOther)
	return nil
}
