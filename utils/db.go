// package utils

// import (
// 	"database/sql"
// 	"log"
// 	"strconv"

// 	_ "github.com/lib/pq"
// )
  
// var DB *sql.DB

// const (
//     host     = "localhost"
//     port     = 5432
//     user     = "postgres"
//     password = "postgres"
//     dbname   = "postgres"
// )

// const connectionString = "host=" + host + " port=" + strconv.Itoa(port) + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"


// func initDB() {
// 	db, err := sql.Open("postgres", connectionString)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// } 