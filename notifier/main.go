package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const channelName = "channel"

func main() {
	var conninfo string = os.Args[1]
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		panic(err)
	}

	for {
		message := uuid.New().String()
		_, err := db.Exec(fmt.Sprintf("NOTIFY %s, '%s'", channelName, message))
		if err != nil {
			panic(err)
		}
		time.Sleep(3 * time.Second)
	}
}
