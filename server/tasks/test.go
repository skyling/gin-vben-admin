package tasks

import (
	"context"
	"encoding/json"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/dao/query"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	TypeTest       = "test"
	TypeTestStruct = "test-struct"
)

type TestPayload struct {
	Msg string
}

// RunTestTask 运行测试案例
func RunTestTask(msg string) (*asynq.TaskInfo, error) {
	payload, err := json.Marshal(TestPayload{Msg: msg})
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(TypeTest, payload)
	return client.Enqueue(task)
}

func HandleTestTask(ctx context.Context, t *asynq.Task) error {
	var p TestPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	logrus.Infof("======== start %s %s\n", "test", p.Msg)
	time.Sleep(2 * time.Second)
	uq := query.User
	u := models.User{
		Name:     "test",
		Username: "test" + time.Now().Format("2006-01-02150405"),
	}
	uq.WithContext(ctx).Create(&u)
	logrus.Infof("======== end %s %s\n", "test", p.Msg)
	return nil
}

type TestStructPayload struct {
	Msg string
}

func NewTestStructTask(msg string) (*asynq.Task, error) {
	payload, err := json.Marshal(TestPayload{Msg: msg})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeTestStruct, payload), nil
}

type TestStructProcessor struct {
}

func (pro *TestStructPayload) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p TestStructPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	logrus.Infof("======== start %s %s\n", "test struct", p.Msg)
	time.Sleep(5 * time.Second)
	logrus.Infof("======== end %s %s\n", "test struct", p.Msg)
	return nil
}

func NewTestStructProcessor() *TestStructProcessor {
	return &TestStructProcessor{}
}
