package Redis

import (
	"github.com/redigo/redis"
	"fmt"
)

// Работа с HASH (HSET)
func (r *RedisType) HGetAll(key string) (map[string]interface {},error) {
	rowRed,errRed := r.RedisConn.Do("HGETALL",key)
	if errRed == redis.ErrNil{
		return nil,errRed
	}
	row := r.GetRedisReply(rowRed,errRed,[]string{})
	return row,errRed
}

func (r *RedisType) HMGet(key string,fields ...interface {}) (map[string]interface {},error) {
	params := make([]interface {},0)
	vals := make([]string,0)
	params = append(params,key)
	for _,k := range fields{
		params = append(params,k)
		vals = append(vals,fmt.Sprintf("%s",k))
	}
	rowRed,errRed := r.RedisConn.Do("HMGET",params...)
	if errRed == redis.ErrNil{
		return nil,nil
	}
	row := r.GetRedisReply(rowRed,errRed,vals)
	return row,errRed
}

func (r *RedisType) HGet(key string,field interface {}) (interface {},error) {
	row,err := redis.String(r.RedisConn.Do("HGET",key,field))
	if err == redis.ErrNil{
		return nil,nil
	}
	return row,err
}

func (r *RedisType) HSet (key string,field interface {},value interface {}) error {
	params := make([]interface {},0)
	params = append(params,key)
	params = append(params,field)
	params = append(params,value)
	err := r.RedisConn.Send("HSET",params...)
	if err != nil {
		return err
	}
	return nil
}
func (r *RedisType) HMSet(key string, data map[string]interface {}) error {
	params := make([]interface {},0)
	params = append(params,key)
	for k,v := range data {
		params = append(params,k)
		params = append(params,v)
	}
	err := r.RedisConn.Send("HMSET",params...)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisType) HExists(key string,field interface {}) (bool,error) {
	params := make([]interface {},0)
	params = append(params,key)
	params = append(params,field)
	exist,err := redis.Bool(r.RedisConn.Do("HEXISTS",params...))
	return exist, err
}

func (r *RedisType) HDel(key string,fields ...interface {}) error {
	params := make([]interface {},0)
	params = append(params,key)
	for _,v := range fields {
		params = append(params, v)
	}
	err := r.RedisConn.Send("HDEL",params...)
	return err
}

func (r *RedisType) HKeys(key string) ([]string,error) {
	row,err := redis.Strings(r.RedisConn.Do("HKEYS",key))
	if err == redis.ErrNil{
		return nil,nil
	}
	return row,err
}

func (r *RedisType) HLen(key string) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("HLEN",key))
	if err == redis.ErrNil{
		return 0,nil
	}
	return row, err
}

func (r *RedisType) HIncrBy(key string,field interface {},increment interface {}) (int, error) {
	params := make([]interface {},0)
	params = append(params,key)
	params = append(params,field)
	params = append(params,increment)
	row,err := redis.Int(r.RedisConn.Do("HINCRBY",params...))
	return row,err
}

func (r *RedisType) HIncrByFloat (key string,field interface {},increment interface {}) (interface {}, error) {
	params := make([]interface {},0)
	params = append(params,key)
	params = append(params,field)
	params = append(params,increment)
	row,err := redis.Float64(r.RedisConn.Do("HINCRBYFLOAT",params...))
	return row,err
}

func (r *RedisType) HSetNx (key string,field interface {},value interface {}) (interface{},error) {
	params := make([]interface {},0)
	params = append(params,key)
	params = append(params,field)
	params = append(params,value)
	row,err := r.RedisConn.Do("HSETNX",params...)
	return row, err
}

func (r *RedisType) HVals (key string) (interface {},error) {
	row,err := r.RedisConn.Do("HVALS",key)
	if err == redis.ErrNil{
		return nil,nil
	}
	row2:= r.GetRedisReplyArray(row,err)
	return row2, err
}

//func (r *RedisType) HScan () () {
//	//HSCAN key cursor [MATCH pattern] [COUNT count]
//}
