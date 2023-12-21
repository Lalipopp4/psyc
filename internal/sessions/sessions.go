package sessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key")
	Store = sessions.NewCookieStore(key)
)

func Check(w http.ResponseWriter, r *http.Request) bool {
	session, _ := Store.Get(r, "psycProfiling")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	}
	return true
}

func Login(w http.ResponseWriter, r *http.Request, id int, name string) {
	session, _ := Store.Get(r, "psycProfiling")
	session.Values["authenticated"] = true
	session.Values["id"] = id
	session.Values["name"] = name
	session.Options.MaxAge = 10000
	session.Save(r, w)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "psycProfiling")
	session.Values["authenticated"] = false
	session.Save(r, w)
}
