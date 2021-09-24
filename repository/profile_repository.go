package repository

import (
	"database/sql"
	"golangapi/database"
	"golangapi/helper"
	"golangapi/model"
	"golangapi/view"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

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
	var profiles []view.Profile

	urlParams := c.Request.URL.Query()

	//Convert []string to string --> string to int
	numPageString := strings.Join(urlParams["numPage"], "")
	numRowString := strings.Join(urlParams["numRow"], "")

	// Convert string to int
	numPageInt, _ := strconv.Atoi(numPageString)
	numRowInt, _ := strconv.Atoi(numRowString)
	count := countRecords(db, c)

	maxPage := int(math.Ceil(float64(count) / float64(numRowInt)))

	//Check if urlParams empty ==> return All Profiles
	if cReq := checkRequest(numPageString, numRowString, numPageInt, numRowInt); cReq != "" {
		query := "SELECT * FROM profiles WHERE del_flag = 0 "
		rows, err := db.Query(query)
		if err != nil {
			helper.WriteLog("/profile.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": "Profiles Not Found",
			})
		}
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
			helper.WriteLog("/profile.log", err.Error(), query)
			panic(err.Error())
		}

		c.JSON(200, gin.H{
			"maxPage": maxPage,
			"result":  profiles,
		})
	} else {
		offsetInt := (numPageInt - 1) * numRowInt
		offsetString := strconv.Itoa(offsetInt)

		query := "SELECT * FROM profiles WHERE del_flag = 0 LIMIT " + offsetString + ", " + numRowString
		rows, err := db.Query(query)
		if err != nil {
			helper.WriteLog("/profile.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": "Profiles Not Found",
			})
		}
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
				helper.WriteLog("/profile.log", err.Error(), query)
				panic(err.Error())
			}
			profiles = append(profiles, pfl)
		}
		if err = rows.Err(); err != nil {
			helper.WriteLog("/profile.log", err.Error(), query)
			panic(err.Error())
		}

		c.JSON(200, gin.H{
			"maxPage": maxPage,
			"result":  profiles,
		})
	}
	defer db.Close()
}

func CreateProfile(c *gin.Context) {
	db := database.DBConn()
	var json model.ProfileRequest
	c.Request.ParseMultipartForm(8 << 20)

	if err := c.ShouldBind(&json); err == nil {
		image := imageUpload(c)

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

		//+ " WHERE NOT EXIST (SELECT id FROM profiles WHERE employee_id = '" + json.Employee_id + "')"
		insProfile, err := db.Prepare(query)
		if err != nil {
			helper.WriteLog("/profile.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
				"prepare":  "WRONG",
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
			image,
			json.Del_flag,
			getTime(),
			json.Created_user,
			json.Updated_time,
			json.Updated_user,
		)
		if err != nil {
			helper.WriteLog("/profile.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
				"exec":     "WRONG",
			})
		}

		//Return data profile
		lastRowId, err := rs.LastInsertId()
		if err != nil {
			helper.WriteLog("/profile.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
				"lastID":   "WRONG",
			})
		}

		id := strconv.FormatInt(lastRowId, 10)

		pfl := getProfile(db, id)

		c.JSON(200, gin.H{
			"messages": "Inserted",
			"data":     pfl,
		})
	} else {
		helper.WriteLog("/profile.log", err.Error(), "")
		c.JSON(500, gin.H{
			"error": err.Error(),
			"err":   "WRONG",
		})
	}

	defer db.Close()
}

func UpdateProfile(c *gin.Context) {
	db := database.DBConn()
	c.Request.ParseMultipartForm(8 << 20)

	var json model.ProfileRequest
	id := c.Param("id")
	if err := c.ShouldBind(&json); err == nil {
		image := imageUpload(c)
		pflCheck := getProfile(db, id)
		if pflCheck.Image != "" && image == "" {
			image = pflCheck.Image
		}

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
			helper.WriteLog("/profile.log", err.Error(), query)
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
			image,
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
		helper.WriteLog("/profile.log", err.Error(), "")
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
		helper.WriteLog("/profile.log", err.Error(), query)
		panic(err.Error())
	}

	delete.Exec()
	c.JSON(200, gin.H{
		"messages": "Deleted",
	})
}

func getTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

func getProfile(db *sql.DB, id string) view.Profile {
	//Return data profile
	query := "SELECT * FROM profiles WHERE id = " + id + " AND del_flag = 0"
	row, err := db.Query(query)
	if err != nil {
		helper.WriteLog("/profile.log", err.Error(), query)
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
			helper.WriteLog("/profile.log", err.Error(), query)
			panic(err.Error())
		}
	}
	return pfl
}

func countRecords(db *sql.DB, c *gin.Context) int {
	query := "SELECT COUNT(id) FROM profiles WHERE del_flag = 0"
	rows, err := db.Query(query)
	if err != nil {
		helper.WriteLog("/profile.log", err.Error(), query)
		c.JSON(500, gin.H{
			"messages": err,
		})
	}

	defer rows.Close()
	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			helper.WriteLog("/profile.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
			})
		}
	}
	return count
}

func checkRequest(numPageString string, numRowString string, numPageInt int, numRowInt int) string {
	if numPageString == "" && numRowString == "" {
		return "400 Error: Bad Request"
	}
	if numPageInt == 0 && numRowInt == 0 {
		return "numPage, numRow invalid"
	}
	if numPageInt == 0 && numRowInt != 0 {
		return "numPage invalid"
	}
	if numPageInt != 0 && numRowInt == 0 {
		return "numRow invalid"
	}
	return ""
}

func imageUpload(c *gin.Context) string {
	c.Request.ParseMultipartForm(8 << 20)

	file, handler, err := c.Request.FormFile("imageFile")
	if err != nil {
		return ""
	}
	defer file.Close()

	f, err := os.OpenFile("public/users/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return ""
	}
	defer f.Close()
	io.Copy(f, file)

	return handler.Filename
}
