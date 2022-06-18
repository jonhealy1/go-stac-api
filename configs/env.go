package configs

import (
	"fmt"
	"os"
	"strings"
)

func EnvMongoURI() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Hostname: ", hostname)
	if strings.Contains(hostname, "local") {
		return "mongodb://stac:root@localhost:27017"
	}

	return os.Getenv("MONGOURI")
}
