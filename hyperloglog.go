package Redis

import (
	"github.com/redigo/redis"
)

func (r *RedisType) PfAdd (key string,element ...interface {}) (int, error) {
	params := make([]interface {},0)
	params = append(params,key)
	for _,v := range element {
		params = append(params,v)
	}
	row,err := redis.Int(r.RedisConn.Do("PFADD", params))
	return row, err
}

func (r *RedisType) PfCount (key ...interface {}) (int, error) {
	row,err := redis.Int(r.RedisConn.Do("PFCOUNT", key))
	return row, err
}
//PFMERGE destkey sourcekey
func (r *RedisType) Pfadd (destkey string,sourcekey ...string) (string, error) {
	params := make([]interface {},0)
	params = append(params,destkey)
	for _,v := range sourcekey {
		params = append(params,v)
	}
	row,err := redis.String(r.RedisConn.Do("PFMERGE", params))
	return row, err
}

