package app

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/janusky/java-vs-go/app-go/app/utils"
	_ "github.com/lib/pq"
)

var DB_HOST = utils.GetEnv("POSTGRES_HOST", "localhost")
var DB_USER = utils.GetEnv("POSTGRES_USER", "postgres")
var DB_PASSWORD = utils.GetEnv("POSTGRES_PASSWORD", "test1234")
var DB_NAME = utils.GetEnv("POSTGRES_NAME", "postgres")

var DB *sql.DB

func Init() {
	fmt.Println("Created DB - Begin")
	db, e := newDb()
	checkErr(e)
	_, e = db.Query("DROP TABLE IF EXISTS client")
	_, e = db.Query("DROP TABLE IF EXISTS transaction")
	checkErr(e)
	_, e = db.Query("CREATE TABLE client(id SERIAL PRIMARY KEY NOT NULL, name VARCHAR(20), email VARCHAR(20), phone VARCHAR(20));")
	_, e = db.Query("CREATE TABLE transaction(id SERIAL PRIMARY KEY NOT NULL, from_client_id INTEGER, to_client_id INTEGER, amount INTEGER);")
	checkErr(e)
	//db.Close()
	fmt.Println("Created DB - End")
}

func CreateClient(client Client) Client {
	db, err := newDb()
	checkErr(err)
	var id int
	err = db.QueryRow(
		"INSERT INTO client(name, email, phone) VALUES ($1, $2, $3) RETURNING id",
		client.Name, client.Email, client.Phone).Scan(&id)
	fmt.Println("Created client with id", id)
	checkErr(err)
	//db.Close()
	client.Id = id
	return client
}

func CreateTransaction(trans Transaction) Transaction {
	db, err := newDb()
	checkErr(err)
	var id int
	err = db.QueryRow(
		"INSERT INTO transaction(from_client_id, to_client_id, amount) VALUES ($1, $2, $3) RETURNING id",
		trans.From_client_id, trans.To_client_id, trans.Amount).Scan(&id)
	fmt.Println("Created transaction with id", id)
	checkErr(err)
	//db.Close()
	trans.Id = id
	return trans
}

func GetBalance(client_id int) int {
	db, err := newDb()
	var balance int
	err = db.QueryRow(`
				SELECT debit - credit
				FROM
				  (
					SELECT COALESCE(sum(amount), 0) AS debit
					FROM transaction
					WHERE to_client_id = $1
				  ) a,
				  (
					SELECT COALESCE(sum(amount), 0) AS credit
					FROM transaction
					WHERE from_client_id = $1
				  ) b;
		`, client_id).Scan(&balance)
	checkErr(err)
	//db.Close()
	fmt.Println("Calculated balance with client id", client_id)
	return balance
}

func newDb() (*sql.DB, error) {
	var err error
	if DB == nil {
		dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)
		DB, err = sql.Open("postgres", dbinfo)
		DB.SetMaxOpenConns(20) // Sane default
		DB.SetMaxIdleConns(0)
		DB.SetConnMaxLifetime(time.Nanosecond)
	}
	return DB, err
}

func CloseDB() {
	fmt.Println("Close db")
	if DB != nil {
		err := DB.Close()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Close FAIL")
		} else {
			fmt.Println("Close OK")
		}
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
