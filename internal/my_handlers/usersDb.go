package my_handlers

import (
	"context"
	me "permit_nats/internal/utils/my_errors"
	mu "permit_nats/internal/utils/my_utils"
	"permit_nats/postgres"
)

func GetAllUsers() []postgres.User {
	conn := mu.ConnectDB()
	users, err := conn.UserGetAll(context.Background())
	me.ErrorPrint(err)

	return users
}
func GetEmails() []postgres.UserGetEmailsRow {
	conn := mu.ConnectDB()
	emails, err := conn.UserGetEmails(context.Background())
	me.ErrorPrint(err)

	return emails
}

func GetUserByEmail(email string) (postgres.User, error) {
	conn := mu.ConnectDB()
	user, err := conn.UserGetByEmail(context.Background(), email)
	me.ErrorPrint(err)

	return user, err
}

func UserSignup(args postgres.UserSignupParams) (postgres.User, error) {
	conn := mu.ConnectDB()
	user, err := conn.UserSignup(context.Background(), args)
	me.ErrorPrint(err)
	return user, err
}
