package simulator

import (
 "sync"
 "strconv"
)

func (c *Client) process(msg string, wg *sync.WaitGroup) {

  defer wg.Done()

  if (c.isActive() == false) {
    simpleLog("Node deactive, not participating")
    return
  }

  m := &Message{client: c.getIdent()}
  m.setKey(msg)
  m.setVal(msg)
  m.setTTL()

  for _, r := range c.redisNodes {
      var nolock = false
      if err := r.lock(m, c); err != nil {

        switch err {
          case lockfail:
                nolock = true
                noLockgiven(err.Error(), msg, c, r)
                break
          case getPartError(err):
                failLog(err.Error(), msg, c, r)
                m.incNAnodes()
        }

        if (nolock && len(m.getLockedNodes()) < c.quorum)  {
          c.releaseLocks(m)
          break
        }

      } else {
        m.pushLockedNodes(r)
      }
  }

  if (m.getNAnodes() >= c.quorum) {
      inactiveLog("Marking node deactive", msg, c)
      c.deactivate()
      c.releaseLocks(m)

      inactiveLog("Re-start suggested by", msg, c)

      redisn := GetRedisNodes(Env.Servers)
      clients := GetClients()
      simulate(msg, redisn, clients)
  }

  if (len(m.getLockedNodes()) >= c.quorum) {
      successLog("Lock acquired with quorum " + strconv.Itoa(c.quorum), msg, c)
  }
}

func (c *Client) releaseLocks(cm *Message) {
    locksAc := cm.getLockedNodes()
    for i := len(locksAc) - 1 ; i >= 0; i-- {
          rn := locksAc[i]
          lockRelease("Releasing lock for", cm.key, c, rn)
          err := rn.unLock(cm)

          //If not able to release lock, how to handle
          if (err != nil) {
              simpleLog(err.Error())
          }
    }
}
