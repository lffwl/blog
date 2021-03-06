// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"www.lffwl.com/app/dao/internal"
)

// wwwApiDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type wwwApiDao struct {
	*internal.WwwApiDao
}

var (
	// WwwApi is globally public accessible object for table www_api operations.
	WwwApi wwwApiDao
)

func init() {
	WwwApi = wwwApiDao{
		internal.NewWwwApiDao(),
	}
}

// Fill with you ideas below.
