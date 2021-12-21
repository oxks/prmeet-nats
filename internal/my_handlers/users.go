package my_handlers

import (
	"fmt"
	mu "permit_nats/internal/utils/my_utils"
	"permit_nats/postgres"

	"github.com/mitchellh/mapstructure"
	"github.com/nats-io/nats.go"
)

func GetAllUsersHandler(msg *nats.Msg) {

	request, response := mu.RequestUnmarshall(msg.Data)
	fmt.Printf("\nThe request is: %v\n", request)

	users := GetAllUsers()

	response["users"] = users

	mu.Respond(response, msg)
}

func GetAllEmailsHandler(msg *nats.Msg) {

	request, response := mu.RequestUnmarshall(msg.Data)
	fmt.Printf("\nThe request is: %v\n", request)

	users := GetAllUsers()

	response["users"] = users

	mu.Respond(response, msg)
}

func GetUserByEmailHandler(msg *nats.Msg) {

	request, response := mu.RequestUnmarshall(msg.Data)
	fmt.Printf("\nThe request for email is: %v\n", request)

	user, err := GetUserByEmail(request["email"].(string))

	response["user"] = user
	response["err"] = err

	mu.Respond(response, msg)
}
func UserSignupHandler(msg *nats.Msg) {

	request, response := mu.RequestUnmarshall(msg.Data)

	var args postgres.UserSignupParams

	u := request["user"].(map[string]interface{})
	mapstructure.Decode(u, &args)

	fmt.Printf("\nThe args for UserSignup is: %v\nAnd it's type is: %T", request["user"], request["user"])
	user, err := UserSignup(args)

	response["user"] = user
	response["err"] = err

	mu.Respond(response, msg)
}
