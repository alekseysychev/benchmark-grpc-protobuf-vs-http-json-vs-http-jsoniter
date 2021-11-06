package httpjsoniter

import (
	"errors"
	"log"
	"net/http"
	"net/mail"

	httpjson "github.com/alekseysychev/benchmark-grpc-protobuf-vs-http-json-vs-http-jsoniter/http-json"
	jsoniter "github.com/json-iterator/go"
)

// Start entrypoint
func Start() {
	http.HandleFunc("/httpjsoniter", CreateUserIter)
	log.Println(http.ListenAndServe(":60002", nil))
}

// CreateUser handler
func CreateUserIter(w http.ResponseWriter, r *http.Request) {
	decoder := jsoniter.NewDecoder(r.Body)
	var user httpjson.User
	decoder.Decode(&user)
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	validationErr := validate(user)
	if validationErr != nil {
		jsoniter.NewEncoder(w).Encode(httpjson.Response{
			Code:    500,
			Message: validationErr.Error(),
		})
		return
	}

	user.ID = "1000000"
	jsoniter.NewEncoder(w).Encode(httpjson.Response{
		Code:    200,
		Message: "OK",
		User:    &user,
	})
}

func validate(in httpjson.User) error {
	_, err := mail.ParseAddress(in.Email)
	if err != nil {
		return err
	}

	if len(in.Name) < 4 {
		return errors.New("Name is too short")
	}

	if len(in.Password) < 4 {
		return errors.New("Password is too weak")
	}

	return nil
}
