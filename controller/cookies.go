package controller

import (
	"fmt"
	"net/http"
	"github.com/gofrs/uuid"
)

func cookies(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r. Cookies())
	cookie := &http.Cookie{
		Name: "session",
		Value: "testValue",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}