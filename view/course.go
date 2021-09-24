package view

type Course struct {
	ID           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	Course_Type  string `json:"course_type" binding:"required"`
	Time         string `json:"time"`
	Weekdays     string `json:"weekdays"`
	Start_Date   string `json:"start_date"`
	End_Date     string `json:"end_date"`
	Note         string `json:"note"`
	Del_flag     int    `json:"del_flag"`
	Created_time string `json:"created_time"`
	Created_user string `json:"created_user"`
	Updated_time string `json:"updated_time"`
	Updated_user string `json:"updated_user"`
}
