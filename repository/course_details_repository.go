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

func GetCourseParticipants(c *gin.Context) {
	db := database.DBConn()
	var courseParticipants []view.CourseParticipants

	id := c.Param("id")
	query := `SELECT
							course_details.id, profiles.name, profiles.employee_id,
							courses.name, courses.course_type, courses.id, profiles.id, courses.end_date, courses.weekdays, courses.time,
							course_details.del_flag, course_details.updated_time,
							course_details.updated_user, course_details.created_time, course_details.created_user
						FROM course_details, profiles, courses
						WHERE course_details.course_id = courses.id
							AND course_details.profile_id = profiles.id
							AND course_details.del_flag = 0
							AND courses.id = ` + id
	rows, err := db.Query(query)
	if err != nil {
		helper.WriteLog("/course_details.log", err.Error(), query)
		c.JSON(500, gin.H{
			"messages": "Course Details Not Found",
		})
	}
	for rows.Next() {
		var courseParticipant view.CourseParticipants

		if err := rows.Scan(
			&courseParticipant.ID,
			&courseParticipant.Employee_Name,
			&courseParticipant.Employee_Id,
			&courseParticipant.Course_Name,
			&courseParticipant.Course_Type,
			&courseParticipant.Course_Id,
			&courseParticipant.Profile_Id,
			&courseParticipant.End_Date,
			&courseParticipant.Weekdays,
			&courseParticipant.Time,
			&courseParticipant.Del_flag,
			&courseParticipant.Updated_time,
			&courseParticipant.Updated_user,
			&courseParticipant.Created_time,
			&courseParticipant.Created_user,
		); err != nil {
			helper.WriteLog("/course_details.log", err.Error(), query)
			panic(err.Error())
		}
		courseParticipants = append(courseParticipants, courseParticipant)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	c.JSON(200, gin.H{
		"message": "getCourseParticipants",
		"result":  courseParticipants,
	})
}

func CreateCourseParticipant(c *gin.Context) {
	db := database.DBConn()
	var json model.CourseParticipantsRequest

	if err := c.ShouldBindJSON(&json); err == nil {
		// Exist if > -1
		// If exist ==> update del_flag = 0
		if idParticipantExist := checkParticipantExist(c, db, json); idParticipantExist > 0 {
			id := strconv.Itoa(idParticipantExist)

			query := `	UPDATE course_details SET
									del_flag = 0
								WHERE id = ` + id

			update, err := db.Prepare(query)
			if err != nil {
				helper.WriteLog("/course_details.log", err.Error(), query)
				panic(err.Error())
			}

			update.Exec()

			courseParticipant := getCourseParticipant(db, id)

			c.JSON(200, gin.H{
				"messages": "UPDATE",
				"data":     courseParticipant,
			})
		} else { // If not exist ==> INSERT
			query := `	INSERT INTO course_details (
							course_id,
							profile_id,
							del_flag,
							updated_time,
							updated_user,
							created_time,
							created_user
						) VALUES(?,?,?,?,?,?,?)`

			insCourseParticipant, err := db.Prepare(query)
			if err != nil {
				helper.WriteLog("/course_details.log", err.Error(), query)
				c.JSON(500, gin.H{
					"messages": err,
					"prepare":  "WRONG",
				})
			}

			rs, err := insCourseParticipant.Exec(
				json.Course_Id,
				json.Profile_Id,
				json.Del_flag,
				getTime(),
				json.Created_user,
				json.Updated_time,
				json.Updated_user,
			)
			if err != nil {
				helper.WriteLog("/course_details.log", err.Error(), query)
				c.JSON(500, gin.H{
					"messages": err,
					"exec":     "WRONG",
				})
			}

			//Return data
			lastRowId, err := rs.LastInsertId()
			if err != nil {
				helper.WriteLog("/course_details.log", err.Error(), query)
				c.JSON(500, gin.H{
					"messages": err,
					"lastID":   "WRONG",
				})
			}

			id := strconv.FormatInt(lastRowId, 10)

			courseParticipant := getCourseParticipant(db, id)

			c.JSON(200, gin.H{
				"messages": "Inserted",
				"data":     courseParticipant,
			})
		}
	} else {
		helper.WriteLog("/course_details.log", err.Error(), "")
		c.JSON(500, gin.H{
			"error": err.Error(),
			"err":   "WRONG",
		})
	}
}

func DeleteCourseParticipant(c *gin.Context) {
	db := database.DBConn()
	id := c.Param("id")
	query := `	UPDATE course_details SET
								del_flag = 1
							WHERE id = ` + id

	delete, err := db.Prepare(query)
	if err != nil {
		helper.WriteLog("/course_details.log", err.Error(), query)
		panic(err.Error())
	}

	delete.Exec()
	c.JSON(200, gin.H{
		"messages": "Deleted",
	})
}

//Return 0 if not exist
func checkParticipantExist(c *gin.Context, db *sql.DB, json model.CourseParticipantsRequest) int {
	query := `SELECT id FROM course_details
						WHERE course_id = ?
						AND profile_id = ?`

	row, err := db.Query(query, json.Course_Id, json.Profile_Id)
	if err != nil {
		helper.WriteLog("/course_details.log", err.Error(), query)
		panic(err.Error())
	}

	var id int
	for row.Next() {
		if err := row.Scan(
			&id,
		); err != nil {
			helper.WriteLog("/course_details.log", err.Error(), query)
			panic(err.Error())
		}
	}

	return id
}

func getCourseParticipant(db *sql.DB, id string) view.CourseParticipants {
	//Return data profile
	query := `SELECT
							course_details.id, profiles.name, profiles.employee_id,
							courses.name, courses.course_type, courses.id, profiles.id, courses.end_date, courses.weekdays, courses.time,
							course_details.del_flag, course_details.updated_time,
							course_details.updated_user, course_details.created_time, course_details.created_user
						FROM course_details, profiles, courses
						WHERE course_details.course_id = courses.id
							AND course_details.profile_id = profiles.id
							AND course_details.id = ` + id
	row, err := db.Query(query)
	if err != nil {
		helper.WriteLog("/course_details.log", err.Error(), query)
		panic(err.Error())
	}

	var courseParticipant view.CourseParticipants
	for row.Next() {
		if err := row.Scan(
			&courseParticipant.ID,
			&courseParticipant.Employee_Name,
			&courseParticipant.Employee_Id,
			&courseParticipant.Course_Name,
			&courseParticipant.Course_Type,
			&courseParticipant.Course_Id,
			&courseParticipant.Profile_Id,
			&courseParticipant.End_Date,
			&courseParticipant.Weekdays,
			&courseParticipant.Time,
			&courseParticipant.Del_flag,
			&courseParticipant.Updated_time,
			&courseParticipant.Updated_user,
			&courseParticipant.Created_time,
			&courseParticipant.Created_user,
		); err != nil {
			helper.WriteLog("/course_details.log", err.Error(), query)
			panic(err.Error())
		}
	}
	return courseParticipant
}
