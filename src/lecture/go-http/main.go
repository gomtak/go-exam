package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//해당 uri에 대한 전체 binding 되어 person 함수로 c *gin.Context를 통해 전달됨
	r.GET("/:name/:age", get)
	// r.POST("/:name/:age", post)
	// 비즈니스로직 상 err return : 호출은 정상적이나 로직상 에러
	if err != nil {
		c.JSON(204, gin.H{ // 204 or http.StatusNoContent
			"message": "Fail Verified"})
	} else {
		c.JSON(200, gin.H{ // 200 or http.StatusOK
			"message": "ok",
			"token":   AccessToken,
			"body":    Body,
		})
	}

	return
}
