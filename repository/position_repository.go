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

func GetPositions(c *gin.Context) {
	db := database.DBConn()
	var items []view.Position

	query := "SELECT * FROM positions WHERE del_flag = 0 "
	rows, err := db.Query(query)
	if err != nil {
		helper.WriteLog("/position.log", err.Error(), query)
		c.JSON(500, gin.H{
			"messages": "Positions Not Found",
		})
	}
	for rows.Next() {
		var item view.Position

		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Short_name,
			&item.Note,
			&item.Del_flag,
			&item.Updated_time,
			&item.Updated_user,
			&item.Created_time,
			&item.Created_user,
		); err != nil {
			helper.WriteLog("/position.log", err.Error(), query)
			panic(err.Error())
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	c.JSON(200, gin.H{
		"message": "GetPositions",
		"result":  items,
	})
}

func CreatePosition(c *gin.Context) {
	db := database.DBConn()
	var json model.PositionRequest

	if err := c.ShouldBindJSON(&json); err == nil {
		query := `	INSERT INTO positions(
						name,
						short_name,
						note,
						del_flag,
						updated_time,
						updated_user,
						created_time,
						created_user
					) VALUES(?,?,?,?,?,?,?,?)`
		insPosition, err := db.Prepare(query)
		if err != nil {
			helper.WriteLog("/position.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
			})
		}

		rs, err := insPosition.Exec(
			json.Name,
			json.Short_name,
			json.Note,
			json.Del_flag,
			json.Updated_time,
			json.Updated_user,
			getTime(),
			json.Created_user,
		)
		if err != nil {
			helper.WriteLog("/positions.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
				"exec":     "WRONG",
			})
		}
		//Return data profile
		lastRowId, err := rs.LastInsertId()
		if err != nil {
			helper.WriteLog("/positions.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
				"lastID":   "WRONG",
			})
		}

		id := strconv.FormatInt(lastRowId, 10)

		position := getPosition(db, id)

		c.JSON(200, gin.H{
			"messages": "Inserted",
			"data":     position,
		})
	}
}

func UpdatePosition(c *gin.Context) {
	db := database.DBConn()

	var json model.PositionRequest
	id := c.Param("id")

	if err := c.ShouldBindJSON(&json); err == nil {

		query := `	UPDATE positions SET
						name = ?,
						short_name = ?,
						note = ?,
						del_flag = ?,
						updated_time = ?,
						updated_user = ?
					WHERE id = ` + id
		edit, err := db.Prepare(query)
		if err != nil {
			helper.WriteLog("/position.log", err.Error(), query)
			panic(err.Error())
		}

		edit.Exec(
			json.Name,
			json.Short_name,
			json.Note,
			json.Del_flag,
			getTime(),
			json.Updated_user,
		)
		//Return data profile
		position := getPosition(db, id)

		c.JSON(200, gin.H{
			"messages": "Updated",
			"data":     position,
		})
	} else {
		helper.WriteLog("/position.log", err.Error(), "")
		c.JSON(500, gin.H{"error": err.Error()})
	}
	defer db.Close()
}
func DeletePosition(c *gin.Context) {
	db := database.DBConn()
	id := c.Param("id")
	query := `	UPDATE positions SET
								del_flag = 1
							WHERE id = ` + id

	delete, err := db.Prepare(query)
	if err != nil {
		helper.WriteLog("/position.log", err.Error(), query)
		panic(err.Error())
	}

	delete.Exec()
	c.JSON(200, gin.H{
		"messages": "Deleted",
	})
}

func getPosition(db *sql.DB, id string) view.Position {
	//Return data profile
	query := "SELECT * FROM positions WHERE id = " + id + " AND del_flag = 0"
	row, err := db.Query(query)
	if err != nil {
		helper.WriteLog("/position.log", err.Error(), query)
		panic(err.Error())
	}

	var item view.Position
	for row.Next() {
		if err := row.Scan(
			&item.ID,
			&item.Name,
			&item.Short_name,
			&item.Note,
			&item.Del_flag,
			&item.Updated_time,
			&item.Updated_user,
			&item.Created_time,
			&item.Created_user,
		); err != nil {
			helper.WriteLog("/position.log", err.Error(), query)
			panic(err.Error())
		}
	}
	return item
}
