package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func createUserHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mongoDAL, err := NewMongoDAL("localhost:27017", "elegant-unit-testing-mocking-go")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/createUser", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("content-type", "application/json")
		var user User
		json.NewDecoder(request.Body).Decode(&user)
		result, err := mongoDAL.CreateUser(user)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(response, err.Error())
			return
		}
		response.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(response, result)
	})

	http.HandleFunc("/getUser", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("content-type", "application/json")
		var user User
		json.NewDecoder(request.Body).Decode(&user)
		println(user.ID)
		foundUser, err := mongoDAL.GetUser(user.ID)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(response, err.Error())
			return
		}
		response.WriteHeader(http.StatusFound)
		j, _ := json.Marshal(foundUser)
		fmt.Fprintf(response, string(j))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
