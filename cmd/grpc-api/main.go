package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"bitbucket.org/chronomics/todo-list-app/gen/todolist/v1/todolistv1connect"
	"bitbucket.org/chronomics/todo-list-app/internal/command"
	"bitbucket.org/chronomics/todo-list-app/internal/domain/todolist"
	"bitbucket.org/chronomics/todo-list-app/internal/infrastructure/grpc"
	"bitbucket.org/chronomics/todo-list-app/internal/query"
	"bitbucket.org/chronomics/todo-list-app/lib/ddd/redisddd"
	"bitbucket.org/chronomics/todo-list-app/lib/must"
)

func run() error {
	config, err := ParseConfig()
	if err != nil {
		return fmt.Errorf("grpc-api: failed to parse config, %v", err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return fmt.Errorf("grpc-api: failed to create logger, %v", err)
	}

	defer logger.Sync()

	redisClient := redis.NewClient(&redis.Options{
		Addr:       config.RedisAddr(),
		ClientName: "todo-list-app/grpc-api",
		Password:   config.Redis.Password,
	})

	if cmd := redisClient.Ping(context.Background()); cmd.Err() != nil {
		return fmt.Errorf("grpc-api: failed to ping redis, %v", cmd)
	}

	logger.Sugar().Debugw("connection with Redis established",
		"address", config.RedisAddr(),
	)

	todoListRepository := redisddd.Repository[uuid.UUID, *todolist.TodoList]{
		Prefix: "TodoList",
		Client: redisClient,
	}

	todoListService := &grpc.TodoListService{
		GenerateIDFunc: uuid.New,
		GetTodoListHandler: query.GetTodoListHandler{
			Repository: todoListRepository,
		},
		CreateTodoListHandler: command.CreateTodoListHandler{
			Clock:      time.Now,
			Repository: todoListRepository,
		},
		AddItemToTodoListHandler: command.AddItemToTodoListHandler{
			Clock:      time.Now,
			Repository: todoListRepository,
		},
		ToggleTodoItemHandler: command.ToggleTodoItemHandler{
			Repository: todoListRepository,
		},
	}

	mux := http.NewServeMux()

	mux.Handle(todolistv1connect.NewTodoListServiceHandler(todoListService))
	mux.Handle(grpchealth.NewHandler(grpchealth.NewStaticChecker(todolistv1connect.TodoListServiceName)))
	mux.Handle(grpcreflect.NewHandlerV1(grpcreflect.NewStaticReflector(todolistv1connect.TodoListServiceName)))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(grpcreflect.NewStaticReflector(todolistv1connect.TodoListServiceName)))

	logger.Sugar().Infow("grpc server started",
		"address", config.Server.Address,
	)

	// TODO: implement graceful shutdown
	err = http.ListenAndServe(
		config.Server.Address,
		h2c.NewHandler(mux, &http2.Server{}),
	)

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("grpc-api: grpc server exited with error, %v", err)
	}

	return nil
}

func main() {
	must.NotFail(run())
}
