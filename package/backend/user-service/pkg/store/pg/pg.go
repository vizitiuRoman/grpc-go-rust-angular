package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/user-service/config"
)

const schema = `
		create table if not exists users (
			id serial primary key,
			email varchar(255) NOT NULL UNIQUE,
			password varchar(255) NOT NULL,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);
	`

func Dial() (*sqlx.DB, error) {
	DBSpec := fmt.Sprintf(
		"user=%s dbname=%s password=%s port=%s host=%s sslmode=disable",
		config.Get().DBUser, config.Get().DBName,
		config.Get().DBPassword, config.Get().DBPort,
		config.Get().DBHost,
	)
	database, err := sqlx.Connect(config.Get().DBDriver, DBSpec)
	if err != nil {
		return &sqlx.DB{}, err
	}

	err = database.Ping()
	if err != nil {
		return &sqlx.DB{}, err
	}

	database.MustExec(schema)
	return database, nil
}
