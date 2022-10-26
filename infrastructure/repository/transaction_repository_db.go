package repository

import "database/sql"

type TransactionRepositoryDb struct {
	db *sql.DB
}
