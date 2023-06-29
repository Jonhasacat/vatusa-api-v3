package legacydb

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"vatusa-api-v3/config"
)

var DB *gorm.DB

func DSN(c *config.DatabaseConfig) (string, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.User, c.Password, c.Host, c.Port, c.Database)
	return dsn, nil
}

func Connect() error {
	dsn, err := DSN(&config.LegacyDatabaseConfig)
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
