package api

import (
	"echo-study/db"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// TestStruct role table
type TestStruct struct {
	RoleID   string `db:"role_id" json:"role_id"`
	RoleName string `db:"role_name" json:"role_name"`
}

// GetList test
func GetList(c echo.Context) error {
	data := []TestStruct{}

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
