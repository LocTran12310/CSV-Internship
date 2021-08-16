package repository

import (
	"database/sql"
	"golangapi/database"
	"golangapi/model"
	"golangapi/view"
	"strconv"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	db := database.DBConn()
	id := c.Param("id")
	pfl := getProfile(db, id)
	c.JSON(200, gin.H{
		"messages": "GetProfile",
		"data":     pfl,
	})
	defer db.Close()
}

func GetProfiles(c *gin.Context) {
	db := database.DBConn()
	query := "SELECT * FROM profiles"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Profiles Not Found",
		})
	}

	var profiles []view.Profile

	for rows.Next() {
		var pfl view.Profile
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

func CreateProfile(c *gin.Context) {
	db := database.DBConn()
	var json model.ProfileRequest

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

		rs, err := insProfile.Exec(
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
		if err != nil {
			c.JSON(500, gin.H{
				"messages": err,
			})
		}

		//Return data profile
		lastRowId, err := rs.LastInsertId()
		if err != nil {
			c.JSON(500, gin.H{
				"messages": err,
			})
		}

		id := strconv.FormatInt(lastRowId, 10)

		pfl := getProfile(db, id)

		c.JSON(200, gin.H{
			"messages": "Inserted",
			"data":     pfl,
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

	var json model.ProfileRequest
	id := c.Param("id")
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
					WHERE id = ` + id
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

		//Return data profile
		pfl := getProfile(db, id)

		c.JSON(200, gin.H{
			"messages": "Updated",
			"data":     pfl,
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	defer db.Close()
}

func DeleteProfile(c *gin.Context) {
	db := database.DBConn()
	id := c.Param("id")
	query := `	UPDATE profiles SET
					del_flag = 1
				WHERE id = ` + id

	delete, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	delete.Exec()
	pfl := getProfile(db, id)
	c.JSON(200, gin.H{
		"messages": "Deleted",
		"data":     pfl,
	})
}

func getTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

func getProfile(db *sql.DB, id string) view.Profile {
	//Return data profile
	query := "SELECT * FROM profiles WHERE id = " + id
	row, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	var pfl view.Profile
	for row.Next() {
		if err := row.Scan(
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
	}
	return pfl
}
