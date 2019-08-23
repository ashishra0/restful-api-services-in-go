package handlers

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

// UsersRouter handles the users route.
func UsersRouter(w http.ResponseWriter, r *http.Request) {
	// For the user list
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "/users" {
		switch r.Method {
		case http.MethodGet:
			getUsers(w, r)
			return
		case http.MethodPost:
			createUser(w, r)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}
	// For individual user
	path = strings.TrimPrefix(path, "/users/")
	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
		return
	}
	id := bson.ObjectIdHex(path)
	switch r.Method {
	case http.MethodGet:
		getSingleUser(w, r, id)
		return
	case http.MethodPut:
		putUser(w, r, id)
		return
	case http.MethodPatch:
		return
	case http.MethodDelete:
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
