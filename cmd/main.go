package main

import (
	"flag"
	"ls-daily-signal/pkg/model"
	"time"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../configs/config.yaml", "-conf ./config.yaml")
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var conf model.Conf
	if err := c.Scan(&conf); err != nil {
		panic(err)
	}
	log.Infof("config: %+v", conf)
	srv, close, err := initService(&conf)
	if err != nil {
		panic(err)
	}
	
	var (
		// y, m, d = time.Now().Date()
		// date    = time.Date(y, m, d, 0, 0, 0, 0, time.Local)
		date    = time.Date(2022, 8, 25, 0, 0, 0, 0, time.Local)
	)

	if err := srv.CreateDailySignal(&date); err != nil {
		log.Errorf("err: %+v", err)
		panic(err)
	}
	close()
}
