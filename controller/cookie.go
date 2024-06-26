package controller

import (
	model "forum/model"
	"net/http"
	"strconv"
	"github.com/gofrs/uuid"
)

func newSession(w http.ResponseWriter, r *http.Request, id int) {
	cookie, err := newCookie()
	if err != nil {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	} else {
		err := model.NewSession(cookie.Value, id)
		
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		} else {
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func newCookie() (*http.Cookie, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	uuidS := uuid.String()

	cookie := &http.Cookie{
		Name: "session",
		Value: uuidS,
		Secure: true,
		HttpOnly: true,
	}
	return cookie, nil
}

func disconnect(w http.ResponseWriter, r *http.Request) {
	idUser, err := strconv.Atoi(r.FormValue("idUser"))
	if err != nil {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
	model.NewSession(nil, idUser)
	http.Redirect(w, r, "/", http.StatusFound)
}