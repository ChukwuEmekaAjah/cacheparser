package cacheparser

import "fmt"

// KeyValue represents how we would store values internally
type KeyValue struct {
	key       string
	value     []string
	valueType string
}

func parseSetFunction(arguments []string, valueType string) *KeyValue {
	keyValueObject := new(KeyValue)
	keyValueObject.key = arguments[0]
	keyValueObject.value = arguments[1:]
	keyValueObject.valueType = valueType

	fmt.Printf("value is %v", keyValueObject)
	return keyValueObject
}

// ParserFunctions transforms arguments to a Struct
var ParserFunctions = map[string]func(arguments []string, valueType string) *KeyValue{
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
