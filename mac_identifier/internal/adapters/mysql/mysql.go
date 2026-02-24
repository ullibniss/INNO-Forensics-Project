package mysql

import (
	"database/sql"
	"fmt"
	"mac_identifier/pkg/config"
)

type MysqlRepo struct {
	db *sql.DB
}

func NewMysqlRepo() *MysqlRepo {
	cfg := config.LoadConfig()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME))
	if err != nil {
		panic(err.Error())
	}
	return &MysqlRepo{
		db: db,
	}
}
