package skeleton

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

// Tambahkan query di dalam const
// getAllUser = "GetAllUser"
// qGetAllUser = "SELECT * FROM users"
const ()

// Tambahkan query ke dalam key value order agar menjadi prepared statements
// readStmt = []statement{
// 	{getAllUser, qGetAllUser},
// }
var (
	readStmt = []statement{}
)

// New ...
func New(db *sqlx.DB) Data {
	d := Data{
		db: db,
	}

	d.initStmt()
	return d
}

func (d *Data) initStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize statement key %v, err : %v", v.key, err)
		}
	}

	d.stmt = stmts
}
