package gomyadmin

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Db struct {
	Tables map[string]Table
}

func (d Db) GetTables(tname string) Table {
	return d.Tables[tname]
}

func init() {
	/*db, err := sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	db.Exec("use test")
	rows, err := db.Query("show tables")
	for rows.Next() {
		var t string
		rows.Scan(&t)
		Db.Tables[t] = Table.GetRows(0)
	}*/
}
