package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/shunbiao/simplebank/util"
)

var testQueries *Queries

// var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("connot load config:", err)
	}

	testDB, err := sql.Open(config.DBDrive, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	exitVal := m.Run()

	if exitVal == 0 {
		os.Exit(exitVal)
	}
}
