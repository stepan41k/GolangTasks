package gitTest

import (
	"net/http"
	"time"
	"context"
	"fmt"
	"io"
	"errors"
	"encoding/json"
)

type ForecastParams struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone string
}

type ForecastResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Current   struct {
		Time        string  `json:"time"`
		Temperature float64 `json:"temperature_2m"`
	} `json:"current"`
	Error  bool   `json:"error"`
	Reason string `json:"reason"`
}

type Client struct {
	host string
	client *http.Client
}

const defaultAPIHost = "https://api.open-meteo.com"

var (
	ErrOpenMeteoService       = errors.New("open-meteo service error")
	ErrOpenMeteoForecastError = errors.New("open-meteo forecast error")
)

func New(host string) *Client {
	apiHost := defaultAPIHost
	if host != "" {
		apiHost = host
	}

	return &Client{
		host: apiHost,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

const forecastURL = "/v1/forecast?latitude=%f&longitude=%f&current=temperature_2m&hourly=temperature_2m&timezone=%s"

func (cl *Client) Forecast(ctx context.Context, params ForecastParams) (float64, error) {
	addr := fmt.Sprintf(cl.host+forecastURL, params.Latitude, params.Longitude, params.Timezone)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, addr, http.NoBody)

	resp, err := cl.client.Do(req)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, ErrOpenMeteoService
	}

	var response ForecastResponse

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return 0, err
	}

	if response.Error {
		return 0, fmt.Errorf("%w: %s:", ErrOpenMeteoForecastError, response.Reason)
	}

	return response.Current.Temperature, nil
}
