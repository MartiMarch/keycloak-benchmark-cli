package managers

type Managers struct {
  env   Env
  conf  Conf
}

var env = newEnv()
var conf = newConf()

func Get() Managers {
  var object = Managers{
    env:  env,
    conf: conf,
  }

  return object
}
