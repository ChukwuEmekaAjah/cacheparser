package cacheparser

import (
	"errors"
)

// KeyValue represents how we would store values internally
type KeyValue struct {
	key       string
	value     []string
	valueType string
	command   string
}

func parseSetFunction(command string, arguments []string, valueType string) *KeyValue {
	keyValueObject := new(KeyValue)
	keyValueObject.key = arguments[0]
	keyValueObject.value = arguments[1:]
	keyValueObject.valueType = valueType
	keyValueObject.command = command

	return keyValueObject
}

// ParserFunctions transforms arguments to a Struct
var ParserFunctions = map[string]func(command string, arguments []string, valueType string) *KeyValue{
	"SET":   parseSetFunction,
	"SADD":  parseSetFunction,
	"SETEX": parseSetFunction,
	"SETNX": parseSetFunction,
	"ZADD":  parseSetFunction,
	"HSET":  parseSetFunction,
	"LSET":  parseSetFunction,
	"LPUSH": parseSetFunction,
	"MSET":  parseSetFunction,
	"HMSET": parseSetFunction,
}

func retrieveGetFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	keyValue, exists := cacheMap[arguments[0]]

	if exists == false {
		return "", errors.New("Unable to find value")
	}

	return keyValue.value[0], nil
}

// RetrievalFunctions lists all the functions this cache would support for retrieving values
var RetrievalFunctions = map[string]func(commandKey string, arguments []string, cacheMap map[string]*KeyValue) (string, error){
	"GET":     retrieveGetFunction,
	"HGET":    retrieveGetFunction,
	"HKEYS":   retrieveGetFunction,
	"LPOP":    retrieveGetFunction,
	"LINDEX":  retrieveGetFunction,
	"GETSET":  retrieveGetFunction,
	"HGETALL": retrieveGetFunction,
	"HLEN":    retrieveGetFunction,
	"HMGET":   retrieveGetFunction,
	"PING":    retrieveGetFunction,
	"HEXISTS": retrieveGetFunction,
	"EXISTS":  retrieveGetFunction,
	"LLEN":    retrieveGetFunction,
	"MGET":    retrieveGetFunction,
	"STRLEN":  retrieveGetFunction,
	"ZCARD":   retrieveGetFunction,
	"KEYS":    retrieveGetFunction,
}
