package check

import (
	"os"
	"strconv"
	"time"

	"github.com/PostSample/models"
	"github.com/gin-gonic/gin"
)

func GetPath(c *gin.Context) {
	var resultErr model.ResultErr
	var req model.Check

	path := c.PostForm("path")
	timeout := c.PostForm("timeout")
	req.Path = path
	req.Timeout, _ = strconv.Atoi(timeout)

	if req.Path == "" || req.Timeout <= 0 {
		resultErr.Error = "ERR_INPUT"
		c.JSON(400, resultErr)
		return
	}

	req.Time = time.Now().Format("2006-01-02 15:04:05")

	t1 := time.Now().UnixNano() / int64(time.Millisecond)
	var t2 int64
	for i := 0; i < req.Timeout; i++ {
		_, err := os.Stat(req.Path)
		if err != nil {
			continue
		} else {
			t2 = time.Now().UnixNano() / int64(time.Millisecond)
			break
		}
	}
	if t2 == 0 {
		t2 = time.Now().UnixNano() / int64(time.Millisecond)
	}

	req.Duration = t2 - t1
	c.JSON(200, req)
	return
}
