// Package main is the my-service service.
package main

import (
  "fmt"
  "github.com/namsral/flag"
)

// ServiceName is the public name of this service.
const ServiceName = "my-service"

var config = Service1Config{}

func parseFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "Host on which to run")
	flag.IntVar(&config.Port, "port", 9000, "Port on which to run")
	flag.Parse()
}

func main() {
	parseFlags()

		s :=Service1{
			Service1Config: config,
		}
		if err := s.Run(); err != nil {
			fmt.Printf("Run error: %s", err)
		}
}
