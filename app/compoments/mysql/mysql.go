package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var M *sql.DB

type DbConfig struct {
	Addr         string
	Username     string
	Password     string
	Name         string
	MaxConnNum   int
	MaxIdleConns int
}

func InitDB(cfg *DbConfig) error {
	connArgs := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Addr, cfg.Name)
	M, err := sql.Open("mysql", connArgs)
	if err != nil {
		return err
	}
	M.SetMaxOpenConns(cfg.MaxConnNum)
	M.SetMaxIdleConns(cfg.MaxIdleConns)

	return nil
}
