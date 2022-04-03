// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"www.lffwl.com/app/dao/internal"
)

// wwwRoleDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type wwwRoleDao struct {
	*internal.WwwRoleDao
}

var (
	// WwwRole is globally public accessible object for table www_role operations.
	WwwRole wwwRoleDao
)

func init() {
	WwwRole = wwwRoleDao{
		internal.NewWwwRoleDao(),
	}
}

// Fill with you ideas below.
