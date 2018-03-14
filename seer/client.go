package seer

import (
	"time"

	"google.golang.org/grpc"
)

// Client is a seer client
type Client struct {
	SeerClient
}

// New dials the specified seer server and returns a client.
func New(address string) (c *Client, err error) {
	conn, err := grpc.Dial(address, grpc.WithDefaultCallOptions())
	if err != nil {
		return nil, err
	}
	c = &Client{NewSeerClient(conn)}
	return c, nil
}

// CreateStream creates a stream with the specified data on the server.
func CreateStream(name string, period float64) (s *Stream, err error) {
	return
}

// GetStream retrieves the stream with the specified name.
func GetStream(name string) (s *Stream, err error) {
	return
}

// DeleteStream deletes the stream with the specified name.
func DeleteStream(name string) (err error) {
	return
}

// ListStreams returns a paged slice of streams.
func ListStreams(pageNum, pageSize int) (s []*Stream, err error) {
	return
}

// UpdateStream sends the provided data to the specific stream.
func UpdateStream(name string, times []time.Time, values []float64) (s *Stream, err error) {
	return
}

// GetForecast generates a forecast from the specified stream.
func GetForecast(name string, length int) (f *Forecast, err error) {
	return
}
