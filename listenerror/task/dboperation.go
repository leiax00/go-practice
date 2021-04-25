package task

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"os"
)

var (
	db *sql.DB
)

func init() {
	tmp, err := sql.Open("mysql", "root:root@tcp(192.168.10.10:3306)/go-practice?charset=utf8")
	if err != nil {
		fmt.Printf("Failed to init db: %+v", errors.WithStack(err))
		os.Exit(0)
	}
	db = tmp
	if err := db.Ping(); err != nil {
		fmt.Printf("ping failed: %+v", errors.WithStack(err))
		os.Exit(0)
	}
	initData()
}

func QueryData(userId int) (*User, error) {
	var (
		id         int
		name, desc string
		age        int8
	)
	err := db.QueryRow("SELECT * FROM go_user where id = ?", userId).Scan(&id, &name, &age, &desc)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, errors.Wrap(err, "Failed to query data.")
	default:
		return &User{
			id, name, age, desc,
		}, nil
	}
}

func CloseDb() error {
	return db.Close()
}

func initData() {
	_, _ = db.Query("create table if not exists go_user (id int auto_increment primary key , name varchar(64) default '-' not null, age smallint null, `desc` text null, constraint table_name_id_uindex unique (id)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;")
	_, _ = db.Query("truncate go_user")
	for i := 0; i < 10; i++ {
		_, _ = db.Query("insert into go_user (name, age, `desc`) VALUES (?, ?, ?)",
			fmt.Sprintf("user%02d", i),
			int8(18+i),
			fmt.Sprintf("This is user%02d", i),
		)
	}

}
