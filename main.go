package main

import (
        "database/sql"
        _ "github.com/lib/pq"
        "fmt"
        "log"

)


func main() {
	ReadSomeData()
}


var db *sql.DB


//this function is executed before rest of the functionality is called
func init(){
        cs := "postgres://..."
        var err error
        db, err = sql.Open("postgres",  cs)

        if err != nil {
                panic(err.Error())
        }
}


func ReadSomeData() error {
        rows, err := db.Query("SELECT id, issuer  FROM  edt_issuers")
        FailOnError(err,"querying db")
        for rows.Next() {
                var id int
                var issuer string
                err = rows.Scan(&id, &issuer)
                FailOnError(err, "reading sensor from DB")
                fmt.Printf("%v %s", id, issuer)
        }
        return err
}


func FailOnError(err error, msg string){
        if err != nil {
                log.Fatalf("%s: %s", msg, err)
                panic(fmt.Sprintf("%s: %s", msg, err))
        }
}


