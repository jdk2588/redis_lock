package main

import (
  "os"
  "fmt"
  "flag"
  "sync"
  "time"
  "strconv"
   S  "simulator"
   Conf "simulator/config"
  "github.com/pborman/uuid"
)

var usageStr = `
Usage: go run main.go [options]

Options:
    -c  --config                     Set parameters in config file
Common Options:
    -h, --help                       Show this message
    -v, --version                    Show version
`

func usage() {
        fmt.Printf("%s\n", usageStr)
        os.Exit(0)
}

func main() {
  var wg sync.WaitGroup

  var configFile string

  flag.StringVar(&configFile, "c", "", "path to configuration YML file.")
  flag.StringVar(&configFile, "config", "", "path to configuration YML file.")
  flag.Usage = usage
  flag.Parse()

  S.Env = Conf.DefaultConfig()

  if (configFile != "") {
    S.Env, _ = Conf.LoadConfig(configFile)
  }

  if len(os.Args) < 2 {
    usage()
  }

  S.Init()

  var redisn = S.GetRedisNodes(S.Env.Servers)
  var clients = S.GetClients()

  fmt.Println("Running simulation with " + strconv.Itoa(len(redisn)) + " servers and " + strconv.Itoa(len(clients)) + " clients for " + strconv.Itoa(S.Env.MessageCount) + " messages.\n")

  if S.Env.Mock {
    fmt.Println("In mock mode, set mock to false in config file to run with actual redis nodes. \n")
  }

  time.Sleep(1*time.Second)

  ch := make(chan string)
  for i:=0;i<S.Env.MessageCount;i++ {

    wg.Add(1)

    go func() {
      msg := uuid.New()
      ch <- msg
    }()

    go S.Dispatcher(ch, &wg, redisn, clients)
  }
  wg.Wait()
}
