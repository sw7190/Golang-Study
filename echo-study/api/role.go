package api

import (
	"echo-study/db"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// RoleStruct role table
type RoleStruct struct {
	RoleID   string `db:"role_id" json:"role_id"`
	RoleName string `db:"role_name" json:"role_name"`
}

// getRole
func getRole(c echo.Context) error {
	data := RoleStruct{}
	roleID := c.Param("role_id")

	err := db.DB.Get(&data, "SELECT role_id, role_name FROM role WHERE role_id = $1", roleID)

	if nil != err {
		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  fmt.Sprintf("failed %v", err),
			Internal: err,
		}
	}

	return c.JSON(http.StatusOK, db.HTTPResponseSingle{
		Data: data,
	})
}

// getRoles
func getRoles(c echo.Context) error {
	data := []RoleStruct{}

	err := db.DB.Select(&data, "SELECT role_id, role_name FROM role")

	if nil != err {
		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  fmt.Sprintf("failed %v", err),
			Internal: err,
		}
	}

	return c.JSON(http.StatusOK, db.HTTPResponseSingle{
		Data: data,
	})
}

func createRole(c echo.Context) error {
	reqBody := db.HTTPRequest{}
	if err := c.Bind(&reqBody); err != nil || reqBody.Data == nil {
		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  "Bad Request",
			Internal: err,
		}
	}
	return c.JSON(http.StatusOK, reqBody.Data)
}
