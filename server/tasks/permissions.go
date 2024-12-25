package tasks

import (
	"context"
	"encoding/json"
	"gin-vben-admin/dao/repo"
	"github.com/bwmarrin/snowflake"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

const (
	TypeRoleCasbinPolicy = "role-casbin-policy"
	TypeAddRolesForUser  = "add-roles-for-user"
)

type RoleCasbinPolicyPayload struct {
	RoleID snowflake.ID `json:"role_id"`
}

// RunRoleCasbinPolicyTask 更新角色权限点
func RunRoleCasbinPolicyTask(roleID snowflake.ID) (*asynq.TaskInfo, error) {
	payload, err := json.Marshal(RoleCasbinPolicyPayload{RoleID: roleID})
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(TypeRoleCasbinPolicy, payload)
	return client.Enqueue(task)
}

func HandleRoleCasbinPolicyTask(ctx context.Context, t *asynq.Task) error {
	var p RoleCasbinPolicyPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	logrus.Infof("======== start %s %s\n", "role casbin policy", p.RoleID)
	defer logrus.Infof("======== end %s %s\n", "role casbin policy", p.RoleID)
	repo.CasbinSrv.UpdatePolicy(p.RoleID)
	repo.CasbinSrv.LoadCasbinPolicy()
	return nil
}

type AddRolesForUserPayload struct {
	UserID snowflake.ID `json:"user_id"`
}

func RunAddRolesForUserTask(userId snowflake.ID) (*asynq.TaskInfo, error) {
	payload, err := json.Marshal(AddRolesForUserPayload{UserID: userId})
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(TypeAddRolesForUser, payload)
	return client.Enqueue(task)
}

func HandleAddRolesForUserTask(ctx context.Context, t *asynq.Task) error {
	var p AddRolesForUserPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	logrus.Infof("======== start %s %s\n", "role casbin policy", p.UserID)
	defer logrus.Infof("======== end %s %s\n", "role casbin policy", p.UserID)
	u, err := repo.UserSrv.GetUserWithRoles(p.UserID)
	if err != nil {
		return err
	}
	repo.CasbinSrv.AddRolesForUser(u)
	return nil
}
