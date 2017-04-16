package simulator

import (
  "fmt"
  "github.com/fatih/color"
)

var red = color.New(color.FgRed)

var boldRed = red.Add(color.Bold)

var green = color.New(color.FgGreen)

var boldGreen = green.Add(color.Bold)

var blue = color.New(color.FgBlue)

var boldBlue = blue.Add(color.Bold)

var yellow = color.New(color.FgYellow)

var boldYellow = yellow.Add(color.Bold)

var magenta = color.New(color.FgMagenta)

var boldMagenta = magenta.Add(color.Bold)

var cyan = color.New(color.FgCyan)

var boldCyan = cyan.Add(color.Bold)

func detailLog(text, senderId string, node1, node2 Instance) {
     if Env.Debug {
       fmt.Println(text+" "+node1.getIdent()+" "+node2.getIdent()+" "+senderId+"\n")
     }
}

func clientLog(text, senderId string, node1 Instance) {
     if Env.Debug {
       fmt.Println(text+" "+node1.getIdent()+" "+senderId+"\n")
     }
}

func simpleLog(text string){
    if Env.Debug {
      fmt.Println(text+"\n")
    }
}

func failLog(text, senderId string, node1, node2 Instance) {
    if Env.Debug {
     boldRed.Println(text+" "+node1.getIdent()+" "+node2.getIdent()+" "+senderId+"\n")
    }
}

func lockGiven(text, senderId string, node1, node2 Instance) {
    if Env.Debug {
       boldYellow.Println(text+" "+node1.getIdent()+" "+node2.getIdent()+" "+senderId+"\n")
    }
}

func successLog(text, senderId string, node1 Instance) {
      boldGreen.Println(text+" "+node1.getIdent()+" "+senderId+"\n")
}

func inactiveLog(text, senderId string, node1 Instance) {
    if Env.Debug {
      boldBlue.Println(text+" "+node1.getIdent()+" "+senderId+"\n")
    }
}

func lockRelease(text, senderId string, node1, node2 Instance) {
    if Env.Debug {
     boldMagenta.Println(text+" "+node1.getIdent()+" "+node2.getIdent()+" "+senderId+"\n")
    }
}

func noLockgiven(text, senderId string, node1, node2 Instance) {
    if Env.Debug {
     boldCyan.Println(text+" "+node1.getIdent()+" "+node2.getIdent()+" "+senderId+"\n")
   }
}
