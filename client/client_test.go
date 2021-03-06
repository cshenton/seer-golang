package client_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/cshenton/seer-golang/client"
	"github.com/golang/protobuf/ptypes"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func setUp(t *testing.T) (name string, c *client.Client) {
	c, err := client.New("localhost:8080")
	if err != nil {
		t.Fatal("unexpected error in New:", err)
	}
	name = randString(10)
	_, err = c.CreateStream(name, 86400)
	if err != nil {
		t.Fatal("unexpected error in CreateStream:", err)
	}
	return name, c
}

func TestNew(t *testing.T) {
	_, err := client.New("localhost:8080")
	if err != nil {
		t.Error("unexpected error in New:", err)
	}
}

func TestCreateStream(t *testing.T) {
	_, c := setUp(t)

	name := randString(10)

	s, err := c.CreateStream(name, 3600)
	if err != nil {
		t.Fatal("unexpected error in CreateStream:", err)
	}
	if s.Name != name {
		t.Errorf("expected name %v, but got %v", name, s.Name)
	}
	if s.Period != 3600 {
		t.Errorf("expected period %v, but got %v", 3600, s.Period)
	}
}

func TestGetStream(t *testing.T) {
	name, c := setUp(t)

	s, err := c.GetStream(name)
	if err != nil {
		t.Fatal("unexpected error in GetStream:", err)
	}
	if s.Name != name {
		t.Errorf("expected name %v, but got %v", name, s.Name)
	}
	if s.Period != 86400 {
		t.Errorf("expected period %v, but got %v", 86400, s.Period)
	}
}

func TestDeleteStream(t *testing.T) {
	name, c := setUp(t)

	err := c.DeleteStream(name)
	if err != nil {
		t.Error("unexpected error in DeleteStream:", err)
	}
}

func TestListStreams(t *testing.T) {
	name, c := setUp(t)

	s, err := c.ListStreams(1, 10)
	if err != nil {
		t.Error("unexpected error in ListStreams:", err)
	}
	if len(s) > 10 {
		t.Errorf("stream list length %v greater than page size %v", len(s), 10)
	}

	contains := false
	for i := range s {
		if s[i].Name == name {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected list to contain stream %v, but it didn't", name)
	}
}

func TestUpdateStream(t *testing.T) {
	name, c := setUp(t)

	s, err := c.UpdateStream(
		name,
		[]time.Time{
			time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2016, 1, 3, 0, 0, 0, 0, time.UTC),
		},
		[]float64{10, 9, 6},
	)
	if err != nil {
		t.Error("unexpected error in UpdateStream:", err)
	}
	if s.Name != name {
		t.Errorf("expected name %v, but got %v", name, s.Name)
	}

	ts, _ := ptypes.Timestamp(s.LastEventTime)
	if !ts.Equal(time.Date(2016, 1, 3, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("expected last event %v, but got %v", time.Date(2016, 1, 3, 0, 0, 0, 0, time.UTC), ts)
	}
}

func TestGetForecast(t *testing.T) {
	name, c := setUp(t)

	_, err := c.UpdateStream(name, []time.Time{time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)}, []float64{20})
	if err != nil {
		t.Error("unexpected error in UpdateStream:", err)
	}

	f, err := c.GetForecast(name, 100)
	if err != nil {
		t.Error("unexpected error in GetForecast:", err)
	}

	if len(f.Values) != 100 {
		t.Errorf("expected values length %v, but got %v", 100, len(f.Values))
	}
}
