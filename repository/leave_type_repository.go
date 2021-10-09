package repository

import (
	"database/sql"
	"golangapi/database"
	"golangapi/helper"
	"golangapi/model"
	"golangapi/view"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLeaveTypes(c *gin.Context) {
	db := database.DBConn()
	var leaveTypes []view.LeaveType

	query := "SELECT * FROM leave_type WHERE del_flag = 0 "
	rows, err := db.Query(query)
	if err != nil {
		helper.WriteLog("/leave_type.log", err.Error(), query)
		c.JSON(500, gin.H{
			"messages": "Leave Type Not Found",
		})
	}
	for rows.Next() {
		var leaveType view.LeaveType

		if err := rows.Scan(
			&leaveType.ID,
			&leaveType.Name,
			&leaveType.Del_flag,
			&leaveType.Updated_time,
			&leaveType.Updated_user,
			&leaveType.Created_time,
			&leaveType.Created_user,
		); err != nil {
			helper.WriteLog("/leave_type.log", err.Error(), query)
			panic(err.Error())
		}
		leaveTypes = append(leaveTypes, leaveType)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	c.JSON(200, gin.H{
		"message": "getLeaveTypes",
		"result":  leaveTypes,
	})
}

func CreateLeaveType(c *gin.Context) {
	db := database.DBConn()
	var json model.LeaveTypeRequest

	if err := c.ShouldBindJSON(&json); err == nil {
		query := `	INSERT INTO leave_type(
						name,
						del_flag,
						updated_time,
						updated_user,
						created_time,
						created_user
					) VALUES(?,?,?,?,?,?)`
		insItem, err := db.Prepare(query)
		if err != nil {
			helper.WriteLog("/leave_type.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
			})
		}

		rs, err := insItem.Exec(
			json.Name,
			json.Del_flag,
			json.Updated_time,
			json.Updated_user,
			getTime(),
			json.Created_user,
		)
		if err != nil {
			helper.WriteLog("/leave_type.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
				"exec":     "WRONG",
			})
		}
		//Return data profile
		lastRowId, err := rs.LastInsertId()
		if err != nil {
			helper.WriteLog("/leave_type.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
				"lastID":   "WRONG",
			})
		}

		id := strconv.FormatInt(lastRowId, 10)

		leaveType := getLeaveType(db, id)

		c.JSON(200, gin.H{
			"messages": "Inserted",
			"data":     leaveType,
		})
	} else {
		helper.WriteLog("/leave_type.log", err.Error(), "")
		c.JSON(500, gin.H{
			"error": err.Error(),
			"err":   "WRONG",
		})
	}

	defer db.Close()
}

func UpdateLeaveType(c *gin.Context) {
	db := database.DBConn()

	var json model.LeaveTypeRequest
	id := c.Param("id")

	if err := c.ShouldBindJSON(&json); err == nil {

		query := `	UPDATE leave_type SET
						name = ?,
						del_flag = ?,
						updated_time = ?,
						updated_user = ?
					WHERE id = ` + id
		edit, err := db.Prepare(query)
		if err != nil {
			helper.WriteLog("/leave_type.log", err.Error(), query)
			panic(err.Error())
		}

		edit.Exec(
			json.Name,
			json.Del_flag,
			getTime(),
			json.Updated_user,
		)
		//Return data profile
		leaveType := getLeaveType(db, id)

		c.JSON(200, gin.H{
			"messages": "Updated",
			"data":     leaveType,
		})
	} else {
		helper.WriteLog("/leave_type.log", err.Error(), "")
		c.JSON(500, gin.H{"error": err.Error()})
	}
	defer db.Close()
}

func DeleteLeaveType(c *gin.Context) {
	db := database.DBConn()

	id := c.Param("id")
	query := `	UPDATE leave_type SET
								del_flag = 1
							WHERE id = ` + id

	delete, err := db.Prepare(query)
	if err != nil {
		helper.WriteLog("/leave_type.log", err.Error(), query)
		panic(err.Error())
	}

	delete.Exec()
	c.JSON(200, gin.H{
		"messages": "Deleted",
	})
}

func getLeaveType(db *sql.DB, id string) view.LeaveType {
	//Return data profile
	query := "SELECT * FROM leave_type WHERE id = " + id + " AND del_flag = 0"
	row, err := db.Query(query)
	if err != nil {
		helper.WriteLog("/leave_type.log", err.Error(), query)
		panic(err.Error())
	}

	var item view.LeaveType
	for row.Next() {
		if err := row.Scan(
			&item.ID,
			&item.Name,
			&item.Del_flag,
			&item.Updated_time,
			&item.Updated_user,
			&item.Created_time,
			&item.Created_user,
		); err != nil {
			helper.WriteLog("/leave_type.log", err.Error(), query)
			panic(err.Error())
		}
	}
	return item
}
