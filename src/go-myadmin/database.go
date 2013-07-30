package gomyadmin

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Database struct {
	DB map[string]Db
}

func (d Database) GetTables(dname string) Db {
	return d.DB[dname]
}
