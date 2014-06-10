package Redis

import (
	"github.com/redigo/redis"
	"errors"
)

// Добавление
func (r *RedisType) ZAdd (key string, members interface {}, scores interface {}) (bool, error) {
	switch members.(type){
	case string:
		params := make([]interface{}, 0)
		params = append(params, key)
		params = append(params, scores)
		params = append(params, members)
		row, err := r.GetBool(r.RedisConn.Do("ZADD", params...))
		return row, err
	case []string:
		params := make([]interface{}, 0)
		params = append(params, key)
		for k,v := range scores.([]string){
			params = append(params,v)
			params = append(params,members.([]string)[k])
		}
		row, err := r.GetBool(r.RedisConn.Do("ZADD", params...))
		return row, err
	default:
		return false,errors.New("")
	}
}

func (r *RedisType) ZCard (key string) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("ZCARD",key))
	if err == redis.ErrNil{
		return 0, nil
	}
	return row, err
}

func (r *RedisType) ZCount (key string, min,max int) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("ZCOUNT",key,min,max))
	if err == redis.ErrNil{
		return 0, nil
	}
	return row, err
}

func (r *RedisType) ZIncrBy (key,member string, increment int) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("ZINCRBY",key,increment,member))
	return row, err
}

func (r *RedisType) ZRange (key string, start,stop int, withScore bool) (interface {},error) {
	params := make([]interface {},0)
	params = append(params,key)
	params = append(params,start)
	params = append(params,stop)
	row, err := redis.Strings(r.RedisConn.Do("ZRANGE", params...))
	if err == redis.ErrNil{
		return nil, nil
	}
	if withScore{
		data := make(map[string]interface {})
		for _,key2 := range row {
			row2, err2 := r.ZScore(key, key2)
			if err2 != nil {
				data[key2] = nil
				continue
			} else { data[key2] = row2 } }
		return data, nil }
	return row, err
}

func (r *RedisType) ZScore (key, member string) (interface {},error) {
	params := make([]interface {},0)
	params = append(params,key)
	params = append(params,member)
	row,err := redis.Int(r.RedisConn.Do("ZSCORE",params...))
	if err == redis.ErrNil{
		return nil, nil
	}
	return row, err
}


func (r *RedisType) ZRevRange (key string, start,stop int, withScore bool) (interface {},error) {
	params := make([]interface {},0)
	params = append(params,key)
	params = append(params,start)
	params = append(params,stop)
	row, err := redis.Strings(r.RedisConn.Do("ZREVRANGE", params...))
	if err == redis.ErrNil{
		return nil, nil
	}
	if withScore{
		data := make(map[string]interface {})
		for _,key2 := range row {
			row2, err2 := r.ZScore(key, key2)
			if err2 != nil {
				data[key2] = nil
				continue
			} else { data[key2] = row2 } }
		return data, nil }
	return row, err
}

func (r *RedisType) ZRangeByScore (key string, min,max int, withScore bool) (interface {},error) {
	params := make([]interface {},0)
	params = append(params,key)
	params = append(params,min)
	params = append(params,max)
	row, err := redis.Strings(r.RedisConn.Do("ZRANGEBYSCORE", params...))
	if err == redis.ErrNil{
		return nil, nil
	}
	if withScore{
		data := make(map[string]interface {})
		for _,key2 := range row {
			row2, err2 := r.ZScore(key, key2)
			if err2 != nil {
				data[key2] = nil
				continue
			} else { data[key2] = row2 } }
		return data, nil }
	return row, err
}

// Возвращает позицию подключа в Z - списке
func (r *RedisType) ZRank (key, member string) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("ZRANK",key,member))
	if err == redis.ErrNil{
		return 0, nil
	}
	return row, err
}

// Удаления подключей из Z - списка, возвращает количество удалённых подключей
func (r *RedisType) ZRem (key string ,members ...string) (int,error) {
	params := make([]interface {},0)
	params = append(params,key)
	for _,v := range members{
		params = append(params,v)
	}
	row,err := redis.Int(r.RedisConn.Do("ZREM",params...))
	return row, err
}

func (r *RedisType) ZRemRangeByRank (key string,start, stop int) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("ZREMRANGEBYRANK",key,start,stop))
	return row, err
}
func (r *RedisType) ZRemRangeByScore (key string,min, max int) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("ZREMRANGEBYSCORE",key,min, max))
	return row, err
}
func (r *RedisType) ZRevRangeByScore (key string,min, max int, withScore bool) (interface {},error) {
	row,err := redis.Strings(r.RedisConn.Do("ZREVRANGEBYSCORE",key,min, max))
	if err == redis.ErrNil{
		return nil, nil
	}
	if withScore{
		data := make(map[string]interface {})
		for _,key2 := range row {
			row2, err2 := r.ZScore(key, key2)
			if err2 != nil {
				data[key2] = nil
				continue
			} else { data[key2] = row2 } }
		return data, nil }
	return row, err
}

func (r *RedisType) ZRevRank (key, member string) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("ZREVRANK",key,member))
	if err == redis.ErrNil{
		return 0, nil
	}
	return row, err
}
