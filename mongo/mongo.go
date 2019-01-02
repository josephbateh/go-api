package mongo

import (
	"errors"
	"hash/fnv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AddTeam in teams collection
func AddTeam(team Team) error {
	session, err := connect()
	if err != nil {
		return err
	}

	_, exists := GetOneTeam(team.ID)
	if exists == nil {
		err = errors.New("Team already exists")
	} else {
		c := session.DB("nfl").C("teams")
		err = c.Insert(team)
	}
	disconnect(session)
	return err
}

// GetOneTeam returns the team with the given name
func GetOneTeam(id uint32) (Team, error) {
	session, err := connect()
	result := Team{}

	if err != nil {
		return result, err
	}

	c := session.DB("nfl").C("teams")
	err = c.Find(bson.M{"id": id}).One(&result)
	disconnect(session)

	return result, err
}

// DeleteOneTeam deletes the team with the given ID and returns that ID if successful
func DeleteOneTeam(id uint32) (uint32, error) {
	session, err := connect()

	if err != nil {
		return 0, err
	}

	c := session.DB("nfl").C("teams")
	err = c.Remove(bson.M{"id": id})
	disconnect(session)

	return id, err
}

// GetAllTeams returns all teams in the database
func GetAllTeams() ([]Team, error) {
	session, err := connect()
	if err != nil {
		return nil, err
	}
	c := session.DB("nfl").C("teams")

	result := []Team{}
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		return nil, err
	}
	disconnect(session)

	return result, err
}

func connect() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://admin:password1@ds059135.mlab.com:59135/nfl")
	if err == nil {
		session.SetMode(mgo.Monotonic, true)
	}
	return session, err
}

func disconnect(session *mgo.Session) {
	session.Close()
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
