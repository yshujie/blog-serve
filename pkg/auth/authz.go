package auth

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"

	"gorm.io/gorm"
)

const (
	// casbin 访问控制模型.
	aclModel = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)`
)

// Authz 定义了一个授权器，提供授权功能.
type Authz struct {
	*casbin.SyncedEnforcer
}

// NewAuthz 创建一个使用 casbin 完成授权的授权器.
func NewAuthz(db *gorm.DB) (*Authz, error) {
	// Initialize a Gorm adapter and use it in a Casbin enforcer
	adapter, err := adapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	m, _ := model.NewModelFromString(aclModel)

	// Initialize the enforcer.
	enforcer, err := casbin.NewSyncedEnforcer(m, adapter)
	if err != nil {
		return nil, err
	}

	// Load the policy from DB.
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	// 设置自动加载策略的间隔时间为 60 秒
	// enforcer.StartAutoLoadPolicy(60 * time.Second)

	a := &Authz{enforcer}

	return a, nil
}

// Authorize 用来进行授权.
func (a *Authz) Authorize(sub, obj, act string) (bool, error) {
	return a.Enforce(sub, obj, act)
}

// ReloadPolicy 重新加载策略
func (a *Authz) ReloadPolicy() error {
	return a.LoadPolicy()
}
