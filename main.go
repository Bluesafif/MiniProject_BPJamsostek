package main

import (
	"MiniProjectBPJamsostek/dao"
	"MiniProjectBPJamsostek/model"
	"MiniProjectBPJamsostek/router"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	dao.DBConnection()
	applicationModel.InitiateDefaultOperator()
	router.ApiController()
}
