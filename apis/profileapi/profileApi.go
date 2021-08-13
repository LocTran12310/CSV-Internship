package profileapi

import (
	"golangapi/database"
	"golangapi/entities"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllProfiles(c *gin.Context) {
	db := database.DBConn()
	query := "SELECT * FROM profiles"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Profiles Not Found",
		})
	}
	defer rows.Close() // Hoan lai row Close cho den khi func GetList thuc hien xong
	var profiles []entities.Profile

	for rows.Next() {
		var pfl entities.Profile
		if err := rows.Scan(
			&pfl.ID,
			&pfl.Employee_id,
			&pfl.Name,
			&pfl.Email,
			&pfl.Birthday,
			&pfl.Position_id,
			&pfl.Department_id,
			&pfl.Status,
			&pfl.Address,
			&pfl.Telephone,
			&pfl.Mobile,
			&pfl.Official_date,
			&pfl.Probation_date,
			&pfl.Gender,
			&pfl.Image,
			&pfl.Del_flag,
			&pfl.Created_time,
			&pfl.Created_user,
			&pfl.Updated_time,
			&pfl.Updated_user,
		); err != nil {
			panic(err.Error())
		}
		profiles = append(profiles, pfl)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	c.IndentedJSON(http.StatusOK, profiles)
	defer db.Close()
}

func AddProfile(c *gin.Context) {
	db := database.DBConn()
	var json entities.CRUDProfile

	if err := c.ShouldBindJSON(&json); err == nil {
		query := `	INSERT INTO profiles(
						employee_id,
						name,
						email,
						birthday,
						position_id,
						department_id,
						status,
						address,
						telephone,
						mobile,
						official_date,
						probation_date,
						gender,
						image,
						del_flag,
						created_time,
						created_user,
						updated_time,
						updated_user
					) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		insProfile, err := db.Prepare(query)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": err,
			})
		}

		insProfile.Exec(
			json.Employee_id,
			json.Name,
			json.Email,
			json.Birthday,
			json.Position_id,
			json.Department_id,
			json.Status,
			json.Address,
			json.Telephone,
			json.Mobile,
			json.Official_date,
			json.Probation_date,
			json.Gender,
			json.Image,
			json.Del_flag,
			getTime(),
			json.Created_user,
			json.Updated_time,
			json.Updated_user,
		)
		c.JSON(200, gin.H{
			"messages": "Inserted",
			"json":     insProfile,
		})
	} else {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	defer db.Close()
}

func UpdateProfile(c *gin.Context) {
	db := database.DBConn()

	var json entities.CRUDProfile

	if err := c.ShouldBindJSON(&json); err == nil {
		query := `	UPDATE profiles SET
						employee_id = ?,
						name = ?,
						email = ?,
						birthday = ?,
						position_id = ?,
						department_id = ?,
						status = ?,
						address = ?,
						telephone = ?,
						mobile = ?,
						official_date = ?,
						probation_date = ?,
						gender = ?,
						image = ?,
						del_flag = ?,
						updated_time = ?,
						updated_user = ?
					WHERE id = ` + c.Param("id")
		edit, err := db.Prepare(query)
		if err != nil {
			panic(err.Error())
		}

		edit.Exec(
			json.Employee_id,
			json.Name,
			json.Email,
			json.Birthday,
			json.Position_id,
			json.Department_id,
			json.Status,
			json.Address,
			json.Telephone,
			json.Mobile,
			json.Official_date,
			json.Probation_date,
			json.Gender,
			json.Image,
			json.Del_flag,
			getTime(),
			json.Updated_user,
		)

		c.JSON(200, gin.H{
			"messages": "Updated",
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	defer db.Close()
}

func getTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}
