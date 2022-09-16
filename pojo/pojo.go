package pojo

import (
	"gorm.io/gorm"
)

type WorkHours struct {
	gorm.Model
	WorkContent   string
	WorkHour      float64
	ProjectId     uint
	Demander      uint
	DemanderSide  uint
	AddHour       float64
	Level         uint
	Workday       string
	UserId        uint
	DemandList    string
	Code          string
	TaskModuleId  uint
	SubBusinessId uint
	JiraId        string
	IsCheck       int8
	Remark        string
}

type APIWorkHours struct {
	ID       uint
	WorkHour float64
	WorkDay  string
	UserId   uint
}

type SysUsers struct {
	gorm.Model
	UUID          string
	Username      string
	Password      string
	NickName      string
	HeaderImg     string
	AuthorityId   string
	SideMode      string
	ActiveColor   string
	BaseColor     string
	UserType      uint
	Email         string
	Phone         string
	DepartmentId  uint
	WorkdayNotice int8
	Status        int8
}

type APISysUsers struct {
	ID            uint
	Username      string
	NickName      string
	WorkdayNotice int8
	Status        int8
}
type APIData struct {
	ID            uint
	Username      string
	NickName      string
	WorkHour      float64
	WorkdayNotice int8
	Status        int8
}

type GroupMember struct {
	Username string
	NickName string
}
