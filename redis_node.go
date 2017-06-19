package simulator

import (
  "sync"
)

type RedisNode struct {
     mu sync.RWMutex
     pool *RedisPool
     ident string
     lockProvided  map[string]string
     killed bool
     rack int
     region int
}

func (r *RedisNode) getRegion() int {
      return r.region
}

func (r *RedisNode) setRegion(region int) {
      r.region = region
}

func (r *RedisNode) getRack() int {
      return r.rack
}

func (r *RedisNode) setRack(rack int) {
      r.rack = rack
}

func (r *RedisNode) getIdent() string {
      return r.ident
}

func (r *RedisNode) setIdent(addr string) {
     r.ident = addr
}

func(r *RedisNode) activate() {
    r.killed = false
}

func (r *RedisNode) deactivate() {
    r.killed = true
}

func (r *RedisNode) isActive() bool {
    return !r.killed
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func (r *RedisNode) lock(m *Message, c *Client) error {

  if Env.FailAnyRedis && contains(Env.FailRedisNode, r.getIdent()) {
      r.deactivate()
      return nodeerror
  }

  neterr := checkNetworkPartition(r, c)

  if neterr != nil {
      return neterr
  }

  r.mu.Lock()
  defer r.mu.Unlock()

  if Env.Mock {

    val, ok := r.lockProvided[m.key]
    if val == c.getIdent() {
        // detailLog("Already have lock", m.key, c, r)
        return alreadyprocessed
    }

    if ok {
        return lockfail
    } else {
      r.lockProvided[m.key] = c.getIdent()
    }

  } else {
    conn := r.pool.pool.Get()
    v, e := conn.Do("SET", m.key, m.value, "NX", "EX", m.ttl)
    conn.Close()

    if e != nil {
        return e
    }

    if v == nil {
        return lockfail
    }
  }

  lockGiven("Giving lock to", m.key, c, r)
  return nil

}

func (r *RedisNode) unLock(m *Message) error {
  r.mu.Lock()
  defer r.mu.Unlock()

  if Env.Mock {
    delete(r.lockProvided, m.key)
  } else {
    conn := r.pool.pool.Get()
    v, e := delScript.Do(conn, m.key, m.value)
    conn.Close()

    if e != nil {
        return e
    }

    if v == nil {
        return unlockfail
    }
  }

  return nil
}

func GetRedisNodes(servers []string) []*RedisNode {
  rNodes := []*RedisNode{}
  for i, addr := range servers {
      r := oneServer(addr, i)
      if !r.isActive() {
        continue
      }
      rNodes = append(rNodes, r)
  }

  return rNodes
}
