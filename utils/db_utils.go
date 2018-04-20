package utils

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

//connect to DB
func ConnectDB() (*sql.DB, error) {

	Db, err := sql.Open("postgres", "postgresql://docker:docker@postgres/docker?sslmode=disable")
	// dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.PORT)
	// Db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		log.Fatal(err)
	}

	for index := 0; index < 10; index++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Establishing Connection .....")

		if err = Db.Ping(); err != nil {
			log.Fatal(err)
		}
		// } else {
		// 	fmt.Println("able to ping DB")
		// 	break
		// }
	}

	return Db, err
}
