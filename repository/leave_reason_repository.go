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

func GetLeaveReasons(c *gin.Context) {
	db := database.DBConn()
	var leaveReasons []view.LeaveReason

	query := "SELECT * FROM leave_reason WHERE del_flag = 0 "
	rows, err := db.Query(query)
	if err != nil {
		helper.WriteLog("/leave_reason.log", err.Error(), query)
		c.JSON(500, gin.H{
			"messages": "Leave Reason Not Found",
		})
	}
	for rows.Next() {
		var leaveReason view.LeaveReason

		if err := rows.Scan(
			&leaveReason.ID,
			&leaveReason.Name,
			&leaveReason.Del_flag,
			&leaveReason.Updated_time,
			&leaveReason.Updated_user,
			&leaveReason.Created_time,
			&leaveReason.Created_user,
		); err != nil {
			helper.WriteLog("/leave_reason.log", err.Error(), query)
			panic(err.Error())
		}
		leaveReasons = append(leaveReasons, leaveReason)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	c.JSON(200, gin.H{
		"message": "getLeaveReasons",
		"result":  leaveReasons,
	})
}

func CreateLeaveReason(c *gin.Context) {
	db := database.DBConn()
	var json model.LeaveReasonRequest

	if err := c.ShouldBindJSON(&json); err == nil {
		query := `	INSERT INTO leave_reason(
						name,
						del_flag,
						updated_time,
						updated_user,
						created_time,
						created_user
					) VALUES(?,?,?,?,?,?)`
		insItem, err := db.Prepare(query)
		if err != nil {
			helper.WriteLog("/leave_reason.log", err.Error(), query)
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
			helper.WriteLog("/leave_reason.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
				"exec":     "WRONG",
			})
		}
		//Return data profile
		lastRowId, err := rs.LastInsertId()
		if err != nil {
			helper.WriteLog("/leave_reason.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
				"lastID":   "WRONG",
			})
		}

		id := strconv.FormatInt(lastRowId, 10)

		leaveReason := getLeaveReason(db, id)

		c.JSON(200, gin.H{
			"messages": "Inserted",
			"data":     leaveReason,
		})
	} else {
		helper.WriteLog("/leave_reason.log", err.Error(), "")
		c.JSON(500, gin.H{
			"error": err.Error(),
			"err":   "WRONG",
		})
	}

	defer db.Close()
}

func UpdateLeaveReason(c *gin.Context) {
	db := database.DBConn()

	var json model.LeaveReasonRequest
	id := c.Param("id")

	if err := c.ShouldBindJSON(&json); err == nil {

		query := `	UPDATE leave_reason SET
						name = ?,
						del_flag = ?,
						updated_time = ?,
						updated_user = ?
					WHERE id = ` + id
		edit, err := db.Prepare(query)
		if err != nil {
			helper.WriteLog("/leave_reason.log", err.Error(), query)
			panic(err.Error())
		}

		edit.Exec(
			json.Name,
			json.Del_flag,
			getTime(),
			json.Updated_user,
		)
		//Return data profile
		leaveReason := getLeaveReason(db, id)

		c.JSON(200, gin.H{
			"messages": "Updated",
			"data":     leaveReason,
		})
	} else {
		helper.WriteLog("/leave_reason.log", err.Error(), "")
		c.JSON(500, gin.H{"error": err.Error()})
	}
	defer db.Close()
}

func DeleteLeaveReason(c *gin.Context) {
	db := database.DBConn()

	id := c.Param("id")
	query := `	UPDATE leave_reason SET
								del_flag = 1
							WHERE id = ` + id

	delete, err := db.Prepare(query)
	if err != nil {
		helper.WriteLog("/leave_reason.log", err.Error(), query)
		panic(err.Error())
	}

	delete.Exec()
	c.JSON(200, gin.H{
		"messages": "Deleted",
	})
}

func getLeaveReason(db *sql.DB, id string) view.LeaveReason {
	//Return data profile
	query := "SELECT * FROM leave_reason WHERE id = " + id + " AND del_flag = 0"
	row, err := db.Query(query)
	if err != nil {
		helper.WriteLog("/leave_reason.log", err.Error(), query)
		panic(err.Error())
	}

	var item view.LeaveReason
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
			helper.WriteLog("/leave_reason.log", err.Error(), query)
			panic(err.Error())
		}
	}
	return item
}
