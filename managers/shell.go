package managers

import(
  "log"
  "os/exec"
  "strings"
)

func shellExec(command string) string {
  shell := exec.Command(command)
  output, err := shell.Output()
  if err != nil {
    log.Fatalf("Error executing command \"",command,"\":\n",err)
  }
  return strings.ReplaceAll(string(output), "\n", "")
}
