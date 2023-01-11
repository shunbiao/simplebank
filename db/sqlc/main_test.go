package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)



const (
	dbDrive = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB


// func getDB() *sql.DB{
	
// 	testDB, _ := sql.Open(dbDrive, dbSource)
// 	return testDB
// }

func TestMain(m *testing.M){
	var err error
	testDB, err := sql.Open(dbDrive, dbSource)
	if err !=nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	exitVal := m.Run()

	if exitVal==0 {
		os.Exit(exitVal)
	}
}

