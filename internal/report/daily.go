package report

import (
	"log"
	"reportmicro/internal/clients"
)

type Daily struct {
	Aggregator *clients.Aggregator
	Station    *clients.Station
}

func (d *Daily) SendReport() {
	stations, err := d.Station.GetStations()
	if err != nil {
		log.Println(err)
		return
	}
	//todo:implement ...
	err = d.Aggregator.SendReport()
	if err != nil {
		log.Println(err)
		return
	}
}

func NewDaily(aggregator *clients.Aggregator, station *clients.Station) *Daily {
	return &Daily{
		Aggregator: aggregator,
		Station:    station,
	}
}
