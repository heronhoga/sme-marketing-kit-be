package repositories

import "database/sql"

type AuthRepository struct {
	dbconn *sql.DB
}

func NewAuthRepository(dbconn *sql.DB) *AuthRepository {
	return &AuthRepository{
		dbconn: dbconn,
	}
}