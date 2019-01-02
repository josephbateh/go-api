package api

import (
	"encoding/json"
	"errors"
	"hash/fnv"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/josephbateh/go-api/log"
	"github.com/josephbateh/go-api/mongo"
	"github.com/josephbateh/go-api/verbs"
)

// Teams handles the endpoint /teams
func Teams(response http.ResponseWriter, request *http.Request) {
	log.Info(request.Method + " /teams")
	switch request.Method {
	case http.MethodGet:
		get(response, request)
	case http.MethodPost:
		post(response, request)
	case http.MethodPut:
		put(response, request)
	case http.MethodPatch:
		patch(response, request)
	case http.MethodDelete:
		delete(response, request)
	default:
		http.Error(response, http.StatusText(405), 405)
		log.Error("405 - Method not allowed")
		return
	}
}

func get(response http.ResponseWriter, request *http.Request) {
	var id uint32

	idString, err := parseID(request.RequestURI)
	if err != nil {
		log.Error(err.Error())
		http.Error(response, http.StatusText(500), 500)
	}

	if idString != "" {
		idInt64, err := strconv.ParseInt(idString, 10, 64)

		if err != nil {
			log.Error(err.Error())
			http.Error(response, http.StatusText(500), 500)
		}

		id = uint32(idInt64)
		team, err := mongo.GetOneTeam(id)

		if err != nil {
			log.Error(err.Error())
			http.Error(response, http.StatusText(500), 500)
		}
		verbs.Get(response, request, team)
	} else {
		teams, err := mongo.GetAllTeams()
		if err != nil {
			log.Error(err.Error())
			http.Error(response, http.StatusText(500), 500)
		}
		verbs.Get(response, request, teams)
	}
}

func post(response http.ResponseWriter, request *http.Request) {
	responseString := "Added team"

	// Parse JSON into byte array
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Error(err.Error())
	}

	// Unmarshal byte array into Team struct
	var team mongo.Team
	err = json.Unmarshal(body, &team)
	if err != nil {
		log.Error(err.Error())
	}

	team.ID = hash(team.Name)

	// Add team to database
	err = mongo.AddTeam(team)
	if err != nil {
		log.Error(err.Error())
		responseString = err.Error()
	}

	responseString = strconv.FormatInt(int64(team.ID), 10)
	verbs.Post(response, request, responseString)
}

func put(response http.ResponseWriter, request *http.Request) {
	var id uint32

	idString, err := parseID(request.RequestURI)
	if err != nil {
		log.Error(err.Error())
		http.Error(response, http.StatusText(500), 500)
		return
	}

	idInt64, err := strconv.ParseInt(idString, 10, 64)

	if err != nil {
		log.Error(err.Error())
		http.Error(response, http.StatusText(500), 500)
	}

	id = uint32(idInt64)

	// Parse JSON into byte array
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Error(err.Error())
	}

	// Unmarshal byte array into Team struct
	var team mongo.Team
	err = json.Unmarshal(body, &team)
	if err != nil {
		log.Error(err.Error())
	}

	team.ID = hash(team.Name)

	// Add team to database
	err = mongo.AddTeam(team)
	if err != nil {
		mongo.DeleteOneTeam(team.ID)
		newErr := mongo.AddTeam(team)
		if newErr != nil {
			log.Error(newErr.Error())
			http.Error(response, http.StatusText(500), 500)
		}
	}

	verbs.Put(response, request, id)
}

func patch(response http.ResponseWriter, request *http.Request) {
	var id uint32

	idString, err := parseID(request.RequestURI)
	if err != nil {
		log.Error(err.Error())
		http.Error(response, http.StatusText(500), 500)
		return
	}

	idInt64, err := strconv.ParseInt(idString, 10, 64)

	if err != nil {
		log.Error(err.Error())
		http.Error(response, http.StatusText(500), 500)
		return
	}

	id = uint32(idInt64)

	// Parse JSON into byte array
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Error(err.Error())
	}

	// Unmarshal byte array into Team struct
	var team mongo.Team
	err = json.Unmarshal(body, &team)
	if err != nil {
		log.Error(err.Error())
	}

	oldTeam, err := mongo.GetOneTeam(id)

	if err != nil {
		http.Error(response, http.StatusText(500), 500)
		return
	}

	// Replace values if not nil
	// Its a bunch of ifs... its bad I know
	if team.HeadCoach != "" {
		oldTeam.HeadCoach = team.HeadCoach
	}

	if team.Name != "" {
		oldTeam.Name = team.Name
	}

	if team.Record != "" {
		oldTeam.Record = team.Record
	}

	// Add team to database
	_, err = mongo.DeleteOneTeam(id)
	if err != nil {
		log.Error(err.Error())
		http.Error(response, http.StatusText(500), 500)
		return
	}

	err = mongo.AddTeam(oldTeam)
	if err != nil {
		log.Error(err.Error())
		http.Error(response, http.StatusText(500), 500)
		return
	}

	verbs.Patch(response, request, id)
}

func delete(response http.ResponseWriter, request *http.Request) {
	var id uint32

	idString, err := parseID(request.RequestURI)
	if err != nil {
		log.Error(err.Error())
		http.Error(response, http.StatusText(500), 500)
		return
	}

	if idString != "" {
		idInt64, err := strconv.ParseInt(idString, 10, 64)

		if err != nil {
			log.Error(err.Error())
			http.Error(response, http.StatusText(500), 500)
			return
		}

		id = uint32(idInt64)
		_, err = mongo.DeleteOneTeam(id)
		if err != nil {
			log.Error("Attempted to delete a non-existent team")
			http.Error(response, http.StatusText(500), 500)
			return
		}
		verbs.Delete(response, request, id)
	} else {
		http.Error(response, http.StatusText(400), 400)
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func parseID(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", errors.New("failed to parse URL")
	}
	q := u.Query()

	return q.Get("id"), nil
}
