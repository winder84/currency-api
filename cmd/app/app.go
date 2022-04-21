package main

import (
	"context"
	"currency-api/internal/app/config"
	"currency-api/internal/app/delivery/rest"
	"currency-api/internal/app/repository"
	"currency-api/internal/app/service"
	"currency-api/pkg/process"
	"flag"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/multierr"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Listener interface {
	Listen(addr string) error
}

const configPath = "configs/app.yaml"

var option = flag.String("o", "all", "Раздельный запуск(api || worker)")

func app() error {
	flag.Parse()
	osSignalCh := make(chan os.Signal)
	allOut := make(chan bool)
	signal.Notify(osSignalCh, syscall.SIGINT)
	go func() {
		for {
			select {
			case _, ok := <-osSignalCh:
				if ok {
					allOut <- true
					close(osSignalCh)
					return
				}
			}
		}
	}()

	logrus.Info("initialize app config")
	appConfig, err := config.New(configPath)
	if err != nil {
		return err
	}

	logrus.Info("initialize DB connection")
	db, err := sqlx.Open("postgres", appConfig.DB.Source)
	if err != nil {
		return errors.Wrap(err, "connect to DB")
	}
	defer func(err error) {
		multierr.AppendInto(&err, db.Close())
	}(err)

	logrus.Info("initialize repository")
	storage := repository.New(db)

	logrus.Info("initialize service usecases")
	useCase := service.New(storage, appConfig.Api)

	logrus.Info("initialize process usecases")
	proc := process.New(time.Duration(appConfig.Api.Process.Ticker) * time.Minute)

	logrus.Info("add usecase process")
	proc.Add(useCase.Currency.UpdateAll)

	logrus.Info("initialize service rest core")
	api := rest.New(useCase)

	logrus.Info("initialize application context")
	ctx, cancel := context.WithCancel(context.Background())

	if *option != "api" {
		go func() {
			proc.Wait(ctx)
		}()
	}

	if *option != "worker" {
		go func() {
			logrus.Infof("listen and serving http on port %s", appConfig.App.Port)
			lerr := listenHttp(api, ":"+appConfig.App.Port)
			if lerr != nil {
				multierr.AppendInto(&err, lerr)
				return
			}
		}()
	}

	<-allOut
	logrus.Info("shutdown service")
	cancel()
	close(allOut)
	time.Sleep(time.Second * 2)
	return nil
}

func listenHttp(l Listener, port string) error {
	return l.Listen(port)
}
