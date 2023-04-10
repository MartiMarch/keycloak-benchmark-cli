package commands

type Help struct {
  exec func() string
}

func newHelp() Help {
  var object = Help {
    exec: helpExec,
  }

  return object
}

func helpExec() string {
  return `
***********************************************************************************
*   __   ___   __  _    ___  _  _   __ __    __    ___   _  __     ___  _     _   *
*  |  \ | __| |  \| |  / _/ | || | |  V  |  /  \  | _ \ | |/ /    / _/ | |   | |  *
*  | -< | _|  | | ' | | \__ | >< | | \_/ | | /\ | | v / |   <    | \__ | |_  | |  *
*  |__/ |___| |_|\__|  \__/ |_||_| |_| |_| |_||_| |_|_\ |_|\_\    \__/ |___| |_|  *
*                                                                                 *
***********************************************************************************

 This CLI is used to aotumatize Keycloak benchmark module:
 > https://www.keycloak.org/keycloak-benchmark

 Usage:
   
  1. Create executable using "go build"

  2. Create the file .../managers/keycloak.env with the next content:
    
    ** keycloak.env
    * DOMAIN=http://IP:POR  > IP:PORT or DOMAIN of Keycloak server
    * TYPE=wildfly          > Keycloak installation mode, valid values: wildfly, minikube, quarkus
    ***************

  3. Execute CLI 

     .../keycloak-cli <command> <flag 1> <flag 2> ... <flag n>

     To assign a value to a specific flag use "=",  <flag>=<value>

 Commands:
  help        Provides information abouth CLI

  * To obtain extra information of any command use "<command> help"
`
}
