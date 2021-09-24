package view

type Department struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Short_name   string `json:"short_name"`
	Note         string `json:"note"`
	Del_flag     int    `json:"del_flag"`
	Updated_time string `json:"updated_time"`
	Updated_user string `json:"updated_user"`
	Created_time string `json:"created_time"`
	Created_user string `json:"created_user"`
}
