package commands

func Exec(command string, flags []string) string {
  var output string
  var help = newHelp()

  if command == "help" || command == "" {
    output = help.exec()
  }

  return output
}
