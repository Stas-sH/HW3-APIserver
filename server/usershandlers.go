package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func WeryffyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		adminName := r.Header.Get("admin-name")
		adminPassword := r.Header.Get("admin-password")
		var isAdmin bool
		admins, err := GetAdminsFromDB()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if (adminName == "") && (adminPassword == "") {
			isAdmin = false
		} else if (chekAdminName(adminName, admins) == 0) || (chekAdminPassword(adminPassword, admins) == 0) {
			log.Printf("trying to log in as administrator\n")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You are not Administrator"))
			return
		} else {
			isAdmin = true
		}

		ctx := r.Context()
		if isAdmin {
			adminID := chekAdminName(adminName, admins)
			ctx = context.WithValue(ctx, "adminID", adminID)
			r = r.WithContext(ctx)
		} else {
			ctx = context.WithValue(ctx, "adminID", 0)
			r = r.WithContext(ctx)
		}

		next(w, r)
	}
}

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		adminIdFromCtx := r.Context().Value("adminID")
		adminID, ok := adminIdFromCtx.(int)
		if !ok {
			log.Printf("[%s] %s - some problem with adminID\n", r.Method, r.URL)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if adminID == 0 {
			log.Printf("[%s] %s - by someone\n", r.Method, r.URL)
		} else {
			log.Printf("[%s] %s - by administrator with id [%d]\n", r.Method, r.URL, adminID)
		}
		next(w, r)
	}
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
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
		addUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented) //для функционала которій должен біть но его пока нет
	}

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := GetUsersFromDB()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(resp)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newUser := NewUser()
	if err = json.Unmarshal(req, &newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = AddUserToDB(newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func chekAdminPassword(str string, admins []Administrator) int {

	for i := 0; i < len(admins); i++ {
		if admins[i].Password == str {
			return admins[i].Id
		}
	}
	return 0
}

func chekAdminName(str string, admins []Administrator) int {

	for i := 0; i < len(admins); i++ {
		if admins[i].AdminName == str {

			return admins[i].Id
		}
	}
	return 0
}
