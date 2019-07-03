package cmd

import (
	"fmt"

	"github.com/arschles/go-in-5-minutes/episode26/dsclient"
	"github.com/spf13/cobra"
)

func temp(apiKey string) (*cobra.Command, error) {
	cl, err := dsclient.New(apiKey)
	if err != nil {
		return nil, err
	}
	cmd := &cobra.Command{
		Use: "temp",
	}
	flags := cmd.PersistentFlags()
	lat := flags.Float64("lat", 0, "The Latitude to fetch")
	long := flags.Float64("long", 0, "The Longitude to fetch")
	cmd.RunE = func(*cobra.Command, []string) error {
		fmt.Printf("Getting temp for (%f, %f)\n", *lat, *long)
		fcast, err := cl.Forecast(*lat, *long)
		if err != nil {
			return err
		}
		if len(fcast.Hourly.Data) < 1 {
			return fmt.Errorf("No hourly data returned!")
		}
		hourly := fcast.Hourly.Data[0]
		fmt.Printf(
			"Temperature for your location (%f, %f): %f\n",
			*lat,
			*long,
			hourly.ApparentTemp, // the API only allows apparent temp on hourly
		)
		return nil
	}
	return cmd, nil
}
