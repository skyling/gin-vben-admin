package tasks

import (
	"gin-vben-admin/global"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

var (
	client *asynq.Client
	server *asynq.Server
)

func Server() {
	server = asynq.NewServer(asynq.RedisClientOpt{
		Addr:     global.Conf.Redis.Network,
		Password: global.Conf.Redis.Password,
		DB:       global.Conf.Redis.TaskDB,
	}, asynq.Config{
		Concurrency: 10, // worker 个数
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
		Logger: logrus.StandardLogger(),
	})

	mux := asynq.NewServeMux()

	for k, t := range tasks {
		mux.HandleFunc(k, t)
	}

	go func() {
		err := server.Run(mux)
		if err != nil {
			logrus.WithError(err).Error("start task server error")
		}
	}()
}

func Client() {
	client = asynq.NewClient(asynq.RedisClientOpt{
		Addr:     global.Conf.Redis.Network,
		Password: global.Conf.Redis.Password,
		DB:       global.Conf.Redis.TaskDB,
	})
}

func ClientClose() {
	if client != nil {
		client.Close()
	}
}
func ServerClose() {
	if server != nil {
		server.Stop()
	}
}
