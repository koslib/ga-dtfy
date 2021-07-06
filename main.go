package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	ApiKey           string
	ScanProfileToken string
)

const (
	apiKeyInputVar           = "INPUT_API_KEY"
	scanProfileTokenInputVar = "INPUT_SCAN_PROFILE_TOKEN"
)

func main() {
	readEnv()

	provider := NewProvider(10 * time.Second)
	output, err := provider.StartScan()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(fmt.Sprintf(`::set-output name=result::%s`, output))
}

func readEnv() {
	ApiKey = os.Getenv(apiKeyInputVar)
	ScanProfileToken = os.Getenv(scanProfileTokenInputVar)
}
