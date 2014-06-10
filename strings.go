package Redis

import (
	"github.com/redigo/redis"
)
func (r *RedisType) Set (key string, value interface {}) (bool, error) {
	row,err := r.GetBool(r.RedisConn.Do("SET",key,value))
	return row, err
}

func (r *RedisType) Get (key string) (interface {},error) {
	row,err := redis.String(r.RedisConn.Do("GET",key))
	return row, err
}

func (r *RedisType) Incr (key string) (int, error){
	row,err := redis.Int(r.RedisConn.Do("INCR",key))
	return row, err
}

func (r *RedisType) IncrBy (key string, increment int) (int, error){
	row,err := redis.Int(r.RedisConn.Do("INCRBY",key,increment))
	return row, err
}

func (r *RedisType) Decr (key string) (int, error){
	row,err := redis.Int(r.RedisConn.Do("DECR",key))
	return row, err
}

func (r *RedisType) DecrBy (key string, decrement int) (int, error){
	row,err := redis.Int(r.RedisConn.Do("DECRBY",key,decrement))
	return row, err
}

func (r *RedisType) MSet (keys map[string]interface {}) (bool,error) {
	params := make([]interface {},0)
	for k,v := range keys {
		params = append(params,k)
		params = append(params,v)
	}
	row,err := r.GetBool(r.RedisConn.Do("MSET",params...))
	return row, err
}

func (r *RedisType) MGet (keys []string) ([]string, error) {
	params := make([]interface {},0)
	for _,v := range keys {
		params = append(params,v)
	}
	row,err := redis.Strings(r.RedisConn.Do("MGET",params...))
	return row, err
}

func (r *RedisType) SetEx (key string, value interface {}, seconds int) (bool,error) {
	row,err := r.GetBool(r.RedisConn.Do("SETEX",key,seconds,value))
	return row, err
}

