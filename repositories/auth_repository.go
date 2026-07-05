package repositories

import (
	"context"
	"database/sql"
	"time"

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

func (r *AuthRepository) FindUserPasswordByEmail(c context.Context, email string) (string, string, error) {
	var userPassword string
	var userId string
	findUserByIdQuery := "SELECT password, id FROM users WHERE email = ? LIMIT 1"
	err := r.dbconn.QueryRow(findUserByIdQuery, email).Scan(&userPassword, &userId)
	if err != nil {
		return "", "", err
	}
	return userPassword, userId, nil
}

func (r *AuthRepository) SaveRefreshToken(c context.Context,id string, userId string, hashedRefreshToken string, expiresAt time.Time) error {
	saveRefreshTokenQuery := "INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at) VALUES (?,?,?,?)"
	_, err := r.dbconn.Exec(saveRefreshTokenQuery, id, userId, hashedRefreshToken, expiresAt)
	if err != nil {
		return err
	}

	return nil
}
