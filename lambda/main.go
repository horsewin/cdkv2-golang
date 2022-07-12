package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func handle(ctx context.Context, event MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", event.Name), nil
}

// Caveat! main func is required
func main() {
	lambda.Start(handle)
}
