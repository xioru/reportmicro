package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reportmicro/internal/models"
)

type Aggregator struct {
	u          string
	httpclient *http.Client
}

func NewAggregator(u string) *Aggregator {
	return &Aggregator{
		u:          u,
		httpclient: &http.Client{},
	}
}

func (a *Aggregator) SendReport(report models.Report) error {
	body := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(body).Encode(report)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, a.u, body)
	if err != nil {
		return err
	}
	resp, err := a.httpclient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("want status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
	return nil
}
