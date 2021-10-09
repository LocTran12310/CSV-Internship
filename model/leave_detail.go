package model

// type LeaveDetailArray struct {
// 	LeaveDetail []LeaveDetailRequest `json: "leave_detail_array"`
// }
type LeaveDetailArray []struct {
	Employee_id     string `form:"employee_id" json:"employee_id" binding:"required"`
	User_id         int    `form:"user_id" json:"user_id" binding:"required"`
	Start_date      string `form:"start_date" json:"start_date" binding:"required"`
	End_date        string `form:"end_date" json:"end_date" binding:"required"`
	Leave_type_id   int    `form:"leave_type_id" json:"leave_type_id" binding:"required"`
	Leave_reason_id int    `form:"leave_reason_id" json:"leave_reason_id" binding:"required"`
	Note            string `form:"note" json:"note" `
	Del_flag        int    `form:"del_flag" json:"del_flag"`
	Updated_time    string `form:"updated_time" json:"updated_time"`
	Updated_user    string `form:"updated_user" json:"updated_user"`
	Created_time    string `form:"created_time" json:"created_time"`
	Created_user    string `form:"created_user" json:"created_user"`
}

type LeaveDetailRequest struct {
	Employee_id     int    `form:"employee_id" json:"employee_id" `
	User_id         int    `form:"user_id" json:"user_id"`
	Start_date      string `form:"start_date" json:"start_date" binding:"required"`
	End_date        string `form:"end_date" json:"end_date" binding:"required"`
	Leave_type_id   int    `form:"leave_type_id" json:"leave_type_id" binding:"required"`
	Leave_reason_id int    `form:"leave_reason_id" json:"leave_reason_id" binding:"required"`
	Note            string `form:"note" json:"note" `
	Del_flag        int    `form:"del_flag" json:"del_flag"`
	Updated_time    string `form:"updated_time" json:"updated_time"`
	Updated_user    string `form:"updated_user" json:"updated_user"`
	Created_time    string `form:"created_time" json:"created_time"`
	Created_user    string `form:"created_user" json:"created_user"`
}
