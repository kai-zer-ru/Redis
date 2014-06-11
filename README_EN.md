# ABOUT

Library to work with [redis] (http://redis.io) in language [Go] (http://golang.org/). 

It requires [redigo] (https://github.com/garyburd/redigo/) 

Refinement still in the process, write basic functions.

## Author

Kaizer666 - [http://vk.com/](http://vk.com/id_00000000000000000000000000)

## Install

    go get github.com/kaizer666/Redis
    
## Use

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



