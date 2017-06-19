package simulator

import (
  "sync"
)

func Dispatcher(ch chan string, wg *sync.WaitGroup, redisn []*RedisNode, clients []*Client) {
  defer wg.Done()

  message := <-ch

  simulate(message, redisn, clients)

}

func simulate(message string, redisn []*RedisNode, clients []*Client) {

  var wg1 sync.WaitGroup
  for _ , c := range(clients) {
    wg1.Add(1)

    c.redisNodes = redisn
    c.setQuorum()

    go c.process(message, &wg1)

   }

   wg1.Wait()
}
