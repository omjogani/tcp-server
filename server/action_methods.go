package server

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/fatih/color"
	"github.com/go-faker/faker/v4"
	ulid "github.com/oklog/ulid/v2"
)

var persons []Person

func init() {
	if len(persons) == 0 {
		for i := 0; i < 5; i++ {
			var person Person
			person.Id = ulid.Make()
			person.FirstName = faker.FirstName()
			person.LastName = faker.LastName()
			person.Email = faker.Email()

			persons = append(persons, person)
		}
	}
}

func GET(conn net.Conn) {
	data, err := json.Marshal(persons)
	handleError(err, "TCP_SERVER_ERROR: failed to marshal data.")
	_, err = conn.Write(data)
	handleError(err, "TCP_SERVER_ERROR: failed to write data.")
}

func readBody(conn net.Conn) []byte {
	body := make([]byte, (1024 * 10))
	bodyLength, err := conn.Read(body)
	handleError(err, "failed to read body")

	body = body[:bodyLength]
	if bodyLength == 0 {
		color.Red("TCP_SERVER_ERROR: failed to read body | ", err)
		return nil
	}
	return body
}

func POST(conn net.Conn, bodyType string) {
	body := readBody(conn)

	if strings.ToUpper(bodyType) == "JSON" {
		var data map[string]interface{}
		err := json.Unmarshal(body, &data)
		handleError(err, "TCP_SERVER_ERROR: failed to Unmarshal operation.")
		id := ulid.Make()
		data["Id"] = id.String()

		modifiedBody, err := json.Marshal(data)
		handleError(err, "TCP_SERVER_ERROR: failed to Marshal Complete Body")

		var person Person
		err = json.Unmarshal(modifiedBody, &person)
		handleError(err, "TCP_SERVER_ERROR: failed to convert body to json.")

		persons = append(persons, person)

		_, err = conn.Write([]byte(fmt.Sprintf(
			"%s %s %s %s - %s\n",
			person.Id, person.FirstName, person.LastName, person.Email, "SUCCESS!",
		)))
		handleError(err, "TCP_SERVER_ERROR: failed to update success status.")
	}
}
