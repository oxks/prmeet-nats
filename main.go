package main

import (
	"log"
	mh "permit_nats/internal/my_handlers"
	mu "permit_nats/internal/utils/my_utils"
)

func main() {

	enc, nc := mu.NatsConnect()

	enc.Subscribe("index", mh.GetAllUsersHandler)
	enc.Subscribe("users", mh.GetAllEmailsHandler)
	enc.Subscribe("users.login.getByEmail", mh.GetUserByEmailHandler)
	enc.Subscribe("users.signup", mh.UserSignupHandler)

	enc.Flush()

	if err := enc.LastError(); err != nil {
		log.Fatal(err)
	}

	mu.ShutdownGracefully(nc)
}
