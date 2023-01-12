package controller

import (
	"fmt"
	"lecture/go-swag/model"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{}
	fmt.Println(rep)
	return r, nil
}

// GetOK godoc
// @Summary call GetOK, return ok by json.
// @Description api test를 위한 기능.
// @name GetOK
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /acc/v01/ok [get]
// @Success 200 {object} Controller
func (p *Controller) GetOK(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})
	return
}
