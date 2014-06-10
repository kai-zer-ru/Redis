# About

Package to easy use [redis](http://redis.io) with [Go](http://golang.org/).

Worked with [redigo](https://github.com/garyburd/redigo/)

## Author

Kaizer666 - [http://vk.com/id_00000000000000000000000000](http://vk.com/id_00000000000000000000000000)

## Installation

    go get https://github.com/kovalyovmakc1990/Redis
    
## Usage

<pre>

package main

import (
      "https://github.com/kovalyovmakc1990/Redis"
      "fmt"
      )

func main() {

    MyRedis := Redis.RedisType{
        Host:"localhost",
        Port:1234,
        Password:"qweqweqweqw",
        DB:0,
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



