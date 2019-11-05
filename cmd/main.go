package main

import (
	"fmt"
	"os"

	"istio.io/istio/mixer/adapter/auth_adapter_t"
)

func main() {
	addr := ""
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	s, err := auth_adapter_t.NewMyGrpcAdapter(addr)
	if err != nil {
		fmt.Printf("unable to start server: %v", err)
		os.Exit(-1)
	}

	shutdown := make(chan error, 1)
	go func() {
		s.Run(shutdown)
	}()
	_ = <-shutdown
}
