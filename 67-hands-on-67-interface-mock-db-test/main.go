package main

import (
	"fmt"
	"log"
)

type user struct {
	id   int
	name string
}
type mockDataStore struct {
	users map[int]user
}

func (md mockDataStore) getUser(id int) (user, error) {
	value, ok := md.users[id]
	if !ok {
		return user{}, fmt.Errorf("User does not exist with ID: ", id)
	}
	return value, nil
}
func (md mockDataStore) saveUser(u user) error {
	_, ok := md.users[u.id]
	if ok {
		return fmt.Errorf("User already exist with ID: ", u.id)
	}
	md.users[u.id] = u
	return nil
}

type dataStore interface {
	getUser(id int) (user, error)
	saveUser(u user) error
}
type service struct {
	ds dataStore
}

func (s service) getUser(id int) (user, error) {
	return s.ds.getUser(id)
}

func (s service) saveUser(u user) error {
	return s.ds.saveUser(u)
}
func main() {
	db := mockDataStore{
		users: make(map[int]user),
	}

	srvc1 := service{
		ds: db,
	}
	user1 := user{
		id:   1,
		name: "zaman",
	}
	err := srvc1.saveUser(user1)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	u1Returned, err := srvc1.getUser(user1.id)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(user1)
	fmt.Println(u1Returned)
}
