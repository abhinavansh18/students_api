package config

import (
	"flag"
	"log"
	"os"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env          string "yaml: 'env'  env:'ENV' env-required:'true' env-default:'production'"
	Storage_Path string "yaml: 'storage_path' env-required:'true'"
	HTTPServer   "yaml: 'httpserver'"
}

func MustLoad() {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to the config file")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatal("Config PATH IS NOT SET")
		}

	}

	if _, err := os.Stat("configPath"); os.IsNotExist(err) {
		log.Fatalf(" config file does not exist: %s", configPath)
	}
}
