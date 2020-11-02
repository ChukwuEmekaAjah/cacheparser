package cacheparser

import "strconv"

func isValidSetCommand(command string, arguments []string) bool {
	if len(arguments) < 2 {
		return false
	}

	return true
}

func isValidGetCommand(command string, arguments []string) bool {
	if len(arguments) < 1 {
		return false
	}

	return true
}

func isValidSetexCommand(command string, arguments []string) bool {
	if len(arguments) < 3 {
		return false
	}

	_, err := strconv.ParseInt(arguments[1], 10, 64)

	if err != nil {
		return false
	}

	return true
}

func isValidSaddCommand(command string, arguments []string) bool {
	if len(arguments) < 2 {
		return false
	}

	return true
}

func isValidSetnxCommand(command string, arguments []string) bool {
	if len(arguments) < 2 {
		return false
	}

	return true
}

func isValidZaddCommand(command string, arguments []string) bool {
	if len(arguments) < 3 {
		return false
	}

	// member without score is part of the arguments
	if len(arguments)%2 == 0 {
		return false
	}

	return true
}

func isValidHsetCommand(command string, arguments []string) bool {
	if len(arguments) != 3 {
		return false
	}

	return true
}

func isValidLsetCommand(command string, arguments []string) bool {
	if len(arguments) != 3 {
		return false
	}

	_, err := strconv.ParseInt(arguments[1], 10, 64)

	if err != nil {
		return false
	}

	return true
}

func isValidLpushCommand(command string, arguments []string) bool {
	if len(arguments) < 2 {
		return false
	}

	return true
}

func isValidMsetCommand(command string, arguments []string) bool {
	// should have an equal number of keys and values
	if len(arguments)%2 == 1 {
		return false
	}
	return true
}

func isValidHMSetCommand(command string, arguments []string) bool {
	// should have at least one fiel-value pair with key name
	if len(arguments) < 3 {
		return false
	}
	// field without value part of arguments
	if len(arguments)%2 == 0 {
		return false
	}

	return false
}

// Commands lists all the functions this cache would support
var Commands = map[string]func(commandKey string, arguments []string) bool{
	"SET":     isValidSetCommand,
	"SADD":    isValidSaddCommand,
	"SETEX":   isValidSetexCommand,
	"SETNX":   isValidSetnxCommand,
	"ZADD":    isValidZaddCommand,
	"HSET":    isValidHsetCommand,
	"LSET":    isValidLsetCommand,
	"LPUSH":   isValidLpushCommand,
	"MSET":    isValidMsetCommand,
	"HMSET":   isValidHMSetCommand,
	"GET":     isValidGetCommand,
	"HGET":    checkInputValidity,
	"HKEYS":   checkInputValidity,
	"LPOP":    checkInputValidity,
	"LINDEX":  checkInputValidity,
	"GETSET":  checkInputValidity,
	"HGETALL": checkInputValidity,
	"HLEN":    checkInputValidity,
	"HMGET":   checkInputValidity,
	"PING":    checkInputValidity,
	"HEXISTS": checkInputValidity,
	"EXISTS":  checkInputValidity,
	"LLEN":    checkInputValidity,
	"MGET":    checkInputValidity,
	"STRLEN":  checkInputValidity,
	"ZCARD":   checkInputValidity,
	"KEYS":    checkInputValidity,
}
