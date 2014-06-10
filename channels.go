package Redis

import (
	"github.com/redigo/redis"
	"fmt"
)

func (r *RedisType) Publish (channel, message string) (int,error){
	row,err := redis.Int(r.RedisConn.Do("PUBLISH", channel, message))
	return row, err
}

func (r *RedisType) Pubsub() redis.PubSubConn {
	psc := redis.PubSubConn{r.RedisConn}
	return psc
}

func (r *RedisType) Listen (psc redis.PubSubConn) map[string]interface {}{
	data := make(map[string]interface {})
	switch v := psc.Receive().(type) {
	case redis.Message:
		data["channel"] = v.Channel
		data["data"] = fmt.Sprintf("%s",v.Data)
		data["type"] = "message"
		return data
	case redis.Subscription:
		data["channel"] = v.Channel
		data["data"] = v.Kind
		data["type"] = "subscribe"
		return data
	case error:
		panic(v)
	}
	return nil
}
