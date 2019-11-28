package dsclient

type DataPoint struct {
	ApparentTemp float64 `json:"apparentTemperature"`
	// There are way more fields in a data point.
	// see https://darksky.net/dev/docs#data-point for all of them
}
