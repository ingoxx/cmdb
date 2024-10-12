package model

import (
	"time"

	"github.com/Lxb921006/cmdb/project/dao"
)

type CronsCrontabs struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	OperateUser string    `gorm:"not null" json:"operate_user"`
	Mission     string    `gorm:"not null" json:"mission"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Status      uint      `gorm:"default:100;comment:'100-完成,101-未开始,102-失败,103-停止,104-运行中'" json:"status"`
	Project     string    `gorm:"not null" json:"project"`
}

func (cm *CronsCrontabs) CronChangeStatus(id, status uint) (err error) {
	tx := dao.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Model(cm).Where("id = ?", id).Update("status = ", status).Error; err != nil {
		tx.Rollback()
		return
	}

	return tx.Commit().Error
}
