package client

import (
	"context"
	"time"

	"github.com/cshenton/seer-golang/seer"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

// Client is a seer client
type Client struct {
	conn seer.SeerClient
}

// New dials the specified seer server and returns a client.
func New(address string) (c *Client, err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c = &Client{seer.NewSeerClient(conn)}
	return c, nil
}

// CreateStream creates a stream with the specified data on the server.
func (c *Client) CreateStream(name string, period float64) (s *seer.Stream, err error) {
	in := &seer.CreateStreamRequest{
		Stream: &seer.Stream{
			Name:   name,
			Period: period,
		},
	}
	s, err = c.conn.CreateStream(context.Background(), in)
	return s, err
}

// GetStream retrieves the stream with the specified name.
func (c *Client) GetStream(name string) (s *seer.Stream, err error) {
	in := &seer.GetStreamRequest{
		Name: name,
	}
	s, err = c.conn.GetStream(context.Background(), in)
	return s, err
}

// DeleteStream deletes the stream with the specified name.
func (c *Client) DeleteStream(name string) (err error) {
	in := &seer.DeleteStreamRequest{
		Name: name,
	}
	_, err = c.conn.DeleteStream(context.Background(), in)
	return err
}

// ListStreams returns a paged slice of streams.
func (c *Client) ListStreams(pageNum, pageSize int) (s []*seer.Stream, err error) {
	in := &seer.ListStreamsRequest{
		PageNumber: int32(pageNum),
		PageSize:   int32(pageSize),
	}
	resp, err := c.conn.ListStreams(context.Background(), in)
	s = resp.Streams
	return s, err
}

// UpdateStream sends the provided data to the specific stream.
func (c *Client) UpdateStream(name string, times []time.Time, values []float64) (s *seer.Stream, err error) {
	ptimes := make([]*timestamp.Timestamp, len(times))
	for i := range times {
		ptimes[i], _ = ptypes.TimestampProto(times[i])
	}
	in := &seer.UpdateStreamRequest{
		Name: name,
		Event: &seer.Event{
			Times:  ptimes,
			Values: values,
		},
	}
	s, err = c.conn.UpdateStream(context.Background(), in)
	return s, err
}

// GetForecast generates a forecast from the specified stream.
func (c *Client) GetForecast(name string, length int) (f *seer.Forecast, err error) {
	in := &seer.GetForecastRequest{
		Name: name,
		N:    int32(length),
	}
	f, err = c.conn.GetForecast(context.Background(), in)
	return f, err
}
