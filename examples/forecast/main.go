package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/stepan41k/git-test"
)

func main() {
	var timezone string
	var latitude, longitude float64

	flag.Float64Var(&latitude, "lat", 0.0, "latitude")
	flag.Float64Var(&longitude, "lon", 0.0, "longitude")
	flag.StringVar(&timezone, "timezone", "UTC", "timezone")

	flag.Parse()

	// create new openmeteo client instance
	client := git_test.New("")

	// call the forecast method to get current weather
	temperature, err := client.Forecast(context.Background(),
		git_test.ForecastParams{
			Latitude:  float32,
			Longitude: float32,
			Timezone:  timezone,
		})
	if err != nil {
		panic(err)
	}

	fmt.Println(temperature)
}