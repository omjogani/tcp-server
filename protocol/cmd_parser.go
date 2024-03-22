package protocol

import (
	"log"
	"strings"
)

func ParseCommand(fullCommand string) (command, path, bodyType string) {
	breakdownCmd := strings.Split(fullCommand, " ")
	if len(breakdownCmd) < 2 {
		log.Fatal("PARSER_ERROR: Please provide sufficient params to command!")
		return "", "", ""
	}

	command = strings.TrimSpace(breakdownCmd[0])
	path = strings.TrimSpace(breakdownCmd[1])
	bodyType = strings.TrimSpace(breakdownCmd[2])

	if bodyType == "" {
		bodyType = "NORMAL"
	}

	return
}
