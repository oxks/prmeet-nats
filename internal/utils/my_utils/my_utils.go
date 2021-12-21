package my_utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	me "permit_nats/internal/utils/my_errors"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"golang.org/x/crypto/bcrypt"
)

// type to use with templates data
type M map[string]interface{}

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func PasswordCheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// fills in feed["user"] even if user is nil
func LoadAppData(c echo.Context) (M, error) {
	feed := M{}
	sess, err := session.Get("session", c)
	me.ErrorPrint(err)
	feed["user"] = sess.Values["user"]
	if sess.Values["err"] != nil {
		feed["err"] = sess.Values["err"]
		sess.Values["err"] = nil
	}
	if sess.Values["success"] != nil {
		feed["success"] = sess.Values["success"]
		sess.Values["success"] = nil
	}
	defer sess.Save(c.Request(), c.Response())

	return feed, err
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		feed, err := LoadAppData(c)
		me.ErrorPrint(err)

		if feed["user"] == nil {
			feed["err"] = errors.New("Please login first.")

			println("do not pass because not logged in")
			return c.Render(http.StatusBadRequest, "login.go.html", feed)
		}
		return next(c)
	}
}

// func Middleware1(store sessions.Store) echo.MiddlewareFunc {
// 	c := DefaultConfig
// 	c.Store = store
// 	return MiddlewareWithConfig(c)
// }

func SetupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectHandler(func(nc *nats.Conn) {
		log.Printf("Disconnected: will attempt reconnects for %.0fm", totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}

func ShutdownGracefully(nc *nats.Conn) {
	// Setup the interrupt handler to drain so we don't miss
	// requests when scaling down.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println()
	log.Printf("Draining...")
	nc.Drain()
	log.Fatalf("Exiting")
}

func RequestUnmarshall(input []byte) (M, M) {
	request := M{}

	err := json.Unmarshal([]byte(input), &request)
	me.ErrorPrint(err)

	response := M{}

	return request, response
}

func Respond(response M, msg *nats.Msg) {

	response_bytes, err := json.Marshal(response)
	me.ErrorPrint(err)

	msg.Respond(response_bytes)
	// enc.Flush()
}

func NatsConnect() (*nats.EncodedConn, *nats.Conn) {
	// Connect Options.
	opts := []nats.Option{nats.Name("NATS permit")}
	opts = SetupConnOptions(opts)

	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL, opts...)
	me.ErrorPrint(err)

	enc, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	me.ErrorPrint(err)

	return enc, nc
}
