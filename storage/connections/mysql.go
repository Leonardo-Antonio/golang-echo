package connections

import (
	"database/sql"
	"log"
	"sync"

	// !!
	_ "github.com/go-sql-driver/mysql"
)

var (
	once sync.Once
)

// Mysql Connection BD
func Mysql() *sql.DB {
	var conDB *sql.DB
	once.Do(func() {
		db, err := sql.Open("mysql", "leo:chester@tcp(localhost:3306)/BD_GO_ECHO")
		if err != nil {
			log.Fatalf("Hubo problemas con la BD -> %+v", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("Hubo problemas en el Ping -> %+v", err)
		}
		conDB = db
	})
	return conDB
}
