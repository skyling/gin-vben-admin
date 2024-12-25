package models

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type CasbinRule struct {
	gormadapter.CasbinRule
}

var (
	// sub 主体(用户) dom: 域 租户
	CasbinModelText = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _
g2 = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && g2(r.obj, p.obj) && r.act == p.act || r.sub == "root"`
)

/**


 */
