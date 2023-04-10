package managers

type Conf struct {
  domain        string
  keycloakType  string
  ping          func() string
}

func newConf() Conf {
  var env = newEnv()
  var object = Conf{
    domain:       env.domain,
    keycloakType: env.keycloakType,
    ping:         confPing,
  }

  return object
}

func confPing() string {
  return ""  
}
