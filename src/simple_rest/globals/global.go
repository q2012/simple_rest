package globals

import (
	"context"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/pgxpool"
	"os"
	"simple_rest/models"
)

var postgresConnString string
var store *sessions.CookieStore
var all[] *models.Main
var db *pgxpool.Pool


func SetPostgresConnString(l string) {
	postgresConnString = l
}

func Pool() *pgxpool.Pool {
	return db
}

func InitGlobals() error {
	initSession()

	var err = initDatabaseConnection()
	return err
}

func Store() *sessions.CookieStore {
	return store
}

func GetAll() []*models.Main {
	return all
}

func AppendToAll(one *models.Main, toAdd ...*models.Main) []*models.Main {
	toAdd = append(toAdd, one)
	all = append(all, toAdd...)
	return all
}

func initDatabaseConnection() error {
	poolConfig, err := pgxpool.ParseConfig(postgresConnString)
	if err != nil {
		fmt.Println("Unable to parse DATABASE_URL", "error", err)
		os.Exit(1)
	}

	db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		fmt.Println("Unable to create connection pool", "error", err)
		os.Exit(1)
	}
	return err
}

func initSession() {
	//store, err := sessions.NewRediStore(10, "tcp", ":6379", "", []byte(os.Getenv("SESSION_KEY")))
	//if err != nil {
	//	panic(err)
	//}
	//defer store.Close()

	//store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
}