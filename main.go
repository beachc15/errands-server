package main

import (
	"os"
	"os/signal"

	envconfig "github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

/*

	ENVIRONMENT VARIABLES:
	-----------------------------
	Set values via env varialbes, prefixed with ERRANDS_
	eg:

		ERRANDS_PORT=:4545 - Will change the listening port to 4545
		ERRANDS_STORAGE="/errands/" - Will change the DB location to /errands/


*/
var cfg Config

type Config struct {
	Storage string `split_words:"true" default:"./errands.db"`
	Port    string `split_words:"true" default:":5555"`
}

var server *ErrandsServer

func main() {

	// Parse Env Vars:
	err := envconfig.Process("ERRANDS", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	server = NewErrandsServer(&cfg)
	log.Info("listening for signals")
	for {
		select {
		case <-signals:
			// Logger.Info("main: done. exiting")
			log.Info("Exiting")
			server.kill()
			return
		}
	}

}
