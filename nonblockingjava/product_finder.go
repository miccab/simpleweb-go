package nonblockingjava

import (
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

func Handler() http.HandlerFunc  {
	db, err := sql.Open("postgres", "user=dropwizard dbname=dropwizard sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return func (writer http.ResponseWriter, request *http.Request)  {
		id := request.URL.Query().Get("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
		} else {
			fetchProduct(db, idInt, writer)
		}
	}
}

func fetchProduct(db *sql.DB, productId int,writer http.ResponseWriter) {
	queryResult, err := db.Query("select find_product_name($1)", productId)
	if err != nil {
		log.Fatal(err)
	}
	if (queryResult.Next()) {
		var name string = ""
		queryResult.Scan(&name)
		fmt.Fprintf(writer, "hi product %s", name)
	}
}
