package main

import (
	"github.com/simplegrpc/grpc-go-course/primeNumberDecompositon/primepb"
)

type server struct{}

func (c *server) PrimeManyTimes(req *primepb.PrimeManytimesRequest) (*primepb.PrimeManytimesResponse, error) {

}

func main() {

}
