package postgres

import (
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"strings"
	"testing"
	"time"
)

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "siva"
	password = "secret"
	dbname   = "appdb"
)

func TestPostgres(t *testing.T) {
	compose := initPostgresContainer()
	defer compose.Down()

	time.Sleep(time.Second * 5)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func initPostgresContainer() *testcontainers.LocalDockerCompose {
	composeFilePaths := []string{"docker-compose.yml"}
	identifier := strings.ToLower(uuid.New().String())

	compose := testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)
	execError := compose.
		WithCommand([]string{"up", "-d"}).
		WithEnv(map[string]string{
			"key1": "value1",
			"key2": "value2",
		}).
		Invoke()
	err := execError.Error
	if err != nil {
		fmt.Errorf("could not run compose file: %v - %v", composeFilePaths, err)
		return nil
	}
	return compose
}
