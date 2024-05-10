package controller

import (
	"github.com/gofrs/uuid"
	"net/http"
)

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

func checkCookie() {
	
}