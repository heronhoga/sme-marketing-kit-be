package repositories

import (
	"context"
	"database/sql"

	"github.com/heronhoga/sme-marketing-kit-be/models"
)

type AuthRepository struct {
	dbconn *sql.DB
}

func NewAuthRepository(dbconn *sql.DB) *AuthRepository {
	return &AuthRepository{
		dbconn: dbconn,
	}
}

func (r *AuthRepository) InsertNewUser(c context.Context, userId string, registerRequest models.RegisterRequest) error {
	insertNewUserQuery := "INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)" 
	_, err := r.dbconn.Exec(insertNewUserQuery, userId, registerRequest.Name, registerRequest.Email, registerRequest.Password)
	if err != nil {
		print(err.Error())
		return err
	}

	return nil
}