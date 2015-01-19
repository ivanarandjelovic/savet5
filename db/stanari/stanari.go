package stanari

import (
	//	"fmt"
	"errors"
	. "github.com/ivanarandjelovic/savet5/db"
	"log"
)

// User object
type Stanar struct {
	Id        int64  `json:"id"`
	SavetId   int64  `json:"savetId"`
	BrojStana string `json:"brojStana"`
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
	if DB.Where("savet_id = ?", savetId).Find(&stanari).RecordNotFound() {
		return stanari, errors.New("Stanari list error!")
	}
	log.Println("stanari: ", stanari)
	return stanari, nil
}

func Create(stanar Stanar) error {
	DB.Save(&stanar)
	if stanar.Id == 0 {
		return errors.New("Stanar save failed!")
	}
	return nil
}
