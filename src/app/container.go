package app

type ServiceContainer struct {
  config Config
}

func GetApp(config Config) *ServiceContainer {
  return &ServiceContainer{config: config}
}
