package model

type ProfileRequest struct {
	Employee_id    string `form:"employee_id" json:"employee_id" binding:"required"`
	Name           string `form:"name" json:"name" binding:"required"`
	Email          string `form:"email" json:"email" binding:"required"`
	Birthday       string `form:"birthday" json:"birthday"`
	Position_id    int    `form:"position_id" json:"position_id" binding:"required"`
	Department_id  int    `form:"department_id" json:"department_id" binding:"required"`
	Status         int    `form:"status" json:"status"`
	Address        string `form:"address" json:"address"`
	Telephone      string `form:"telephone" json:"telephone"`
	Mobile         string `form:"mobile" json:"mobile"`
	Official_date  string `form:"official_date" json:"official_date"`
	Probation_date string `form:"probation_date" json:"probation_date"`
	Gender         int    `form:"gender" json:"gender"`
	Image          string `form:"image" json:"image"`
	Del_flag       int    `form:"del_flag" json:"del_flag"`
	Created_time   string `form:"created_time" json:"created_time"`
	Created_user   string `form:"created_user" json:"created_user"`
	Updated_time   string `form:"updated_time" json:"updated_time"`
	Updated_user   string `form:"updated_user" json:"updated_user"`
}
