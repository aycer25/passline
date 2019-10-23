package storage

import (
	"errors"
)

type Data struct {
	Items []Item
}

// Item structure
type Item struct {
	Name        string       `json:"name"`
	Credentials []Credential `json:"credentials"`
}

func (item Item) GetCredentialByUsername(username string) (Credential, error) {
	for i := 0; i < len(item.Credentials); i++ {
		if item.Credentials[i].Username == username {
			return item.Credentials[i], nil
		}
	}

	return Credential{}, errors.New("Not found")
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Storage interface
type Storage interface {
	Init() error
	GetByName(string) (Item, error)
	GetByIndex(int) (Item, error)
	GetAll() ([]Item, error)
	AddItem(Item) error
	AddCredential(string, Credential) error
	DeleteItem(Item) error
	DeleteCredential(Item, Credential) error
}

func removeFromItems(s []Item, i int) []Item {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func removeFromCredentials(s []Credential, i int) []Credential {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func getIndexOfItem(slice []Item, item Item) int {
	for p, v := range slice {
		if v.Name == item.Name {
			return p
		}
	}
	return -1
}

func getIndexOfCredential(slice []Credential, credential Credential) int {
	for p, v := range slice {
		if v == credential {
			return p
		}
	}
	return -1
}
