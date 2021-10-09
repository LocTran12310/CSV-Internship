package repository

import (
	// "database/sql"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	// "fmt"
	"golangapi/database"
	"golangapi/helper"
	"golangapi/model"
	"golangapi/view"
	"io/ioutil"

	// "strconv"

	"github.com/gin-gonic/gin"
)

// func GetLeaveDetail(c *gin.Context) {
// 	db := database.DBConn()

// 	user_id := c.Param("user_id")
// 	leaveDetails := getLeaveDetails(db, user_id)

// 	c.JSON(200, gin.H{
// 		"messages": "GetLeaveDetails",
// 		"data":     leaveDetails,
// 	})
// 	defer db.Close()
// }

func GetAllLeaveDetails(c *gin.Context) {
	db := database.DBConn()
	var leaveDetails []view.LeaveDetail

	query := "SELECT * FROM leave_details WHERE del_flag = 0"
	rows, err := db.Query(query)
	if err != nil {
		helper.WriteLog("/leave_detail.log", err.Error(), query)
		c.JSON(500, gin.H{
			"messages": "Leave Details Not Found",
		})
	}
	for rows.Next() {
		var item view.LeaveDetail

		if err := rows.Scan(
			&item.ID,
			&item.User_id,
			&item.Start_date,
			&item.End_date,
			&item.Leave_type_id,
			&item.Leave_reason_id,
			&item.Note,
			&item.Del_flag,
			&item.Updated_time,
			&item.Updated_user,
			&item.Created_time,
			&item.Created_user,
		); err != nil {
			helper.WriteLog("/leave_detail.log", err.Error(), query)
			panic(err.Error())
		}
		leaveDetails = append(leaveDetails, item)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	c.JSON(200, gin.H{
		"message": "getLeaveDetails",
		"result":  leaveDetails,
	})
}

func GetLeaveDetails(c *gin.Context) {
	db := database.DBConn()
	var leaveDetails []view.LeaveDetail

	login_id := c.Param("employee_id")

	query := `SELECT id FROM users WHERE del_flag = 0 AND login_id = ?`
	row, err := db.Query(query, login_id)
	if err != nil {
		helper.WriteLog("/leave_detail.log", err.Error(), query)
		panic(err.Error())
	}

	var userId int
	for row.Next() {
		if err := row.Scan(
			&userId,
		); err != nil {
			helper.WriteLog("/leave_detail.log", err.Error(), query)
			panic(err.Error())
		}
	}
	user_id := strconv.Itoa(userId)

	query = "SELECT * FROM leave_details WHERE del_flag = 0 AND user_id = ?"
	rows, err := db.Query(query, user_id)
	if err != nil {
		helper.WriteLog("/leave_detail.log", err.Error(), query)
		c.JSON(500, gin.H{
			"messages": "Leave Details Not Found",
		})
	}
	for rows.Next() {
		var item view.LeaveDetail

		if err := rows.Scan(
			&item.ID,
			&item.User_id,
			&item.Start_date,
			&item.End_date,
			&item.Leave_type_id,
			&item.Leave_reason_id,
			&item.Note,
			&item.Del_flag,
			&item.Updated_time,
			&item.Updated_user,
			&item.Created_time,
			&item.Created_user,
		); err != nil {
			helper.WriteLog("/leave_detail.log", err.Error(), query)
			panic(err.Error())
		}
		leaveDetails = append(leaveDetails, item)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	c.JSON(200, gin.H{
		"message": "getLeaveDetails",
		"result":  leaveDetails,
	})
}

func CreateLeaveDetails(c *gin.Context) {
	db := database.DBConn()

	body, _ := ioutil.ReadAll(c.Request.Body)
	var data model.LeaveDetailArray

	if err := json.Unmarshal([]byte(body), &data); err != nil {
		helper.WriteLog("/leave_detail.log", err.Error(), "")
		c.JSON(500, gin.H{
			"messages": err,
		})
	}

	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
		obj := data[i]

		//Get user_id by employee_id
		loginId := obj.Employee_id
		query := "SELECT id FROM users WHERE login_id = ?"
		row, err := db.Query(query, loginId)
		if err != nil {
			helper.WriteLog("/leave_detail.log", err.Error(), query)
			panic(err.Error())
		}

		var userId int
		for row.Next() {
			if err := row.Scan(
				&userId,
			); err != nil {
				helper.WriteLog("/leave_detail.log", err.Error(), query)
				panic(err.Error())
			}
		}

		// var obj model.LeaveDetailRequest
		query = `	INSERT INTO leave_details(
							user_id,
							start_date,
							end_date,
							leave_type_id,
							leave_reason_id,
							note,
							del_flag,
							updated_time,
							updated_user,
							created_time,
							created_user
						) VALUES(?,?,?,?,?,?,?,?,?,?,?)`
		insItem, err := db.Prepare(query)
		if err != nil {
			helper.WriteLog("/leave_detail.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
			})
		}

		rs, err := insItem.Exec(
			userId,
			obj.Start_date,
			obj.End_date,
			obj.Leave_type_id,
			obj.Leave_reason_id,
			obj.Note,
			obj.Del_flag,
			obj.Updated_time,
			obj.Updated_user,
			getTime(),
			obj.Created_user,
		)
		if err != nil {
			helper.WriteLog("/leave_detail.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
				"exec":     "WRONG",
			})
		}
		rs.RowsAffected()
	}
	c.JSON(200, gin.H{
		"messages": "Inserted",
	})

	defer db.Close()
}

func UpdateLeaveDetails(c *gin.Context) {
	db := database.DBConn()

	var json model.LeaveDetailRequest
	id := c.Param("id")

	if err := c.ShouldBindJSON(&json); err == nil {

		query := `	UPDATE leave_details SET
						start_date = ?,
						end_date = ?,
						leave_type_id = ?,
						leave_reason_id = ?,
						note = ?,
						del_flag = ?,
						updated_time = ?,
						updated_user = ?
					WHERE id = ` + id
		edit, err := db.Prepare(query)
		if err != nil {
			helper.WriteLog("/leave_detail.log", err.Error(), query)
			panic(err.Error())
		}

		edit.Exec(
			json.Start_date,
			json.End_date,
			json.Leave_type_id,
			json.Leave_reason_id,
			json.Note,
			json.Del_flag,
			getTime(),
			json.Updated_user,
		)
		//Return data profile
		leaveDetails := getLeaveDetails(db, id)

		c.JSON(200, gin.H{
			"messages": "Updated",
			"data":     leaveDetails,
		})
	} else {
		helper.WriteLog("/leave_detail.log", err.Error(), "")
		c.JSON(500, gin.H{"error": err.Error()})
	}
	defer db.Close()
}

func DeleteLeaveDetails(c *gin.Context) {
	db := database.DBConn()
	id := c.Param("id")
	query := `	UPDATE leave_details SET
								del_flag = 1
							WHERE id = ` + id

	delete, err := db.Prepare(query)
	if err != nil {
		helper.WriteLog("/leave_detail.log", err.Error(), query)
		panic(err.Error())
	}

	delete.Exec()
	c.JSON(200, gin.H{
		"messages": "Deleted",
	})
}

func getLeaveDetails(db *sql.DB, id string) view.LeaveDetail {
	query := "SELECT * FROM leave_details WHERE id = ? AND del_flag = 0"
	row, err := db.Query(query, id)
	if err != nil {
		helper.WriteLog("/leave_detail.log", err.Error(), query)
		panic(err.Error())
	}

	var item view.LeaveDetail
	for row.Next() {
		if err := row.Scan(
			&item.ID,
			&item.User_id,
			&item.Start_date,
			&item.End_date,
			&item.Leave_type_id,
			&item.Leave_reason_id,
			&item.Note,
			&item.Del_flag,
			&item.Updated_time,
			&item.Updated_user,
			&item.Created_time,
			&item.Created_user,
		); err != nil {
			helper.WriteLog("/leave_detail.log", err.Error(), query)
			panic(err.Error())
		}
	}
	return item
}
