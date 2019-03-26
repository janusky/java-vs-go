package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/janusky/java-vs-go/app-go/app"
)

// our main function
func main() {
	fmt.Println("app-go Begin")
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/client/new/{deposit}", NewClientHandler).Methods("POST")
	router.HandleFunc("/transaction", NewTransactionHandler).Methods("POST")
	router.HandleFunc("/client/{id}/balance", BalanceHandler).Methods("GET")

	// Creación de la base de datos en modo desarrollo o test
	//app.Init()
	defer app.CloseDB()

	fmt.Println("app-go Listen")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`{"status":"Ok"}`))
}

func NewTransactionHandler(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t app.Transaction
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	new_transaction := app.NewTransaction(t.From_client_id, t.To_client_id, t.Amount)
	json.NewEncoder(writer).Encode(new_transaction)
}

func NewClientHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	s := params["deposit"]
	i, _ := strconv.Atoi(s)
	client := app.NewClient(i)
	json.NewEncoder(writer).Encode(client)
}

func BalanceHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	s := params["id"]
	i, _ := strconv.Atoi(s)
	balance := app.CheckBalance(i)
	response := app.Balance{balance}
	json.NewEncoder(writer).Encode(response)
}
