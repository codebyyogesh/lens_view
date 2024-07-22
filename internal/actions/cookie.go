package actions

import (
	"fmt"
	"net/http"
)

const (
	CookieSession = "session"
)

func newCookie(name, value string) *http.Cookie {
	return &(http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",  // ie this cookie is accesible at all paths
		HttpOnly: true, // protect cookies from XSS using Javascript
	})
	//above is same as, cookie := http.Cookie{...} and then return &cookie
}

func setCookie(w http.ResponseWriter, name, value string) {
	cookie := newCookie(name, value)
	http.SetCookie(w, cookie)
}

func readCookie(r *http.Request, name string) (string, error) {
	c, err := r.Cookie(name)
	if err != nil {
		return "", fmt.Errorf("%s, %w", name, err)
	}
	return c.Value, nil
}
