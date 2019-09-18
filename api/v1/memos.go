package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"e.coding.net/handnote/handnote/models"
	"github.com/gin-gonic/gin"
)

// ListMemo 备忘录/便笺列表
func ListMemo(c *gin.Context) {
	memos := models.GetMemoList()
	c.JSON(http.StatusOK, gin.H{"data": memos})
}

// UpdateMemo 创建备忘录/便笺
func UpdateMemo(c *gin.Context) {
	var request models.MemoRequestBody
	if err := c.Bind(&request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证失败"})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID错误"})
		return
	}
	memo := models.Memo{
		ID:      uint(id),
		Name:    request.Name,
		Content: request.Content,
	}
	if err := models.SaveMemo(&memo); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "创建备忘录/便笺失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": memo})
}
