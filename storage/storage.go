package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/galenguyer/retina/core"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func CreateDatabase() {
	db, _ = sql.Open("sqlite3", "file:retina.db?cache=shared")
	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS results (
			id INTEGER PRIMARY KEY, 
			name TEXT, 
			timestamp DATETIME, 
			statuscode INTEGER, 
			duration INTEGER,
			certexpiry INTEGER
		)`)
	if err != nil {
		log.Panic(err)
	}
	statement.Exec()
}

func InsertResult(res *core.Result) {
	statement, err := db.Prepare("INSERT INTO results (name, timestamp, statuscode, duration, certexpiry) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		log.Panic(err)
	}
	statement.Exec(res.ServiceName, res.Timestamp.Unix(), res.HTTPStatusCode, res.Duration.Milliseconds(), (res.CertificateExpiry.Milliseconds() / 1000))
	js, _ := json.Marshal(res)
	log.Println("inserting", string(js))

	rows, err := db.Query("SELECT * FROM results;")
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var timestamp time.Time
		var statuscode int
		var duration int
		var certexpiry int
		err = rows.Scan(&id, &name, &timestamp, &statuscode, &duration, &certexpiry)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(strconv.Itoa(id)+":", name, timestamp, statuscode, duration, certexpiry)
	}
}
