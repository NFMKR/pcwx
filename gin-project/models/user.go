package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRole 用户角色
type UserRole string

const (
	RoleSuperAdmin  UserRole = "super_admin"  // 超级管理员
	RoleSecondAdmin UserRole = "second_admin" // 二级管理员
	RoleThirdAdmin  UserRole = "third_admin"  // 三级管理员
	RoleUser        UserRole = "user"         // 普通用户
)

// Permission 权限类型
type Permission string

const (
	PermissionView   Permission = "view"   // 浏览权限
	PermissionEdit   Permission = "edit"   // 编辑权限
	PermissionManage Permission = "manage" // 管理权限
	PermissionPeople Permission = "people" // 人员管理权限
)

// User 用户模型
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username    string             `bson:"username" json:"username"`
	Password    string             `bson:"password" json:"-"`                    // 密码不返回给前端
	Role        UserRole           `bson:"role" json:"role"`                     // 用户角色
	Permissions []Permission       `bson:"permissions" json:"permissions"`       // 权限列表
	ParentID    primitive.ObjectID `bson:"parent_id,omitempty" json:"parent_id"` // 上级管理员ID
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	LastLoginAt *time.Time         `bson:"last_login_at" json:"last_login_at"`
	Status      bool               `bson:"status" json:"status"` // 账号状态
}
