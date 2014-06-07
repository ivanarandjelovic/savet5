package savet

import (
	//	"fmt"
	. "aivan/savet5/db"
	"errors"
	"log"
	"time"
)

// User object
type Savet struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

func (_ Savet) TableName() string {
	return "Saveti"
}

//Try to find the user based on email and pass, return the user if found, nil otherwise
func List() ([]Savet, error) {
	var saveti []Savet
	if DB.Find(&saveti).RecordNotFound() {
		return saveti, errors.New("Saveti list error!")
	}
	log.Println("saveti: ", saveti)
	return saveti, nil
}

func Create(savet Savet) error {
	DB.Save(&savet)
	if savet.Id == 0 {
		return errors.New("Savet save failed!")
	}
	return nil
}

func Get(id int64) Savet {
	savet := Savet{Id: id}
	DB.Find(&savet)
	return savet
}
