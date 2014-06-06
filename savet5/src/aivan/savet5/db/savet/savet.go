package savet

import (
	//	"fmt"
	. "aivan/savet5/db"
	"log"
	"errors"
)

// User object
type Savet struct {
	Id          uint64 `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	CreatedAt   string `json:"createdAt"`
}

func (s Savet) TableName() string {
	return "Saveti"
}

//Try to find the user based on email and pass, return the user if found, nil otherwise
func List() ([]Savet, error) {
	var saveti[] Savet
	if DB.Find(&saveti).RecordNotFound() {
		return saveti, errors.New("Saveti list error!")
	}
	log.Println("saveti: ", saveti)
	return saveti, nil
}
