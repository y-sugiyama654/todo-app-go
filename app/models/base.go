package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"todo-go-app/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	uuid STRING NOT NULL UNIQUE,
    	name STRING,
    	email STRING,
    	password STRING,
    	created_at DATETIME)`, tableNameUser)
	Db.Exec(cmdU)
}

func CreateUUID() (uuidObj uuid.UUID) {
	uuidObj, _ = uuid.NewUUID()
	return uuidObj
}

func Encrypt(plaintext string) (encrypted string) {
	encrypted = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return encrypted
}