package stanari

import (
	//	"fmt"
	. "aivan/savet5/db"
	"errors"
	"log"
)

// User object
type Stanar struct {
	Id        int64  `json:"id"`
	BrojStana string `json:"broj_Stana"`
	Redosled  int    `json:"redosled"`
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
}

func (_ Stanar) TableName() string {
	return "Stanari"
}

//Try to find the user based on email and pass, return the user if found, nil otherwise
func Get(savetId int64) ([]Stanar, error) {
	var stanari []Stanar
	if DB.Where("SAVET_ID = ?", savetId).Find(&stanari).RecordNotFound() {
		return stanari, errors.New("Stanari list error!")
	}
	log.Println("stanari: ", stanari)
	return stanari, nil
}
