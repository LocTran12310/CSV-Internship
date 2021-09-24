package view

type CourseParticipants struct {
	ID            int    `json:"id"`
	Employee_Id   string `json:"employee_id"`
	Employee_Name string `json:"employee_name"`
	Course_Id     int    `json:"course_id" binding:"required"`
	Profile_Id    int    `json:"profile_id" binding:"required"`
	Course_Name   string `json:"course_name" binding:"required"`
	Course_Type   string `json:"course_type" binding:"required"`
	Time          string `json:"time"`
	Weekdays      string `json:"weekdays"`
	End_Date      string `json:"end_date"`
	Del_flag      int    `json:"del_flag"`
	Created_time  string `json:"created_time"`
	Created_user  string `json:"created_user"`
	Updated_time  string `json:"updated_time"`
	Updated_user  string `json:"updated_user"`
}
