package repositories

import (
	"OrderPick/models"
	"time"

	"github.com/gocql/gocql"
)

type UserRepository struct {
	session *gocql.Session
}

func NewUserRepository(session *gocql.Session) *UserRepository {
	return &UserRepository{session}
}

func (r *UserRepository) CreateUser(user models.User) error {
	query := "INSERT INTO user (user_id, email, phone, password, created_at, updated_at, first_name, last_name, avatar, token_, refresh_token) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	return r.session.Query(query, user.User_id, *user.Email, *user.Phone, *user.Password, user.Created_at, user.Updated_at, *user.First_name, *user.Last_name, user.Avatar, user.Token, user.Refresh_Token).Exec()
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := "SELECT * FROM user WHERE email = ? LIMIT 1"
	err := r.session.Query(query, email).Consistency(gocql.One).Scan(
		&user.User_id, &user.Avatar, &user.Created_at, &user.Email, &user.First_name, &user.Last_name, &user.Password, &user.Phone, &user.Refresh_Token, &user.Token, &user.Updated_at)
	return user, err
}

func (r *UserRepository) GetUserById(userId string) (models.User, error) {
	var user models.User
	query := "SELECT * FROM user WHERE user_id = ? LIMIT 1"
	err := r.session.Query(query, userId).Consistency(gocql.One).Scan(
		&user.User_id, &user.Email, &user.Phone, &user.Password, &user.Created_at, &user.Updated_at, &user.First_name, &user.Last_name, &user.Avatar, &user.Token, &user.Refresh_Token,
	)
	return user, err
}

func (r *UserRepository) GetUsers(recordPerPage, startIndex int) ([]models.User, error) {
	query := "SELECT * FROM user LIMIT ? OFFSET ?"
	iter := r.session.Query(query, recordPerPage, startIndex).Iter()

	var users []models.User
	var user models.User
	for iter.Scan(&user.User_id, &user.Email, &user.Phone, &user.Password, &user.Created_at, &user.Updated_at, &user.First_name, &user.Last_name, &user.Avatar, &user.Token, &user.Refresh_Token) {
		users = append(users, user)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) UpdateTokens(userId, token, refreshToken string, updatedAt time.Time) error {
	query := "UPDATE user SET token = ?, refresh_token = ?, updated_at = ? WHERE user_id = ?"
	return r.session.Query(query, token, refreshToken, updatedAt, userId).Exec()
}
