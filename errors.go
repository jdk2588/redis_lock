package simulator

import (
  "errors"
)

var lockfail = errors.New("Unable to acquire lock!")
var unlockfail = errors.New("Unable to release lock!")
var alreadyprocessed = errors.New("Already processed")
var regionerror = errors.New("Region partition  error!")
var rackerror = errors.New("Rack partition error!")
var rackplusregionerror = errors.New("Rack & Region partition error!")
var nodeerror = errors.New("Redis node not reachable!")

func getPartError(e error) error {
      switch e {
        case regionerror:
        case rackerror:
        case rackplusregionerror:
        case nodeerror:
      }
      return e
}
