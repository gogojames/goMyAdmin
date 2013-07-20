package gomyadmin

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

const (
	DRIVER_NAME      = "mysql"
	DATA_SOURCE_NAME = "root@/test?charset=utf8"
)

type Row struct{}
