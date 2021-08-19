package view

type User struct {
	ID               int    `json:"id"`
	Login_id         string `json:"login_id"`
	Contract_type_id int    `json:"contract_type_id"`
	Profile_id       int    `json:"profile_id"`
	Password         string `json:"password"`
	Del_flag         int    `json:"del_flag"`
	Updated_time     string `json:"updated_time"`
	Updated_user     string `json:"updated_user"`
	Created_time     string `json:"created_time"`
	Created_user     string `json:"created_user"`
}
