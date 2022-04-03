package model

//添加swagger参数
type AdminPostStoreSwag struct {
	UserName  string `validate:"required"` // 用户名
	Password  string `validate:"required"` // 密码
	Password2 string `validate:"required"` // 确认密码
}

//更新swagger参数
type AdminPostUpdateSwag struct {
	UserName string `validate:"required"` // 用户名
	Password string // 密码，留空不修改
}

// 添加管理员的验证参数
type AdminPostStoreReq struct {
	AdminPostStoreInput
	UserName  string `v:"required|length:5,16"`
	Password  string `v:"required|length:6,16"`
	Password2 string `v:"required|length:6,16|same:Password"`
}

// 添加管理员的参数
type AdminPostStoreInput struct {
	AdminPostStoreUpdateBase
	CreatedId int `d:"0"` // 添加人
}

// 更新管理员参数
type AdminPostStoreUpdateBase struct {
	UserName     string
	Password     string
	HashPassword string
}

// 更新管理员的验证参数
type AdminPutUpdateReq struct {
	AdminPutUpdateInput
	Id       int    `v:"required|min:1"`
	UserName string `v:"required|length:5,16"`
	Password string `v:"length:6,16"`
}

// 更新管理员的参数
type AdminPutUpdateInput struct {
	AdminPostStoreUpdateBase
	Id        int
	UpdatedId int `d:"0"` // 更新人
}

// admin列表请求参数
type AdminGetListReq struct {
	AdminGetListInput
	Page  int `d:"1"  v:"min:0#分页号码错误"`     // 分页号码
	Limit int `d:"10" v:"max:50#分页数量最大50条"` // 分页数量，最大50
}

// admin列表查询参数
type AdminGetListInput struct {
	Page  int // 分页号码
	Limit int // 分页数量，最大50
}

// admin列表输出
type AdminGetListOutput struct {
	List  []AdminGetList `json:"list"`  // 列表
	Page  int            `json:"page"`  // 分页码
	Limit int            `json:"limit"` // 分页数量
	Total int            `json:"total"` // 数据总数
}
type AdminGetList struct {
	Id        string `json:"id"`
	UserName  string `json:"user_name"`
	CreatedId string `json:"created_id"`
	CreatedAt string `json:"created_at"`
	UpdatedId string `json:"updated_id"`
	UpdatedAt string `json:"updated_at"`
}

// admin列表查询参数
type AdminPostUpdateInput struct {
	Page  int // 分页号码
	Limit int // 分页数量，最大50
}

// 删除参数
type AdminDeleteReq struct {
	Id int `v:"required|min:1"`
}

// 设置管理员角色
type AdminPostSetRoleReq struct {
	Id    int    `v:"required|min:0" validate:"required"`                                  // ID
	Roles string `v:"regex:^(\\d+,)*\\d+$|methods|required#只能数字和逗号组合" validate:"required"` //角色Id集合
}
type AdminPostSetRoleData struct {
	RoleId  int
	AdminId int
}

//查看角色
type AdminGetRoleOutput struct {
	List    []RoleGetList `json:"list"`    // 列表
	Checked []string      `json:"checked"` // 选中
}
