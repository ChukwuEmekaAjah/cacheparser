package cacheparser

import "strings"

//Parse is the top-level function that accepts client commands on the query
func Parse(command string) bool {
	commandParts := strings.Fields(command)
	commandKey := commandParts[0]
	commandParsingFunction, commandExists := Commands[strings.ToUpper(commandKey)]

	if commandExists == false {
		return false
	}

	validRequest := commandParsingFunction(commandKey, commandParts[1:])

	if !validRequest {
		return false
	}

	println("Correct parsing")
	return true
}
