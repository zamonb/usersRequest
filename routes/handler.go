package routes

import (
	"github.com/gin-gonic/gin"
	"usersRequest/models"
	"time"
	"net/http"
	"usersRequest/logger"
	"strconv"
)


func userReq(c *gin.Context) {
	userId := c.Query("user_id")
	user := models.User{
		Id:   userId,
		Time: time.Now(),
	}
	user.Add()
	logger.Info.Println("usr", user.Time.Format("02.01.2006 15:04:05.000000"), user.Id)

	c.Status(http.StatusOK)
}


func numberOfReq(c *gin.Context) {

	timeStart := time.Now()
	count := 0

	var m map[string]int
	m = make(map[string]int)

	for _, user := range models.UserSlice {
		if time.Now().Sub(user.Time).Minutes() < 1 {

			if m[user.Id] >= 0 {
				m[user.Id] ++
			}

			if m[user.Id] > 100 {
				count ++
				m[user.Id] = -1 //чтобы в дальнейшем  не инкрементировать для текущего user
			}
		}
	}

	logger.Info.Println("cnt", timeStart.Format("02.01.2006 15:04:05.000000"), count, "sliceNumber=" + strconv.Itoa(len(models.UserSlice)))

	c.JSON(http.StatusOK, count)

}