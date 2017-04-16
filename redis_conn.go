package simulator

import (
    "sync"
    "time"
    "github.com/garyburd/redigo/redis"
)

type RedisPool struct {
  sync.Mutex
  pool *redis.Pool
}

var delScript = redis.NewScript(1, `
if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("del", KEYS[1])
else
	return 0
end`)

func createPool(server string) *redis.Pool {
  return &redis.Pool{
  		MaxIdle:     3,
  		IdleTimeout: 240 * time.Second,
  		Dial: func() (redis.Conn, error) {
  			c, err := redis.Dial("tcp", server)
  			if err != nil {
  				return nil, err
  			}

  			return c, err
  		},
      TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		  },
  }
}

func oneServer(server string, index int) *RedisNode {

    rCache.Lock()
    defer rCache.Unlock()

    for i := 0;i < len(Env.Servers); i++ {
      if i == index {
      	rCache.onceArr[i].Do(func() {
          r := &RedisNode{}
          if Env.Mock {
            r.lockProvided = map[string]string{}
          } else {
            rp := &RedisPool{pool: createPool(server)}
            r.pool = rp
          }

          r.setIdent(server)
          assignRackandRegion(r, i)
      		rCache.caches[server] = r
      	})
      }
    }

	  return rCache.caches[server]
}
