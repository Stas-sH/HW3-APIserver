package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUserById(w, r)
	case http.MethodPut:
		adminIdFromCtx := r.Context().Value("adminID")
		adminID, ok := adminIdFromCtx.(int)
		if !ok {
			log.Printf("[%s] %s - some problem with handlerUsersfunc\n", r.Method, r.URL)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if adminID == 0 {
			log.Printf("[%s] %s - access error\n", r.Method, r.URL)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No access for this operation"))
			return
		}
		updateUserById(w, r)
	case http.MethodDelete:
		adminIdFromCtx := r.Context().Value("adminID")
		adminID, ok := adminIdFromCtx.(int)
		if !ok {
			log.Printf("[%s] %s - some problem with handlerUsersfunc\n", r.Method, r.URL)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if adminID == 0 {
			log.Printf("[%s] %s - access error\n", r.Method, r.URL)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No access for this operation"))
			return
		}
		deleteUserById(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented) //для функционала которій должен біть но его пока нет
	}

}

func getUserById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	user, err := GetUserByIdFromDB(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.Write([]byte("No user with id " + id + "; "))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(resp)
}

func deleteUserById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	err := DeleteUserById(id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func updateUserById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updateUser := NewUser()
	if err = json.Unmarshal(req, &updateUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = UpdateUserById(id, updateUser)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}
