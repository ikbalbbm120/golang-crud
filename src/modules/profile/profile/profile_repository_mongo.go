package profile

import (
	"time"

	"github.com/ikbalbbm120/golang crud/src/modules/profile/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//profileRepositoryMongo

type profileRepositoryMongo struct {
	db *mgo.Database
	collection string
}

//NewProfileRepositoryMongo
func NewProfileRepositoryMongo(db *mgo.Database, collection string) *profileRepositoryMongo {
	return &profileRepositoryMongo {
		db: db,
		collection: collection,
	}
}

//save
func (r *profileRepositoryMongo)save(profile *model.profile) error {
	err := r.db.c(r.collection).insert(profile)
	return err
}

//delete
func (r *profileRepositoryMongo) Delete(id string) error{
	err := r.db.C(r.collection).Remove(bson.M{"id": id})
	return err
}

////FindByID

func (r *profileRepositoryMongo) FindByID(id string) (*model.Profile, error){
	var profile model.Profile

	err := r.db.C(r.collection).Find(bson.M{"id": id}).One(&profile)

	if err != nil {
	return nil, err
	}

	return &profile, nil
}

//FindAll
func (r *profileRepositoryMongo) FindAll() (model.Profiles, error){
	var profiles model.Profiles

	err := r.db.C(r.collection).Find(bson.M{}).All(&profiles)

	if err != nil {
	return nil, err
	}

	return profiles, nil
}