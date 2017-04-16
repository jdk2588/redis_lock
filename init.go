package simulator

import (
  "sync"
  "simulator/config"
)

var clientC = &ClientCaches{}
var rCache = &RedisCaches{}
var Env config.Settings

type ClientCaches struct {
  sync.RWMutex
  clientMap  map[int]*Client
  clientInit []sync.Once
}

type RedisCaches struct {
    sync.RWMutex
    caches map[string]*RedisNode
    onceArr []sync.Once
}

func Init() {
  clientC.clientMap = map[int]*Client{}
  clientC.clientInit = make([]sync.Once, Env.NoOfClients)
  rCache.caches = map[string]*RedisNode{}
  rCache.onceArr = make([]sync.Once, len(Env.Servers))
}
