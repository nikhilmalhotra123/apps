package main

import (
	"github.com/nikhilmalhotra123/apps/jobapps"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter()

  router.HandleFunc("/signup", jobapps.SignUpHandler).Methods("POST")
  router.HandleFunc("/login", jobapps.LoginHandler).Methods("POST")
  router.HandleFunc("/user", jobapps.ProfileHandler).Methods("GET")
  router.HandleFunc("/user/delete", jobapps.DeleteHandler).Methods("DELETE")

	router.HandleFunc("/user/app/{id}", jobapps.GetApplicationHandler).Methods("GET")
	router.HandleFunc("/user/apps", jobapps.GetAllApplicationsHandler).Methods("GET")
  router.HandleFunc("/user/app", jobapps.InsertApplicationHandler).Methods("POST")
	router.HandleFunc("/user/app/{id}",
	jobapps.UpdateApplicationHandler).Methods("PUT")
	router.HandleFunc("/user/app/{id}",
	jobapps.DeleteApplicationHandler).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8080", router))
}
