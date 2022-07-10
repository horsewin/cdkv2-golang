package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

// Lambda 関数コードが実行されるエントリポイント。これは必須です。
func main() {
	// lambda.Start(HandleRequest) を追加すると、Lambda 関数が実行されます。
	lambda.Start(HandleRequest)
}
