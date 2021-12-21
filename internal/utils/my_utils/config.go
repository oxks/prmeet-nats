package my_utils

import (
	"database/sql"
	"fmt"
	"log"
	me "permit_nats/internal/utils/my_errors"
	"permit_nats/postgres"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/gobuffalo/envy"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"github.com/rbcervilla/redisstore"
	"gopkg.in/gomail.v2"
)

// connect nats
func ConnectNats() *nats.Conn {

	nats_url, err := envy.MustGet("NATS_URL")
	me.ErrorPrint(err)
	nc, err := nats.Connect(nats_url)
	if err != nil {
		log.Fatal(err)
	}
	return nc
}

// connect to db
func ConnectDB() *postgres.Queries {
	host, err := envy.MustGet("POSTGRES_HOST")
	me.ErrorPrint(err)
	port_str, err := envy.MustGet("POSTGRES_PORT")
	me.ErrorPrint(err)
	port, err := strconv.Atoi(port_str)
	me.ErrorPrint(err)
	user, err := envy.MustGet("POSTGRES_USER")
	me.ErrorPrint(err)
	password, err := envy.MustGet("sodhfah111lsk")
	me.ErrorPrint(err)
	dbname, err := envy.MustGet("POSTGRES_DB_NAME")
	me.ErrorPrint(err)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	q := postgres.New(db)

	return q
}

func SendMail(message string) {
	email, err := envy.MustGet("EMAIL")
	me.ErrorPrint(err)
	smtp_server, err := envy.MustGet("SMTP_SERVER")
	me.ErrorPrint(err)
	smtp_key, err := envy.MustGet("SMTP_APPLICATION_KEY")
	me.ErrorPrint(err)
	smtp_port, err := envy.MustGet("SMTP_PORT")
	me.ErrorPrint(err)
	smtp_port_int, err := strconv.Atoi(smtp_port)
	me.ErrorPrint(err)

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Prmeet email")
	m.SetBody("text/html", message)

	d := gomail.NewDialer(smtp_server, smtp_port_int, "email", smtp_key)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func GetCookieStore() *redisstore.RedisStore {
	rds, err := envy.MustGet("REDIS_URL")
	me.ErrorPrint(err)
	client := redis.NewClient(&redis.Options{
		Addr: rds,
	})

	// New default RedisStore
	store, err := redisstore.NewRedisStore(client)
	if err != nil {
		log.Fatal("failed to create redis store: ", err)
	}
	return store
}
