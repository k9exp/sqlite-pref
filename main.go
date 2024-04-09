package main

import (
	"fmt"
	"time"

	"github.com/eatonphil/gosqlite"
)

func main() {
	conn, err := gosqlite.Open("test.db")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = conn.Exec(`CREATE TABLE IF NOT EXISTS performance(id INTEGER PRIMARY KEY AUTOINCREMENT, time INTEGER)`)
	if err != nil {
		panic(err)
	}

	timeStart := time.Now()

	recordWriteCnt := 30_000

	// inserting 100 records
	for i := 0; i < recordWriteCnt; i++ {
		err = conn.Exec(`INSERT INTO performance ('time') VALUES (?)`, i)
		if err != nil {
			panic(err)
		}
	}

	timePassed := time.Since(timeStart)

	fmt.Println("Write / Second: ", recordWriteCnt/int(timePassed.Milliseconds()/1000))
}
