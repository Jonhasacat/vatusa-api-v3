package database

import (
	"database/sql"
	"fmt"
	"github.com/VATUSA/api-v3/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DSN(c *config.DatabaseConfig) (string, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.User, c.Password, c.Host, c.Port, c.Database)
	return dsn, nil
}

func Connect() error {
	dsn, err := DSN(&config.PrimaryDatabaseConfig)
	if err != nil {
		return err
	}
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	DB, err = gorm.Open(mysql.New(mysql.Config{Conn: conn}), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
