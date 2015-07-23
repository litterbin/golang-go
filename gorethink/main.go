package main

import (
	r "github.com/dancannon/gorethink"
	"log"
)

func main() {
	var session *r.Session

	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		log.Fatalf(err.Error())

	}

	_, err = r.DB("test").TableDrop("authors").Run(session)
	if err != nil {
		log.Fatalf(err.Error())
	}

	resp, err := r.DB("test").TableCreate("authors").RunWrite(session)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("%d table created", resp.TablesCreated)
}
