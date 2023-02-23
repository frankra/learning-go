package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//BuildSQLConnectionString ...
func buildSQLConnectionString(user, password, host, port, schema string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&tx_isolation='READ-COMMITTED'",
		user, password, host, port, schema)
}

func spawnRunners(db *sqlx.DB) {
	lockNames := []string{"lock1", "lock2", "lock2", "lock1", "lock3", "lock4"}

	for _, lockName := range lockNames {
		go runner(db, lockName)
	}
}

func runner(db *sqlx.DB, lockName string) {
	res, err := db.Exec("select get_lock(?, ?)", lockName, -1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Get lock for %s %v \n", lockName, res)
}

func main() {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Open("mysql", buildSQLConnectionString("root", "root", "localhost", "33006", "vendor_svc_po"))
	if err != nil {
		log.Fatalln(err)
	}

	db.DB.SetMaxOpenConns(2)

	spawnRunners(db)
	time.Sleep(10 * time.Minute)
}
