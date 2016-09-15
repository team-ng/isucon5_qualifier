package main

import(
	"database/sql"
	"gopkg.in/redis.v4"
	_ "github.com/go-sql-driver/mysql"	
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
	host := "localhost"
	portstr := "3306"
	user := "isucon"
  password := "isucon"
	dbname := "isucon5q"
	ssecret := "beermoris"
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
