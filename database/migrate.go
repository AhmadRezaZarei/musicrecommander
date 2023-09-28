package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var currentTryCount = 0

const maxTryCount = 10

func Migrate() {

	for i := 0; i < maxTryCount; i++ {

		time.Sleep(2 * time.Second)

		err := _migrate()
		if err != nil {
			currentTryCount += 1
		}

		if err == nil {
			return
		}

		if i == maxTryCount-1 && err != nil {
			fmt.Println("Can not migrate")
			panic(err)
		}

	}

}

func _migrate() error {

	connectionString := os.Getenv("DB_CONNECTION_STRING")

	fmt.Println("conneciton string: |" + connectionString + "|")

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		fmt.Printf("try: %d con not open sql connection err: %v \n", currentTryCount, err)
		return err
	}

	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Printf("try: %d can not create driver instance: %v \n", currentTryCount, err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migration",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Printf("try: %d create migrate instance : %v \n", currentTryCount, err)
		return err
	}

	err = m.Up()

	if err == migrate.ErrNoChange {
		fmt.Println("db migrated successfully !!")
		return nil
	}

	if err != nil {
		fmt.Printf("try: %d can not perform migrate up : %v \n", currentTryCount, err)
		return err
	}

	return nil
}
