package entity

import "github.com/deigo96/e-wallet.git/app/constant"

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (r *Role) TableName() string {
	return constant.TableRole
}
