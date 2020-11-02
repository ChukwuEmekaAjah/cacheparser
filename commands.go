package cacheparser

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

// Commands lists all the functions this cache would support
var Commands = map[string]func(commandKey string, arguments []string) bool{
	"SET":     isValidSetCommand,
	"SADD":    checkInputValidity,
	"SETEX":   checkInputValidity,
	"SETNX":   checkInputValidity,
	"ZADD":    checkInputValidity,
	"HSET":    checkInputValidity,
	"LSET":    checkInputValidity,
	"LPUSH":   checkInputValidity,
	"MSET":    checkInputValidity,
	"HMSET":   checkInputValidity,
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
