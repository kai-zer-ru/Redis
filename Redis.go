package Redis

import (
	"github.com/redigo/redis"
	"fmt"
	"strings"
)

type RedisType struct {
	RedisConn 	redis.Conn
	Host 		string
	Port 		uint32
	Password 	string
	ErrRed		error
	DB			int
}
// Подключение
func (r *RedisType) Connect() error{
	r.RedisConn,r.ErrRed = redis.Dial("tcp",fmt.Sprintf("%s:%v",r.Host,r.Port))
	if r.ErrRed != nil {
		return r.ErrRed
	}
	_,errAuth := r.RedisConn.Do("AUTH", r.Password)
	if errAuth != nil {
		return errAuth
	}
	if r.DB != 0 {
		r.RedisConn.Send("SELECT",r.DB)
	}
	return nil
}

func (r RedisType) Close() error {
	err := r.RedisConn.Close()
	return err
}


func (r *RedisType) Keys (pattern string) ([]string,error) {
	row,err := redis.Strings(r.RedisConn.Do("KEYS",pattern))
	return row, err
}

func (r *RedisType) Exists (key string) (bool,error) {
	exist,err := redis.Bool(r.RedisConn.Do("EXISTS",key))
	return exist, err
}

func (r *RedisType) Delete (keys ...interface {}) error {
	err := r.RedisConn.Send("DEL",keys...)
	return err
}

func (r *RedisType) Expire (key string, seconds uint32) (int, error) {
	row,err := redis.Int(r.RedisConn.Do("EXPIRE",key,seconds))
	return row, err
}

func (r *RedisType) Ttl (key string) (int,error) {
	row, err := redis.Int(r.RedisConn.Do("TTL",key))
	if row < 0{
		row = 0
	}
	return row, err
}

func (r *RedisType) ExpireAt (key string, timestamp uint32) (int,error) {
	row,err := redis.Int(r.RedisConn.Do("EXPIREAT",key,timestamp))
	return row, err
}

func (r *RedisType) Info () (map[string]interface {},error) {
	row, err := r.RedisConn.Do("INFO")
	if err != nil {
		return nil, err
	}
	row2,err2 := redis.String(row,err)
	info_strings := strings.Split(row2,"\r\n")
	data := make(map[string]interface {})
	title := ""
	pak := make(map[string]interface {})
	for _,values := range info_strings {
		if strings.Count(values,"#") != 0 {
			if title != "" {
				data[title] = pak
			}
			title = values
			pak = make(map[string]interface {})
			continue
		}
		d := strings.Split(values,":")
		if len(d) == 1 {
			// Если пустое значение
			continue
		}
		c := strings.Split(d[1],",")
		if len(c) > 1{
			data2 := make(map[string]string)
			for _,q := range c {
				w := strings.Split(q, "=")
				data2[w[0]] = w[1]
			}
			pak[d[0]] = data2

		} else {
			pak[d[0]] = d[1]
		}
	}
	data[title] = pak
	return data, err2
}

func (r *RedisType) GetRedisReply (answer interface {},err error,names []string) map[string]interface {}{
	var reply []string
	res,_ := redis.Values(answer,err)
	for _, x := range res {
		var v, ok = x.([]byte)
		if ok { reply = append(reply,string(v)) }
		if x == nil { reply = append(reply,"false") }
	}
	var resp = make(map[string]interface {})
	if len(names) > 0 { for i := 0 ; i < len(reply); i++ { resp[names[i]] = reply[i] } } else { for i := 0 ; i < len(reply); i+=2 { resp[reply[i]] = reply[i+1] } }
	return resp
}

func (r *RedisType) GetRedisReplyArray (answer interface {},err error) []string {
	var reply []string
	res,err2 := redis.Values(answer,err)
	if err2 != nil {
		fmt.Println(err2)
		return reply
	}
	for _, x := range res {
		var v, ok = x.([]byte)
		if ok {
			reply = append(reply,string(v))
		}
		if x == nil {
			reply = append(reply,"false")
		}
	}
	return reply
}

func (r *RedisType) GetBool(row interface {},err error) (bool,error) {
	if err != nil {
		return false, err
	}
	switch row.(type){
	case string:
		if row.(string) == "OK" {return true, nil} else {return false, err}
	case int,uint32,uint8,uint16,uint64:
		if row == 1 {return true, nil} else {return false, err}
	case float32,float64:
		if row == 1.0 {return true, nil} else {return false, err}
	}
	return false, nil
}
