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
	if err == redis.ErrNil{ return 0, nil }
	return row, err
}

func (r *RedisType) ZCount (key string, min,max int) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("ZCOUNT",key,min,max))
	if err == redis.ErrNil{ return 0, nil }
	return row, err
}

func (r *RedisType) ZIncrBy (key string, member interface {}, increment interface {}) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("ZINCRBY",key,increment,member))
	return row, err
}
//ZINTERSTORE destination numkeys key
func (r *RedisType) ZInterStore (destination,key string, numkeys int,weight ...interface {}) (int,error) {
	if len(weight) > 0{
		row,err := redis.Int(r.RedisConn.Do("ZINTERSTORE",destination,key,numkeys,weight))
		return row, err
	}
	row,err := redis.Int(r.RedisConn.Do("ZINTERSTORE",destination,key,numkeys))
	return row, err
}
//zlexcount key min max
func (r *RedisType) ZLexCount (key string, min,max int) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("ZLEXCOUNT",key,min,max))
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

//ZRANGEBYLEX key min max [LIMIT offset count
func (r *RedisType) ZRangeByLex (key string, min,max int, other ...interface {}) (interface {},error) {
	params := make([]interface {},0)
	params = append(params,key)
	params = append(params,min)
	params = append(params,max)
	if len(other) > 0 && len(other) == 3 {
		for _,v := range other {
			params = append(params,v)
		}
	}
	row, err := redis.Strings(r.RedisConn.Do("ZRANGEBYLEX", params...))
	if err == redis.ErrNil{
		return nil, nil
	}
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

func (r *RedisType) ZRank (key, member string) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("ZRANK",key,member))
	if err == redis.ErrNil{
		return 0, nil
	}
	return row, err
}

func (r *RedisType) ZRem (key string ,members ...string) (int,error) {
	params := make([]interface {},0)
	params = append(params,key)
	for _,v := range members{
		params = append(params,v)
	}
	row,err := redis.Int(r.RedisConn.Do("ZREM",params...))
	return row, err
}

func (r *RedisType) ZRemRangeByLex (key string ,min,max int) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("ZREMRANGEBYLEX",key,min,max))
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

func (r *RedisType) ZUnionStore (destination string, numkeys int,key interface {}, weights ...interface {}) (interface {},error) {
	params := make([]interface {},0)
	params = append(params,destination)
	params = append(params,numkeys)
	switch key.(type){
	case string:
		params = append(params,key)
	case []string:
		for _,v := range key.([]string){
			params = append(params,v)
		}
	}
	if len(weights) > 0 {
		for _,v := range weights{
			params = append(params, v)
		}
	}
	row,err := redis.Int(r.RedisConn.Do("ZUNIONSTORE",params...))
	if err == redis.ErrNil{
		return nil, nil
	}
	return row, err
}

//// ZSCAN key cursor [MATCH pattern] [COUNT count]
//func (r *RedisType) ZScan (key string,start, stop int) (int,error) {
//	row,err := redis.Int(r.RedisConn.Do("ZREMRANGEBYRANK",key,start,stop))
//	return row, err
//}




