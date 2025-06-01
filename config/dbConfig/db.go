package dbConfig

import (
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"inventory/ent/entgen"
	"sync"
)

var (
	EntClient *entgen.Client
	once      sync.Once
)

func startDB() {
	db := GetDB()

	drv := entsql.OpenDB(dialect.MySQL, db)          // Wrap *sql.DB using MySQL driver
	EntClient = entgen.NewClient(entgen.Driver(drv)) // Pass driver to Ent client

}

func InitDB() *entgen.Client {
	once.Do(func() {
		startDB()
	})
	return EntClient
}
