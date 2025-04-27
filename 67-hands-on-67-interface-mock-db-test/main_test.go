package main

import "testing"

func TestGetUser(t *testing.T) {
	md := &mockDataStore{
		users: map[int]user{
			2: {id: 2, name: "Jenny"},
		},
	}
	s := &service{
		ds: md,
	}
	u, err := s.getUser(2)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if u.name != "Jenny" {
		t.Errorf("got: %s, want: %s", u.name, "Jenny")
	}
}
