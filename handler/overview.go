package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/erikfastermann/lam/db"
)

func (h Handler) overview(ctx context.Context, user *db.User, w *response, r *http.Request) (int, string, error) {
	type account struct {
		Color  string
		Banned bool
		Link   string
		db.Account
	}
	type overviewPage struct {
		Username string
		Accounts []account
	}

	db, err := h.db.Accounts(ctx)
	if err != nil {
		return http.StatusInternalServerError, "", fmt.Errorf("couldn't read accounts from database, %v", err)
	}

	accs := make([]account, 0)
	for _, acc := range db {
		banned := false
		if acc.Perma || (acc.Ban.Valid && acc.Ban.Time.After(time.Now())) {
			banned = true
		}
		color := ""
		if banned {
			color = "table-warning"
		}
		if acc.Perma || acc.PasswordChanged {
			color = "table-danger"
		}
		accs = append(accs, account{color, banned, LeagueOfGraphsURL(acc.Region, acc.IGN), *acc})
	}

	data := overviewPage{Username: user.Username, Accounts: accs}
	h.templates.ExecuteTemplate(w, templateOverview, data)
	return http.StatusOK, "", nil
}
