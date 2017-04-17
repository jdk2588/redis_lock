To run the simulator:

    cd exec
    go run main.go -c config.xml

    Change the following values in config.xml:
      servers: pass the host with port
      mock: if not able to use actual redis servers, set to true.
      numberofclients: number of clients to be used in simulation.
      partitiontype: choose any one from [rack, region, rack+region]
      messages: number of messages to be received by clients
      failanyredis: want to simulate with an unreachable redis node to all clients ?
      failredisnode: pass the index number from 0 to n-1 (n is total servers), currently only one node failure supported.
      debug: set to true, to get debugging message, which will give idea about how the simulation is behaving,
             if set to false, will print only if lock acquired by a client, else nothing
