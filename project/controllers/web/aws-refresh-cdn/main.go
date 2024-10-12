package aws_refresh_cdn

import (
	"github.com/Lxb921006/cmdb/project/controllers/web/aws-refresh-cdn/handle"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RefreshAwsCdn struct {
	Item string `form:"item" binding:"required"`
}

func RefreshAwsCdnCtl(ctx *gin.Context) {
	var ctl RefreshAwsCdn
	if err := ctx.ShouldBind(&ctl); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	result, err := handle.NewHandle(ctl.Item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    result,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
		"data":    "",
	})
}
