package db


import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB


//this function is executed before rest of the functionality is called
func init(){
	var err error
	db, err = sql.Open("postgres",  "postgres://edtuser:.BlackWhiteWithAccent@edt-postgres-server-dev.postgres.database.azure.com/edtsqldev?sslmode=disable")

	if err != nil {
		panic(err.Error())
	}
}

