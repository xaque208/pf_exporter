package main

import (
	log "github.com/sirupsen/logrus"
	pfstats "github.com/xaque208/gopfstats"
)

func getStats() error {
	var fw pfstats.Pf
	var err error

	fw, err = pfstats.Open()
	if err != nil {
		return err
	}

	a, err := fw.Stats()
	log.Infof("%+v", a)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	err := getStats()
	if err != nil {
		log.Error(err)
	}
}
