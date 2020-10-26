package cacheparser

func facade(inputs []string) bool {
	return true
}

// Commands lists all the functions this cache would support
var Commands = map[string]func(arguments []string) bool{
	"SET":     facade,
	"SADD":    facade,
	"SETEX":   facade,
	"SETNX":   facade,
	"ZADD":    facade,
	"HSET":    facade,
	"LSET":    facade,
	"LPUSH":   facade,
	"MSET":    facade,
	"HMSET":   facade,
	"GET":     facade,
	"HGET":    facade,
	"HKEYS":   facade,
	"LPOP":    facade,
	"LINDEX":  facade,
	"GETSET":  facade,
	"HGETALL": facade,
	"HLEN":    facade,
	"HMGET":   facade,
	"PING":    facade,
	"HEXISTS": facade,
	"EXISTS":  facade,
	"LLEN":    facade,
	"MGET":    facade,
	"STRLEN":  facade,
	"ZCARD":   facade,
	"KEYS":    facade,
}
