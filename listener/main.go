package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/lib/pq"
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

	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			panic(err)
		}
	}
	minReconn := 10 * time.Second
	maxReconn := time.Minute
	listener := pq.NewListener(conninfo, minReconn, maxReconn, reportProblem)
	err = listener.Listen(channelName)
	if err != nil {
		panic(err)
	}
	for {
		waitNotification(listener)
	}
}

func waitNotification(l *pq.Listener) {
	for {
		select {
		case notification := <-l.Notify:
			if notification != nil {
				fmt.Println(notification)
			}
			return
		case <-time.After(90 * time.Second):
			go func() {
				if err := l.Ping(); err != nil {
					panic(err)
				}
			}()
			return
		}
	}
}
