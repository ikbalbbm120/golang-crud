package config

import (
	"os"
	mgo "gopkg.in/mgo.v2"
)

func GetMongoDB() (*mgo.Database, error) {
	host := os.Getenv("mongo_host")
	dbName := os.Getenv("mongo_db_name")

	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)

	return db, nil
}