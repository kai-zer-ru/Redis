package Redis

import "github.com/redigo/redis"

func (r *RedisType) SAdd (key string, members ...string) (int, error){
	params := make([]interface {},0)
	params = append(params,key)
	for _,v := range members{ params = append(params, v) }
	row,err := redis.Int(r.RedisConn.Do("SADD",params...))
	return row, err
}

func (r *RedisType) SCard (key string) (int,error){
	row,err := redis.Int(r.RedisConn.Do("SCARD",key))
	return row, err
}

func (r *RedisType) SDiff (key ...interface {}) (interface {},error) {
	row,err := r.RedisConn.Do("SDIFF",key)
	return row, err
}

func (r *RedisType) SDiffStore (destination string,key ...interface {}) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("SDIFFSTORE",destination,key))
	return row, err
}

func (r *RedisType) SInter (key ...interface {}) (interface {},error) {
	row,err := r.RedisConn.Do("SINTER",key)
	return row, err
}

func (r *RedisType) SInterScore (destination string,key ...interface {}) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("SINTERSTORE",destination,key))
	return row, err
}

func (r *RedisType) SIsMember (key, member string) (bool,error){
	row,err := redis.Bool(r.RedisConn.Do("SISMEMBER",key, member))
	return row, err
}

func (r *RedisType) SMembers (key string) (interface {},error){
	row,err := redis.Strings(r.RedisConn.Do("SMEMBERS",key))
	if err == redis.ErrNil{ return nil, nil }
	return row, err
}

func (r *RedisType) SMove (source,destination string, member interface {}) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("SMOVE", source, destination, member))
	return row, err
}
func (r *RedisType) SPop (key string) (string,error){
	row,err := redis.String(r.RedisConn.Do("SPOP",key))
	if err == redis.ErrNil{ return "", nil }
	return row, err
}

func (r *RedisType) SRandMember (key string, count ...int) (interface {},error){
	if len(count) > 0{
		row,err := redis.Strings(r.RedisConn.Do("SRANDMEMBER",key,count[0]))
		return row,err
	} else {
		row, err := redis.String(r.RedisConn.Do("SRANDMEMBER", key))
		return row, err
	}
}

func (r *RedisType) SRem (key string,members ...string) (int,error){
	params := make([]interface {},0)
	params = append(params,key)
	for _,v := range members{ params = append(params,v) }
	row,err := redis.Int(r.RedisConn.Do("SREM",params...))
	return row, err
}

func (r *RedisType) SUnion (key ...interface {}) (interface {},error){
	row,err := r.RedisConn.Do("SUNION",key...)
	return row, err
}

func (r *RedisType) SUnionStore (destination,key string) (int,error){
	row,err := redis.Int(r.RedisConn.Do("SUNIONSTORE",destination,key))
	return row, err
}

//SSCAN key cursor [MATCH pattern] [COUNT count]
//func (r *RedisType) SScan () () {
//
//}
