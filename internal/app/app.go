package app

import (
	"github.com/robfig/cron/v3"
	"reportmicro/internal/clients"
	"reportmicro/internal/config"
	"reportmicro/internal/report"
)

type Application struct {
	stationsClient   *clients.Station
	aggregatorClient *clients.Aggregator
}

func New(cfg config.Config) *Application {
	return &Application{
		stationsClient:   clients.NewStation(cfg.GetStationsURL),
		aggregatorClient: clients.NewAggregator(cfg.SendReportURL),
	}
}

func (a *Application) Run() error {
	reportDaily := report.NewDaily(a.aggregatorClient, a.stationsClient)
	c := cron.New()
	//todo:use cron pattern
	_, err := c.AddFunc("", reportDaily.SendReport)
	if err != nil {
		return err
	}
	_, err = c.AddFunc("", func() {

	})
	if err != nil {
		return err
	}
	c.Start()
	return nil
}
