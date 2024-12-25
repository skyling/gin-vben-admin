package tasks

import (
	"context"
	"github.com/hibiken/asynq"
)

var tasks = map[string]func(context.Context, *asynq.Task) error{
	TypeTest:             HandleTestTask,
	TypeRoleCasbinPolicy: HandleRoleCasbinPolicyTask,
	TypeAddRolesForUser:  HandleAddRolesForUserTask,
}
