package db

import (
	"fmt"
	"log"
)

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