# Seer Golang Client
[![Build Status](https://travis-ci.org/cshenton/seer-golang.svg?branch=master)](https://travis-ci.org/cshenton/seer-golang)
[![Coverage Status](https://img.shields.io/coveralls/github/cshenton/seer-golang/master.svg)](https://coveralls.io/github/cshenton/seer-golang?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/cshenton/seer-golang)](https://goreportcard.com/report/github.com/cshenton/seer-golang)

The golang client for the seer forecasting server.


## Installation

To install the seer client library, use `go get`

```bash
go get github.com/cshenton/seer-golang/seer
```

## Usage

Since this is a client library, you'll need a server to communicate with. To get
up and running just:
```bash
docker run -d -p 8080:8080 cshenton/seer
```
Check out the project repo over [here](https://github.com/cshenton/seer)

Then, to interact over localhost

```go
package main

import (
        "log"
        "time"

        "github.com/cshenton/seer-golang/client"
)

func main() {
        // Create a client
        c, err := client.New("localhost:8080")
        if err != nil {
                log.Fatal(err)
        }

        // Create a stream
        _, err = c.CreateStream("myStream", 3600)
        if err != nil {
                log.Fatal(err)
        }

        // Add in data
        _, err = c.UpdateStream(
                "myStream",
                []float64{10, 9, 6},
                []time.Time{
                        time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
                        time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC),
                        time.Date(2016, 1, 3, 0, 0, 0, 0, time.UTC),
                },
        )
        if err != nil {
                log.Fatal(err)
        }

        // Generate and display forecast
        f, err = c.GetForecast("myStream", 10)
        fmt.Println(f)
}
```

## (For Contributors) Generating gRPC Client stubs

```bash
protoc -I ../seer/seer --go_out=plugins=grpc:seer ../seer/seer/seer.proto
```

We then add `//+build !test` to the top of the file to exclude it from coverage
statistics (since it has a tonne of redundant code, like the Get* methods and
the server interfaces, which are not used, and therefore don't require testing).