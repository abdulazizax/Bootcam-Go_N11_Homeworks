package acl

import (
	"github.com/casbin/casbin/v2"
)

func NewEnforcer() (*casbin.Enforcer, error) {
	return casbin.NewEnforcer("internal/storage/acl/model.conf", "internal/storage/acl/policy.csv")
}
