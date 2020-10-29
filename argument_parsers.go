package cacheparser

var creationInputTypes = map[string]string{
	"SET":   "key-value",
	"SADD":  "key-list",
	"SETEX": "key-value",
	"SETNX": "key-value",
	"ZADD":  "key-list",
	"HSET":  "hash",
	"LSET":  "key-list",
	"LPUSH": "key-list",
	"MSET":  "key-list",
	"HMSET": "hash",
}

var retrievalInputTypes = map[string]string{
	"GET":     "key",
	"HGET":    "key-value",
	"HKEYS":   "key-list",
	"LPOP":    "key",
	"LINDEX":  "key-value",
	"GETSET":  "key-value",
	"HGETALL": "key",
	"HLEN":    "key",
	"HMGET":   "key-list",
	"PING":    "key",
	"HEXISTS": "key-value",
	"EXISTS":  "key",
	"LLEN":    "key",
	"MGET":    "key-list",
	"STRLEN":  "key",
	"ZCARD":   "key",
	"KEYS":    "key-value",
}

func checkInputValidity(commandKey string, inputs []string) bool {

	key := inputs[0]
	println("key is", key)

	valuesAreValid := isValidInputValues(creationInputTypes[commandKey], inputs[1:])
	println(valuesAreValid)
	return true
}

func isValidInputValues(inputType string, inputValues []string) bool {
	if inputType == "key-value" {
		if len(inputValues) == 0 {
			return false
		}
	}

	if inputType == "key-list" {
		if len(inputValues) == 0 {
			return false
		}
	}

	if inputType == "hash" {
		// check for incomplete field value pair
		if len(inputValues)%2 == 1 {
			return false
		}
	}
	return true
}
