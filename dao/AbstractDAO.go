package dao

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "testing"
	DB_NAME     = "bpjamsostek"
)

type FieldStatus struct {
	IsCheck   bool
	FieldName string
	Value     interface{}
}

type DefaultFieldMustCheck struct {
	ID        FieldStatus
	Deleted   FieldStatus
	Status    FieldStatus
	CreatedBy FieldStatus
	UpdatedAt FieldStatus
}

func DBConnection() *sql.DB {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=localhost port=5432", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	CheckError(err)
	if err == nil {
		log.Println("Connected!")
	}

	return db
}

//func DbConn() {
//	_, err := gorm.Open("mysql", "root:paramadaksa@/dataset?charset=utf8&parseTime=True")
//	CheckError(err)
//
//	if err == nil {
//		log.Println("Connected!")
//	}
//}

func CheckError(err error) {
	if err != nil {
		log.Print("Connection Failed : ", err)
	}
}
