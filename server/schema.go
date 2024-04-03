package server

import ulid "github.com/oklog/ulid/v2"

type Person struct {
	Id        ulid.ULID
	FirstName string
	LastName  string
	Email     string
}
