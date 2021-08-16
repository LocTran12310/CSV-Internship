package view

type Profile struct {
	ID             int    `json:"id"`
	Employee_id    string `json:"employee_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Birthday       string `json:"birthday"`
	Position_id    int    `json:"position_id"`
	Department_id  int    `json:"department_id"`
	Status         int    `json:"status"`
	Address        string `json:"address"`
	Telephone      string `json:"telephone"`
	Mobile         string `json:"mobile"`
	Official_date  string `json:"official_date"`
	Probation_date string `json:"probation_date"`
	Gender         int    `json:"gender"`
	Image          string `json:"image"`
	Del_flag       int    `json:"del_flag"`
	Created_time   string `json:"created_time"`
	Created_user   string `json:"created_user"`
	Updated_time   string `json:"updated_time"`
	Updated_user   string `json:"updated_user"`
}
