package main

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/ar3s3ru/todo-list-app/lib/must"
)

func run() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return fmt.Errorf("grpc-gateway: failed to create logger, %v", err)
	}

	//nolint:errcheck // No need for this error to come up if it happens.
	defer logger.Sync()

	return nil
}

func main() {
	must.NotFail(run())
}
