package Redis

import "github.com/redigo/redis"
func (r *RedisType) Set (key string, value interface {}) (bool, error) {
	row,err := redis.Bool(r.RedisConn.Do("SET",key,value))
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
func (r *RedisType) IncrByFloat (key string, increment float32) (string, error){
	row,err := redis.String(r.RedisConn.Do("INCRBY",key,increment))
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
	row,err := redis.Bool(r.RedisConn.Do("MSET",params...))
	return row, err
}

func (r *RedisType) MSetNX (keys map[string]interface {}) (bool,error) {
	params := make([]interface {},0)
	for k,v := range keys {
		params = append(params,k)
		params = append(params,v)
	}
	row,err := redis.Bool(r.RedisConn.Do("MSETNX",params...))
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
	row,err := redis.Bool(r.RedisConn.Do("SETEX",key,seconds,value))
	return row, err
}

func (r *RedisType) Append (key string, value interface {}, seconds int) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("APPEND",key,seconds,value))
	return row, err
}

func (r *RedisType) BitCount (key string, borders ...int) (int,error) {
	if len(borders) > 0 {
		if len(borders) > 1 {
			return nil, errors.New("Too many borders")
		}
		if len(borders[0]) != 2 {
			return nil, errors.New("Too many borders")
		}
		row, err := redis.Int(r.RedisConn.Do("BITCOUNT", key, borders[0], borders[1]))
		return row, err
	} else {
		row, err := redis.Int(r.RedisConn.Do("BITCOUNT", key))
		return row, err
	}
}

func (r *RedisType) BitOp (operation,destkey,key string) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("BITOP",operation,destkey,key))
	return row, err
}

func (r *RedisType) BitPos (key string, bit int, borders ...int) (int,error) {
	if len(borders) > 0 {
		if len(borders) > 1 {
			return nil, errors.New("Too many borders")
		}
		if len(borders[0]) != 2 {
			return nil, errors.New("Too many borders")
		}
		row, err := redis.Int(r.RedisConn.Do("BITPOS", key, bit, borders[0], borders[1]))
		return row, err
	} else {
		row, err := redis.Int(r.RedisConn.Do("BITPOS", key, bit))
		return row, err
	}
}

func (r *RedisType) GetBit (key string, offset int) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("GETBIT",key,offset))
	return row, err
}

func (r *RedisType) GetRange (key string, start,end int) (string,error) {
	row,err := redis.Int(r.RedisConn.Do("GETRANGE",key,start,end))
	return row, err
}

func (r *RedisType) GetSet (key string, value interface {}) (string,error) {
	row,err := redis.String(r.RedisConn.Do("SETEX",key,value))
	return row, err
}
//
func (r *RedisType) SetEx (key string, value interface {}, seconds int) (bool,error) {
	row,err := r.GetBool(r.RedisConn.Do("SETEX",key,seconds,value))
	return row, err
}

func (r *RedisType) PSetEx (key string, value interface {}, milliseconds int) (bool,error) {
	row,err := r.GetBool(r.RedisConn.Do("PSETEX",key,milliseconds,value))
	return row, err
}

func (r *RedisType) SetBit (key string, offset int, value interface {}) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("SETBIT",key,offset,value))
	return row, err
}

func (r *RedisType) SetNX (key string, value interface {}) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("SETNX",key,value))
	return row, err
}

func (r *RedisType) SetRange (key string, offset int, value interface {}) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("SETRANGE",key,offset,value))
	return row, err
}
// STRLEN key
func (r *RedisType) StrLen (key string) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("STRLEN",key))
	return row, err
}
//
//
//
//func (r *RedisType) SetEx (key string, value interface {}, seconds int) (bool,error) {
//	row,err := r.GetBool(r.RedisConn.Do("SETEX",key,seconds,value))
//	return row, err
//}
//
//
