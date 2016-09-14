package main

import(
	"database/sql"
	"gopkg.in/redis.v4"
	"log"
	"os"
	"strconv"
	"fmt"
)

var (
	db    *sql.DB
	client *redis.Client	
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a >  b {
        return a
    }
    return b
}


func initDB() {
	host := os.Getenv("ISUCON5_DB_HOST")
	if host == "" {
		host = "localhost"
	}
	portstr := os.Getenv("ISUCON5_DB_PORT")
	if portstr == "" {
		portstr = "3306"
	}
	port, err := strconv.Atoi(portstr)
	if err != nil {
		log.Fatalf("Failed to read DB port number from an environment variable ISUCON5_DB_PORT.\nError: %s", err.Error())
	}
	user := os.Getenv("ISUCON5_DB_USER")
	if user == "" {
		user = "isucon"
	}
	password := os.Getenv("ISUCON5_DB_PASSWORD")
	if password == "" {
		password = "isucon"
	}
	dbname := os.Getenv("ISUCON5_DB_NAME")
	if dbname == "" {
		dbname = "isucon5q"
	}
	ssecret := os.Getenv("ISUCON5_SESSION_SECRET")
	if ssecret == "" {
		ssecret = "beermoris"
	}

	db, err = sql.Open("mysql", user+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+dbname+"?loc=Local&parseTime=true")
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s.", err.Error())
	}
	defer db.Close()	
}

func main() {
	
	initDB()

	
	rows, err := db.Query(`SELECT one, another FROM relations`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var one, another int
		key := fmt.Sprintf("rel-%d-%d", min(one, another), max(one, another))
		err := client.Set(key, "1", 0).Err()
		if err != nil {
			panic(err)
		}
	}
}
