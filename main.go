package main

import (
  "fmt"
  "os"
  "keycloak-benchmark-cli/commands"
  "keycloak-benchmark-cli/managers"
)


func main(){
  var command string = os.Args[1]
  var flags []string = os.Args[2:]
  var output string

  //Configuration loaded
  //var mngs = managers.Get()
  //fmt.Println(mngs)

  //Execution of a command
  output = commands.Exec(command, flags)
  fmt.Println(output)
}
