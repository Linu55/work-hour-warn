package dao

import (
	"work-hour-warn/conf"
	"work-hour-warn/enity"
)

func GetLazyGuy(usernames []string, dateTime string) ([]enity.APIData, []enity.GroupMember) {
	db := conf.SqlSession
	var sysUsers enity.SysUsers
	var aPIData []enity.APIData
	var gms []enity.GroupMember
	//SELECT s.id,s.username,s.nick_name,sum(w.work_hour) work_hour,s.workday_notice,s.`status` FROM `work_hours` w INNER JOIN sys_users s ON w.user_id=s.id WHERE w.work_day='?' GROUP BY id HAVING sum(w.work_hour)<7.50
	db.Model(&sysUsers).Select("sys_users.id,sys_users.username,sys_users.nick_name,sum(work_hours.work_hour) work_hour,sys_users.workday_notice,sys_users.`status`").
		Joins("INNER JOIN `work_hours` ON work_hours.user_id=sys_users.id").
		Where("work_hours.work_day = ? ", dateTime).
		Where("sys_users.username IN ?", usernames).
		Group("sys_users.id").
		Scan(&aPIData)
	db.Model(&sysUsers).Select("username,nick_name").Where("username IN ?", usernames).Scan(&gms)
	return aPIData, gms
}
