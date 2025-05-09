package http

import (
	"time"

	"github.com/gin-gonic/gin"

	apiv1 "github.com/LiangNing7/miniblog/pkg/api/apiserver/v1"
)

// Healthz 服务健康检查.
func (h *Handler) Healthz(c *gin.Context) {
	// 返回 JSON 响应.
	c.JSON(200, &apiv1.HealthzResponse{
		Status:    apiv1.ServiceStatus_Healthy,
		Timestamp: time.Now().Format(time.DateTime),
	})
}
