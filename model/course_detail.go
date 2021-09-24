package model

type CourseParticipantsRequest struct {
	Course_Id    int    `form:"course_id" json:"course_id" binding:"required"`
	Profile_Id   int    `form:"profile_id" json:"profile_id" binding:"required"`
	Del_flag     int    `form:"del_flag" json:"del_flag"`
	Updated_time string `form:"updated_time" json:"updated_time"`
	Updated_user string `form:"updated_user" json:"updated_user"`
	Created_time string `form:"created_time" json:"created_time"`
	Created_user string `form:"created_user" json:"created_user"`
}
