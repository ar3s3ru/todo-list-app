package main

import (
	"fmt"

	"go.uber.org/zap"

	"bitbucket.org/chronomics/todo-list-app/lib/must"
)

func run() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return fmt.Errorf("grpc-gateway: failed to create logger, %v", err)
	}

	defer logger.Sync()

	return nil
}

func main() {
	must.NotFail(run())
}
