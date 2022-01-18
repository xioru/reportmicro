package config

type Config struct {
	GetStationsURL string `env:"GET_STATION_URL"`
	SendReportURL  string `env:"SEND_REPORT_URL"`
}
