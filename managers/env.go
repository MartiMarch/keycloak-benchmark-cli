package managers

import(
  "fmt"
  "os"
  "log"
  "github.com/joho/godotenv"
)

type Env struct {
  domain        string
  keycloakType  string
}

func newEnv() Env {
  var domain string
  var keycloakType string
  var path string  

  path = shellExec("pwd") + "/managers/keycloak.env"
  var err = godotenv.Load(path)
  if err != nil {
    log.Fatalf("Error loading ../managers/keycloak.env file:\n[%v]", err)
  }
  
  domain = envGet("DOMAIN")
  keycloakType = envGet("TYPE")
  var object = Env {
    domain:        domain,
    keycloakType:  keycloakType,
  }

  return object
}

func envGet(parameter string) string {
  var errorMsg string
  var value string

  value = os.Getenv(parameter)
  errorMsg = fmt.Sprintf(`
    Error loading ../managers/keycloak.env, parameter %s not found.
    
    Expected content:

    ** keycloak.env
    * DOMAIN=http://IP:POR  > IP:PORT or DOMAIN of Keycloak server
    * TYPE=wildfly          > Keycloak installation mode, valid values: wildfly, minikube, quarkus
    ***************
  `, parameter)

  if value == "" {
    log.Fatalf(errorMsg)
  }

  return value
}
