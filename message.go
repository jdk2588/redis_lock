package simulator

import ("sync")
type Message struct {
  sync.RWMutex
  client string
  locksfrom []*RedisNode
  notreachable int
  key string
  value string
  ttl int
}

func (cm *Message) setClient(id string) {
    cm.Lock()
    defer cm.Unlock()
    cm.client = id
}

func (cm *Message) setKey(msg string) {
    cm.key = msg
}

func (cm *Message) setVal(msg string) {
    cm.value = msg + cm.client
}

func (cm *Message) pushLockedNodes(r *RedisNode) {
    cm.locksfrom = append(cm.locksfrom, r)
}

func (cm *Message) getLockedNodes() []*RedisNode {
    return cm.locksfrom
}

func (cm *Message) incNAnodes() {
    cm.notreachable += 1
}

func (cm *Message) getNAnodes() int {
    return cm.notreachable
}

func (cm *Message) setTTL() {
    cm.ttl = 600
}

func (cm *Message) getClient() string {
      return cm.client
}

func (cm *Message) getKey() string {
      return cm.key
}

func (cm *Message) getVal() string {
      return cm.value
}

func (cm *Message) getTTL() int {
      return cm.ttl
}
