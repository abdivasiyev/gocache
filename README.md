# Golang caching library

This library uses Redis or Memcached

![Build](https://github.com/MrWebUzb/gocache/actions/workflows/test.yml/badge.svg)

> Usage example

```golang
package main

import (
    "fmt"
    "github.com/MrWebUzb/gocache"
)

func main() {
    // Generate default configurations
    // Default driver is a redis cache
    // For enabling memcached
    // use following code
    // cfg := gocache.NewConfig().WithDriver(gocache.MemcachedDriver)
    cfg := gocache.NewConfig()

    client, err := gocache.NewClient(cfg)

    if err != nil {
        fmt.Printf("could not connect to redis client: %v\n", err)
        return
    }

    if !client.Ping() {
        fmt.Printf("could not connect to redis client\n")
        return
    }

    if err := client.Set("test", "Test data"); err != nil {
        fmt.Printf("could not set data: %v\n", err)
        return
    }

    data := client.Get("test")

    fmt.Println(data)
}
```

> Configurations

1. `WithDriver(gocache.Driver)` // available driver types are **RedisDriver**, **MemcachedDriver**
2. `WithAddr(addr1, addr2)` // replica addresses or single node address for both of redis and memcached
3. `WithCreds(username, password)` // redis credentials
4. `WithDB(0)` // database number in redis
5. `WithExpiration(10)` // expiration in seconds
