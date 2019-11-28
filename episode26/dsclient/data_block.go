package dsclient

type DataBlock struct {
	Data    []DataPoint `json:"data"`
	Summary string      `json:"summary"`
	Icon    string      `json:"icon"`
	// There are way more fields in a data block.
	// see https://darksky.net/dev/docs#data-block for all of them
}
