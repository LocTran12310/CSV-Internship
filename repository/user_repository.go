package repository

import (
	"database/sql"
	"golangapi/database"
	"golangapi/helper"
	"golangapi/model"
	"golangapi/view"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	db := database.DBConn()
	var json model.UserLoginRequest

	if err := c.ShouldBindJSON(&json); err == nil {
		query := `	SELECT * FROM users
					WHERE login_id = ?`
		row, err := db.Query(query, json.Login_id)
		if err != nil {
			helper.WriteLog("/user.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
			})
		}

		var user view.User
		for row.Next() {
			if err := row.Scan(
				&user.ID,
				&user.Login_id,
				&user.Contract_type_id,
				&user.Profile_id,
				&user.Password,
				&user.Del_flag,
				&user.Created_time,
				&user.Created_user,
				&user.Updated_time,
				&user.Updated_user,
			); err != nil {
				helper.WriteLog("/user.log", err.Error(), query)
				c.JSON(500, gin.H{
					"messages": "Table User have NULL field",
				})
				return
			}
		}

		//Check for no results
		if user.ID == 0 {
			helper.WriteLog("/user.log", "User Not Found", query)
			c.JSON(404, gin.H{
				"messages": "User not found",
			})
			return
		}

		// Comparing the password with the hash
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password)); err != nil {
			// TODO: Properly handle error
			helper.WriteLog("/user.log", err.Error(), query)
			c.JSON(400, gin.H{
				"messages": "Password is wrong",
			})
			return
		}
		id := strconv.Itoa(user.Profile_id)
		pfl := getProfile(db, id)

		session := sessions.Default(c)
		session.Set("id", user.ID)
		session.Set("login_id", user.Login_id)
		session.Save()

		c.JSON(200, gin.H{
			"messages":  "Login Successfully",
			"user":      user,
			"profile":   pfl,
			"json.ID":   json.Login_id,
			"json.Pass": json.Password,
		})
	}

	defer db.Close()
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{
		"message": "Logout Successfully",
	})
}

func CreateUser(c *gin.Context) {
	db := database.DBConn()
	var json model.UserRequest

	if err := c.ShouldBindJSON(&json); err == nil {
		query := `	INSERT INTO users(
						login_id,
						contract_type_id,
						profile_id,
						password,
						del_flag,
						created_time,
						created_user,
						updated_time,
						updated_user
					) VALUES(?,?,?,?,?,?,?,?,?)`
		insUser, err := db.Prepare(query)
		if err != nil {
			helper.WriteLog("/user.log", err.Error(), query)
			c.JSON(500, gin.H{
				"messages": err,
			})
		}

		// Generate "hash" to store from user password
		hash, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
		if err != nil {
			helper.WriteLog("/user.log", err.Error(), query)
			// TODO: Properly handle error
			panic(err.Error())
		}
		rs, err := insUser.Exec(
			json.Login_id,
			json.Contract_type_id,
			json.Profile_id,
			hash,
			json.Del_flag,
			getTime(),
			json.Created_user,
			json.Updated_time,
			json.Updated_user,
		)
		if err != nil {
			helper.WriteLog("/user.log", err.Error(), query)

			c.JSON(500, gin.H{
				"messages": err,
			})
		}
		//Return user
		lastRowId, err := rs.LastInsertId()
		if err != nil {
			helper.WriteLog("/user.log", err.Error(), query)

			c.JSON(500, gin.H{
				"messages": err,
			})
		}

		id := strconv.FormatInt(lastRowId, 10)

		user := getUser(db, c, id)

		c.JSON(200, gin.H{
			"messages": "Inserted",
			"data":     user,
		})
	} else {
		helper.WriteLog("/user.log", err.Error(), "")

		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	defer db.Close()
}
func ChangePassword(c *gin.Context) {
	db := database.DBConn()
	var json model.ChangePasswordRequest
	id := c.Param("id")
	if err := c.ShouldBindJSON(&json); err == nil {
		query := `UPDATE users SET password = ? WHERE id = ` + id
		edit, err := db.Prepare(query)
		if err != nil {
			helper.WriteLog("/user.log", err.Error(), query)
			panic(err.Error())
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
		if err != nil {
			helper.WriteLog("/user.log", err.Error(), query)
			panic(err.Error())
		}
		edit.Exec(hash)
		user := getUser(db, c, id)
		c.JSON(200, gin.H{
			"messages": "ChangePassword",
			"data":     user,
		})
	}
}

func getUser(db *sql.DB, c *gin.Context, id string) view.User {
	//Return user
	query := "SELECT * FROM users WHERE id = ?"
	row, err := db.Query(query, id)
	if err != nil {
		helper.WriteLog("/user.log", err.Error(), query)
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	var user view.User
	for row.Next() {
		if err := row.Scan(
			&user.ID,
			&user.Login_id,
			&user.Contract_type_id,
			&user.Profile_id,
			&user.Password,
			&user.Del_flag,
			&user.Created_time,
			&user.Created_user,
			&user.Updated_time,
			&user.Updated_user,
		); err != nil {
			helper.WriteLog("/user.log", err.Error(), query)
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
	}
	return user
}
