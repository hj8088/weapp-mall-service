package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type APIHandlers struct {
}

func (api APIHandlers) GetBanners(c *gin.Context) {

	reqSchema := "http"
	if c.Request.TLS != nil {
		reqSchema = "https"
	}
	resourceUrl := fmt.Sprintf("%s://%s/resource", reqSchema, c.Request.Host)

	createdAt, _ := time.Parse("2006-01-02 15:04:05", "2019-02-26 08:31:40")
	updatedAt, _ := time.Parse("2006-01-02 15:04:05", "22020-01-07 08:19:29")

	banners := []Banner{
		{
			BusinessId: 0,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
			ID:         19180,
			Order:      10,
			PicUrl:     resourceUrl + "/2020/01/07/c3183558-55e8-4589-9ea7-81670004941a.jpg",
			Status:     0,
			StatusStr:  "显示",
			Title:      "启动图1",
			Type:       "app",
			UserId:     951,
		},
		{
			BusinessId: 0,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
			ID:         19181,
			Order:      20,
			PicUrl:     resourceUrl + "/2020/01/07/b7471e4f-dc0d-4245-afb7-b64fbd9bb29a.jpg",
			Status:     0,
			StatusStr:  "显示",
			Title:      "启动图2",
			Type:       "app",
			UserId:     951,
		},
		{
			BusinessId: 0,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
			ID:         19182,
			Order:      30,
			PicUrl:     resourceUrl + "/2020/01/07/5425289c-fb82-4ab4-a193-933ddba71496.jpg",
			Status:     0,
			StatusStr:  "显示",
			Title:      "启动图2",
			Type:       "app",
			UserId:     951,
		},
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"data": banners,
	})

}
