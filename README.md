# О библиотеке

Библиотека дл работы с [redis](http://redis.io) на языке [Go](http://golang.org/).

Для работы требуется [redigo](!https://github.com/garyburd/redigo/)

Доработка ещё в процессе, написаы основные функции.
## Автор

Kaizer666 - [`http://vk.com/`](http://vk.com/id_00000000000000000000000000)

## Установка

    go get https://github.com/kaizer666/Redis
    
## Использование

<pre>

package main

import (
      "https://github.com/kaizer666/Redis"
      "fmt"
      )

func main() {
    MyRedis := Redis.RedisType{
        Host:"localhost",
        Port:1234,
        Password:"qweqweqweqw",// Необязательный параметр
        DB:0,// Необязательный параметр
        }
    MyRedis.Connect()
    defer MyRedis.Close()
    row,err := MyRedis.HGetAll("TestSetKey")
    if err != nil {
        panic(err)
    }
    fmt.Println(row)
}

</pre>



