package view

type LeaveReason struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Del_flag     int    `json:"del_flag"`
	Created_time string `json:"created_time"`
	Created_user string `json:"created_user"`
	Updated_time string `json:"updated_time"`
	Updated_user string `json:"updated_user"`
}
