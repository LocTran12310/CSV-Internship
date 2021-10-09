package view

type LeaveDetail struct {
	ID              int    `json:"id"`
	User_id         int    `json:"user_id"`
	Start_date      string `json:"start_date"`
	End_date        string `json:"end_date"`
	Leave_type_id   int    `json:"leave_type_id"`
	Leave_reason_id int    `json:"leave_reason_id"`
	Note            string `json:"note" `
	Del_flag        int    `json:"del_flag"`
	Updated_time    string `json:"updated_time"`
	Updated_user    string `json:"updated_user"`
	Created_time    string `json:"created_time"`
	Created_user    string `json:"created_user"`
}
