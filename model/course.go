package model

type CourseRequest struct {
	Name         string `form:"name" json:"name" binding:"required"`
	Course_Type  string `form:"course_type" json:"course_type" binding:"required"`
	Time         string `form:"time" json:"time" binding:"required"`
	Weekdays     string `form:"weekdays" json:"weekdays"`
	Start_Date   string `form:"start_date" json:"start_date" binding:"required"`
	End_Date     string `form:"end_date" json:"end_date"`
	Note         string `form:"note" json:"note"`
	Del_flag     int    `form:"del_flag" json:"del_flag"`
	Created_time string `form:"created_time" json:"created_time"`
	Created_user string `form:"created_user" json:"created_user"`
	Updated_time string `form:"updated_time" json:"updated_time"`
	Updated_user string `form:"updated_user" json:"updated_user"`
}
