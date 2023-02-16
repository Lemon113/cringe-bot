package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

type DB struct {
	Triggers Triggers `json:"Triggers"`
	Pharases Phrases  `json:"Phrases"`
	l        *log.Logger
}

func NewDB(l *log.Logger) *DB {
	file, err := ioutil.ReadFile("./db.json")
	var db DB

	if err != nil {
		l.Println(err.Error())
		return nil
	}

	l.Println(string(file))
	err = json.Unmarshal(file, &db)

	if err != nil {
		l.Println(err.Error())
		return nil
	}

	db.l = l
	return &db
}

type Triggers []string
type Phrases []string

func (db *DB) GenerateRandomResponse() string {
	n := len(db.Pharases)
	i := rand.Intn(n)
	s := db.Pharases[i]
	return s
}

func (db *DB) HasAnyTriggerWords(trigger string) bool {
	for _, w := range db.Triggers {
		if strings.Contains(strings.ToLower(trigger), strings.ToLower(w)) {
			return true
		}
	}

	return false
}
