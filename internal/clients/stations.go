package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reportmicro/internal/models"
)

type Station struct {
	u          string
	httpclient *http.Client
}

func NewStation(u string) *Station {
	return &Station{
		u:          u,
		httpclient: &http.Client{},
	}
}

func (s *Station) GetStations() (models.StationList, error) {
	resp, err := s.httpclient.Get(s.u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("want status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
	var stations models.StationList
	if err := json.NewDecoder(resp.Body).Decode(&stations); err != nil {
		return nil, err
	}
	return stations, nil
}
