package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const (
	dbName      = "trickortreat"
	candiesColl = "candies"
)

// wrapper object to ensure all data in mongo has a key
type objWrapper struct {
	Key  string `bson:"key"`
	Data Model  `bson:"data"`
}

// MongoDB is a DB implementation that talks to an external MongoDB server.
// Note: this is UNTESTED code, only to be used as a getting started reference.
// make sure you test it thouroughly against a real MongoDB server before using it
// in production
type MongoDB struct {
	sess *mgo.Session
}

// NewMongoDB connects to an external MongoDB server and returns a DB implementation
// that fronts that connection. call Close() on the returned value when done.
func NewMongoDB(urlStr string) (*MongoDB, error) {
	s, err := mgo.Dial(urlStr)
	if err != nil {
		return nil, err
	}
	return &MongoDB{sess: s}, nil
}

func (m *MongoDB) GetAllKeys() ([]string, error) {
	sess := m.sess.Copy()
	defer sess.Close()
	coll := sess.DB(dbName).C(candiesColl)
	wrapper := []objWrapper{}
	if err := coll.Find(bson.M{}).All(&wrapper); err != nil {
		return nil, err
	}
	ret := make([]string, len(wrapper))
	for i, obj := range wrapper {
		ret[i] = obj.Key
	}
	return ret, nil
}

// Get is the interface implementation
func (m *MongoDB) Get(key string, val Model) error {
	sess := m.sess.Copy()
	defer sess.Close()
	coll := sess.DB(dbName).C(candiesColl)
	wrapper := objWrapper{}
	if err := coll.Find(bson.M{"key": key}).One(&wrapper); err != nil {
		return err
	}
	val = wrapper.Data
	return nil
}

// Set is the interface implementation
func (m *MongoDB) Set(key string, val Model) error {
	sess := m.sess.Copy()
	defer sess.Close()
	coll := sess.DB(dbName).C(candiesColl)
	return coll.Update(bson.M{"key": key}, objWrapper{Key: key, Data: val})
}

// Upsert is the interface implementation
func (m *MongoDB) Upsert(key string, val Model) (bool, error) {
	sess := m.sess.Copy()
	defer sess.Close()
	coll := sess.DB(dbName).C(candiesColl)
	cInfo, err := coll.Upsert(bson.M{"key": key}, objWrapper{Key: key, Data: val})
	if err != nil {
		return false, err
	}
	// the Updated field is set when already existed, otherwise the UpsertedID field is set.
	// see the func at https://bazaar.launchpad.net/+branch/mgo/v2/view/head:/session.go#L1896
	return cInfo.UpsertedId != nil, nil
}

// Close releases the underlying connections. always call this when
// completely done with operations, not before.
func (m *MongoDB) Close() {
	m.sess.Close()
}
