package utils

import (
	"net/http"
	"os"
)

func IsProd() bool {
	return os.Getenv("GIN_ENV") == "production"
}

func SetSecureCookie(w http.ResponseWriter, name, value string, maxAge int) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   IsProd(),
		SameSite: http.SameSiteLaxMode,
	}

	if IsProd() {
		cookie.SameSite = http.SameSiteStrictMode
	}

	http.SetCookie(w, cookie)
}

func DeleteCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   IsProd(),
		SameSite: http.SameSiteStrictMode,
	})
}
