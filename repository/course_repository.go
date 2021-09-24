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

func GetCourse(c *gin.Context) {
	db := database.DBConn()
	id := c.Param("id")
	course := getCourse(db, id)
	c.JSON(200, gin.H{
		"messages": "GetCourse",
		"data":     course,
	})
	defer db.Close()
}

func GetCourses(c *gin.Context) {
	db := database.DBConn()
	var courses []view.Course

	query := "SELECT * FROM courses WHERE del_flag = 0 "
	rows, err := db.Query(query)
	if err != nil {
		helper.WriteLog("/course.log", err.Error(), query)
		c.JSON(500, gin.H{
			"messages": "Courses Not Found",
		})
	}
	for rows.Next() {
		var course view.Course

		if err := rows.Scan(
			&course.ID,
			&course.Name,
			&course.Course_Type,
			&course.Time,
			&course.Weekdays,
			&course.Start_Date,
			&course.End_Date,
			&course.Note,
			&course.Del_flag,
			&course.Updated_time,
			&course.Updated_user,
			&course.Created_time,
			&course.Created_user,
		); err != nil {
			helper.WriteLog("/course.log", err.Error(), query)
			panic(err.Error())
		}
		courses = append(courses, course)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	c.JSON(200, gin.H{
		"message": "getCourses",
		"result":  courses,
	})
}

func CreateCourse(c *gin.Context) {
	db := database.DBConn()
	var json model.CourseRequest

	if err := c.ShouldBindJSON(&json); err == nil {
		query := `	INSERT INTO courses(
						name,
						course_type,
						time,
						weekdays,
						start_date,
						end_date,
						note,
						del_flag,
						updated_time,
						updated_user,
						created_time,
						created_user
					) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)`
		insCourse, err := db.Prepare(query)
		if err != nil {
			helper.WriteLog("/course.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
			})
		}

		rs, err := insCourse.Exec(
			json.Name,
			json.Course_Type,
			json.Time,
			json.Weekdays,
			json.Start_Date,
			json.End_Date,
			json.Note,
			json.Del_flag,
			json.Updated_time,
			json.Updated_user,
			getTime(),
			json.Created_user,
		)
		if err != nil {
			helper.WriteLog("/course.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
				"exec":     "WRONG",
			})
		}
		//Return data profile
		lastRowId, err := rs.LastInsertId()
		if err != nil {
			helper.WriteLog("/course.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
				"lastID":   "WRONG",
			})
		}

		id := strconv.FormatInt(lastRowId, 10)

		course := getCourse(db, id)

		c.JSON(200, gin.H{
			"messages": "Inserted",
			"data":     course,
		})
	} else {
		helper.WriteLog("/course.log", err.Error(), "")
		c.JSON(500, gin.H{
			"error": err.Error(),
			"err":   "WRONG",
		})
	}

	defer db.Close()
}

func UpdateCourse(c *gin.Context) {
	db := database.DBConn()

	var json model.CourseRequest
	id := c.Param("id")

	if err := c.ShouldBindJSON(&json); err == nil {

		query := `	UPDATE courses SET
						name = ?,
						course_type = ?,
						time = ?,
						weekdays = ?,
						start_date = ?,
						end_date = ?,
						note = ?,
						del_flag = ?,
						updated_time = ?,
						updated_user = ?
					WHERE id = ` + id
		edit, err := db.Prepare(query)
		if err != nil {
			helper.WriteLog("/course.log", err.Error(), query)
			panic(err.Error())
		}

		edit.Exec(
			json.Name,
			json.Course_Type,
			json.Time,
			json.Weekdays,
			json.Start_Date,
			json.End_Date,
			json.Note,
			json.Del_flag,
			getTime(),
			json.Updated_user,
		)
		//Return data profile
		course := getCourse(db, id)

		c.JSON(200, gin.H{
			"messages": "Updated",
			"data":     course,
		})
	} else {
		helper.WriteLog("/course.log", err.Error(), "")
		c.JSON(500, gin.H{"error": err.Error()})
	}
	defer db.Close()
}

func DeleteCourse(c *gin.Context) {
	db := database.DBConn()
	id := c.Param("id")
	query := `	UPDATE courses SET
								del_flag = 1
							WHERE id = ` + id

	delete, err := db.Prepare(query)
	if err != nil {
		helper.WriteLog("/course.log", err.Error(), query)
		panic(err.Error())
	}

	delete.Exec()
	c.JSON(200, gin.H{
		"messages": "Deleted",
	})
}

func getCourse(db *sql.DB, id string) view.Course {
	//Return data profile
	query := "SELECT * FROM courses WHERE id = " + id + " AND del_flag = 0"
	row, err := db.Query(query)
	if err != nil {
		helper.WriteLog("/course.log", err.Error(), query)
		panic(err.Error())
	}

	var item view.Course
	for row.Next() {
		if err := row.Scan(
			&item.ID,
			&item.Name,
			&item.Course_Type,
			&item.Time,
			&item.Weekdays,
			&item.Start_Date,
			&item.End_Date,
			&item.Note,
			&item.Del_flag,
			&item.Updated_time,
			&item.Updated_user,
			&item.Created_time,
			&item.Created_user,
		); err != nil {
			helper.WriteLog("/course.log", err.Error(), query)
			panic(err.Error())
		}
	}
	return item
}
