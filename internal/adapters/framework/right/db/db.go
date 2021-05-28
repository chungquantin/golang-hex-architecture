package db

import (
	"database/sql"
	// "database/sql/driver"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type DbAdapter struct {
	db *sql.DB
}

func NewDbAdapter(driverName, dataSourceName string) (*DbAdapter, error){
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("db pin failureL %v", err)
	}

	return &DbAdapter{db: db}, nil
}

func (da DbAdapter) CloseDbConnection(){
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db `close failure`: %v", err)
	}
}

func (da DbAdapter) AddToHistory(answer int32, operation string) error{
	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").
	Values(time.Now(), answer, operation).ToSql()

	if (err != nil){
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if (err != nil){
		return err
	}

	return nil
}