package controller

import (
	"net/http"

	"github.com/Nathanaelaxl/qresto-api/service"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	menuService service.MenuService
}

func NewMenuController(svc service.MenuService) *MenuController {

	return &MenuController{menuService: svc}

}

func (ctrl *MenuController) GetMenus(ctx *gin.Context) {

	menus, err := ctrl.menuService.GetAllMenus()

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return

	}

	ctx.JSON(http.StatusOK, gin.H{"data": menus})

}
