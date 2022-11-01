package entity

type User struct {
	ID      int64  `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
