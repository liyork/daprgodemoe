package main

import (
	"context"
	"errors"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/liyork/daprgodemoe/internal"
	"log"
	"net/http"
)

/*
dapr run --app-id go-service-c \
  --app-protocol http \
  --app-port 9003 \
  --dapr-http-port 3501 \
  --log-level debug \
  --components-path ./config \
  go run ./cmd/service/http/c/main.go
*/
func main() {
	s := daprd.NewService(":9003")
	if err := s.AddServiceInvocationHandler("/hello", helloHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

func helloHandler(_ context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}
	log.Printf("The go-service-c service method hello has been invoked,recieve message is %v", string(in.Data))
	httpResult := internal.HttpResult{Message: "This message is from Service C."}
	out = &common.Content{
		Data:        httpResult.ToBytes(),
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}
