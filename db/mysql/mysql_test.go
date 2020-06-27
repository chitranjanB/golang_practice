package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/docker/go-connections/nat"
	_ "github.com/go-sql-driver/mysql"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"strconv"
	"testing"
)

func TestMySQL(t *testing.T) {
	dataSourceName, mysqlTerm := initMySQLContainer()
	defer mysqlTerm()
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	createTable(db)
	insertData(db)
	readData(db)
}

func createTable(db *sql.DB) {
	_, err := db.Exec("create table person(id int primary key auto_increment, name varchar(100))")
	if err != nil {
		panic(err.Error())
	}
}

func insertData(db *sql.DB) {
	for i := 1; i <= 5; i++ {
		result, err := db.Exec("insert into person(name) values(?)", "User"+strconv.Itoa(i))
		if err != nil {
			panic(err.Error())
		}
		id, _ := result.LastInsertId()
		fmt.Println("User Id: ", id)
	}
}

func readData(db *sql.DB) {
	rows, err := db.Query("select id, name from person")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func initMySQLContainer() (string, func()) {
	ctx := context.Background()
	mysqlPort, _ := nat.NewPort("tcp", "3306")
	req := testcontainers.ContainerRequest{
		Image:        "mysql",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "admin",
			"MYSQL_DATABASE":      "test",
		},
		WaitingFor: wait.ForListeningPort(mysqlPort),
	}
	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}

	ip, err := mysqlC.Host(ctx)
	if err != nil {
		panic(err)
	}
	port, err := mysqlC.MappedPort(ctx, "3306")
	if err != nil {
		panic(err)
	}
	dataSourceName := fmt.Sprintf("root:admin@tcp(%s:%d)/test", ip, port.Int())
	fmt.Println(dataSourceName)
	cTerm := func() {
		defer mysqlC.Terminate(ctx)
	}
	return dataSourceName, cTerm
}
