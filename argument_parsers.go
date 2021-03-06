package cacheparser

import (
	"errors"
	"fmt"
	"strconv"
)

// KeyValue represents how we would store values internally
type KeyValue struct {
	key     string
	value   []string
	command string
}

func parseSetFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {
	keyValueObject := new(KeyValue)
	keyValueObject.key = arguments[0]
	keyValueObject.value = arguments[1:]
	keyValueObject.command = command

	return keyValueObject
}

func parseSaddFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {

	keyValueObject, exists := cacheMap[arguments[0]]

	if exists == false {
		keyValueObject = new(KeyValue)
		keyValueObject.key = arguments[0]
		keyValueObject.value = arguments[1:]
		keyValueObject.command = command

		return keyValueObject
	}

	for _, newValue := range arguments[1:] {
		for _, value := range keyValueObject.value {
			if value == newValue { // value already in the set
				continue
			}
		}
		keyValueObject.value = append(keyValueObject.value, newValue)
	}

	return keyValueObject
}

func parseZaddFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {

	keyValueObject, exists := cacheMap[arguments[0]]

	if exists == false {
		keyValueObject = new(KeyValue)
		keyValueObject.key = arguments[0]
		keyValueObject.value = arguments[1:]
		keyValueObject.command = command

		return keyValueObject
	}

	for i, newValue := range arguments[1:] {
		for u, value := range keyValueObject.value {
			if i%2 == 0 && u%2 == 0 && value == newValue { // key-value pair already in the hash
				keyValueObject.value[u+1] = arguments[i+2]
				continue
			}
		}

		keyValueObject.value = append(keyValueObject.value, arguments[i+1:i+3]...) // add key-value pair to array
	}

	return keyValueObject
}

func parseLsetFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {

	keyValueObject, exists := cacheMap[arguments[0]]

	if exists == false {
		keyValueObject = new(KeyValue)
		keyValueObject.key = arguments[0]
		keyValueObject.value = arguments[1:]
		keyValueObject.command = command

		return keyValueObject
	}

	index, err := strconv.ParseInt(arguments[1], 10, 64)

	if err != nil {
		return keyValueObject
	}

	if index > int64(len(keyValueObject.value)) {
		keyValueObject.value = append(keyValueObject.value, arguments[2]) // insert value at position index
	} else {
		keyValueObject.value[index] = arguments[2]
	}

	return keyValueObject
}

func parseLpushFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {

	keyValueObject, exists := cacheMap[arguments[0]]

	if exists == false {
		keyValueObject = new(KeyValue)
		keyValueObject.key = arguments[0]
		keyValueObject.value = arguments[1:]
		keyValueObject.command = command

		return keyValueObject
	}

	keyValueObject.value = append(keyValueObject.value, arguments[1:]...)

	return keyValueObject
}

// ParserFunctions transforms arguments to a Struct
var ParserFunctions = map[string]func(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue{
	"SET":   parseSetFunction,
	"SADD":  parseSaddFunction,
	"SETEX": parseSetFunction,
	"SETNX": parseSetFunction,
	"ZADD":  parseZaddFunction,
	"HSET":  parseSetFunction,
	"LSET":  parseLsetFunction,
	"LPUSH": parseLpushFunction,
	"MSET":  parseSetFunction,
	"HMSET": parseZaddFunction,
}

func retrieveGetFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	keyValue, exists := cacheMap[arguments[0]]

	if exists == false {
		return "", errors.New("Unable to find value")
	}

	return keyValue.value[0], nil
}

func retrieveExistsFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	_, exists := cacheMap[arguments[0]]

	if exists == false {
		return "", errors.New("Unable to find key")
	}

	return "1", nil
}

func retrieveHlenFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	value, exists := cacheMap[arguments[0]]

	if exists == false {
		return "", errors.New("Unable to find key")
	}

	return fmt.Sprint(len(value.value) / 2), nil
}

func retrieveLlenFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	value, exists := cacheMap[arguments[0]]

	if exists == false {
		return "", errors.New("Unable to find key")
	}

	return fmt.Sprint(len(value.value)), nil
}

func retrieveStrlenFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	value, exists := cacheMap[arguments[0]]

	if exists == false {
		return "", errors.New("Unable to find key")
	}

	return fmt.Sprint(len(value.value[0])), nil
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
	"HLEN":    retrieveHlenFunction,
	"HMGET":   retrieveGetFunction,
	"PING":    retrieveGetFunction,
	"HEXISTS": retrieveGetFunction,
	"EXISTS":  retrieveExistsFunction,
	"LLEN":    retrieveLlenFunction,
	"MGET":    retrieveGetFunction,
	"STRLEN":  retrieveStrlenFunction,
	"ZCARD":   retrieveGetFunction,
	"KEYS":    retrieveGetFunction,
}
