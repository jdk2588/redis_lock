package simulator

import (
  "time"
)

type Instance interface {
    setRack(int)
    setRegion(int)
    getRack() int
    getRegion() int
    getIdent() string
}

func checkNetworkPartition(node1, node2 Instance) error {
    switch Env.PartitionType {
    case "rack":
        if node1.getRegion() == node2.getRegion() && node1.getRack() != node2.getRack() {
                time.Sleep(1*time.Second)
                return rackerror
        }
      case "region":
        if node1.getRegion() != node2.getRegion() {
                time.Sleep(1*time.Second)
                return regionerror
        }
      case "rack+region":
        if (node1.getRegion() != node2.getRegion() || node1.getRack() != node2.getRack()) {
                time.Sleep(1*time.Second)
                return rackplusregionerror
        }
    }
    return nil
}

func assignRackandRegion(node Instance, index int) {

  switch index {
  case 0,4,8:
          node.setRack(1)
          node.setRegion(1)
  case 1,5,9:
          node.setRack(1)
          node.setRegion(2)
  case 2,6,10:
          node.setRack(2)
          node.setRegion(1)
  case 3,7,11:
          node.setRack(2)
          node.setRegion(2)
  }
}
