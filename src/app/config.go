package app

import "github.com/spf13/viper"

type Config struct {
  HttpPort string
}

func GetConfig() Config {
  return Config{
    HttpPort: viper.GetString("HTTP_PORT"),
  }
}
