package database

import (
	"encoding/json"

	"github.com/decadevs/shoparena/models"
	"github.com/lib/pq"
	//"github.com/globalsign/mgo"
	//"github.com/jackc/pgx/v4"
)

// PostgresDB implements the DB interface
type PostgresDB struct {
	DB *pq.Driver
}

// UpdateUser updates user in the collection
func (pq *PostgresDB) UpdateUser(user *models.User) error {
	return pq.DB.C("user").Update(json.M{"email": user.Email}, user)
}

type DB interface{
	UpdateUser(user *models.User) error
}