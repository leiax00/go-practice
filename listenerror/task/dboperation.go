package task

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

var (
	ctx context.Context
	db  *sql.DB
)

func init()  {
	db, err := sql.Open("MariaDB", "root:root@192.168.10.10/go-practice?charset=utf8");
	if err != nil {
		fmt.Printf("Failed to init db: %+v", errors.WithStack(err))
		os.Exit(0)
	}

	initData()
}

func QueryData()  {

}

func initData() {
	_, _ = db.ExecContext(ctx, "truncate go_user if exists")

}