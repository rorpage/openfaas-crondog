package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/robfig/cron"
	"github.com/rorpage/crondog/types"
)

var functionURL string
var functionData string

func main() {
	osEnv := types.OsEnv{}
	readConfig := ReadConfig{}
	config := readConfig.Read(osEnv)

	functionURL = config.functionURL
	functionData = config.functionData

	c := cron.New()
	c.AddFunc(config.cronSchedule, InvokedFunction)
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

// InvokedFunction is called on CRON schedule
func InvokedFunction() {
	resp, err := http.Post(functionURL,
		"application/json",
		strings.NewReader(functionData))

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", body)
}
