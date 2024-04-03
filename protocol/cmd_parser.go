package protocol

import (
	"log"
	"strings"
)

func ParseCommand(fullCommand string) (command, path, bodyType string) {
	breakdownCmd := strings.Split(fullCommand, " ")
	bodyTypeVar := "NORMAL"
	if len(breakdownCmd) < 2 {
		log.Fatal("PARSER_ERROR: Please provide sufficient params to command!")
		return "", "", ""
	} else if len(breakdownCmd) >= 3 {
		bodyTypeVar = strings.TrimSpace(breakdownCmd[2])
	}

	command = strings.TrimSpace(breakdownCmd[0])
	path = strings.TrimSpace(breakdownCmd[1])
	bodyType = bodyTypeVar
	return
}
