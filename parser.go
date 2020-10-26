package cacheparser

import "strings"

//Parse is the top-level function that accepts client commands on the query
func Parse(command string) bool {
	commandParts := strings.Fields(command)
	commandParsingFunction, commandExists := Commands[strings.ToUpper(commandParts[0])]

	if commandExists == false {
		return false
	}

	validRequest := commandParsingFunction(commandParts[1:])

	if !validRequest {
		return false
	}

	println("Correct parsing")
	return true
}
