package mysql

import (

	"database/sql"
	"fmt"
	"os"
	"github.com/pressly/goose"

)

type Repository struct{

	db *sql.DB

}

func New(db *sql.DB) *Repository{

	migrations(db)

	return &Repository{
		db: db,
	}

}

func migrations(db *sql.DB){

	db, err := goose.OpenDBWithDriver("mysql", "root:test@tcp(127.0.0.1:33061)/base" )

	if err != nil{
		fmt.Println(err.Error())
	}

	// if err := goose.Run("up", db, os.Getenv("MIGRATIONS_DIR")); err != nil{
	// 	fmt.Println(err.Error())
	// }

	if err := goose.Up(db, os.Getenv("MIGRATIONS_DIR")); err != nil {
        fmt.Println(err.Error())
    }

	defer db.Close()
	// fmt.Println(migrations.Dir)

	// _, err := migrate.Exec(db, "mysql", migrations, migrate.Up)

	// if err != nil{
	// 	fmt.Println(err.Error())
	// }

}


