package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"e.coding.net/handnote/handnote/models"
	"e.coding.net/handnote/handnote/pkg/util"
	"github.com/gin-gonic/gin"
)

// ListMemo 备忘录/便笺列表
func ListMemo(c *gin.Context) {
	memos := models.GetMemoList(util.UID)
	c.JSON(http.StatusOK, gin.H{"data": memos})
}

// SyncMemoRequestForm 同步备忘录/便笺表单
type SyncMemoRequestForm struct {
	// 新增备忘录/便笺
	Add []models.MemoRequestBody `json:"add"`
	// 修改备忘录/便笺
	Update []models.MemoRequestBody `json:"update"`
	// 恢复备忘录/便笺
	Restore []models.MemoRequestBody `json:"restore"`
	// 归档备忘录/便笺
	Archive []models.MemoRequestBody `json:"archive"`
	// 删除备忘录/便笺
	Delete []models.MemoRequestBody `json:"delete"`
}

// SyncMemoRequest 同步备忘录/便笺请求参数
// swagger:parameters SyncMemoRequest
type SyncMemoRequest struct {
	// in: body
	Body SyncMemoRequestForm
}

// SyncMemoResponse 同步备忘录/便笺响应参数
// swagger:response SyncMemoResponse
type SyncMemoResponse struct {
}

// SyncMemo 同步备忘录/便笺
func SyncMemo(c *gin.Context) {
	var request SyncMemoRequestForm
	if err := c.Bind(&request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证失败"})
		return
	}
	// 新增备忘录/便笺
	if len(request.Add) > 0 {
		models.InsertBatchMemo(request.Add)
	}
	// 更新备忘录/便笺
	if len(request.Update) > 0 {
		models.UpdateBatchMemo(request.Update)
	}
	// 恢复备忘录/便笺
	if len(request.Restore) > 0 {
		models.RestoreBatchMemo(request.Restore)
	}
	// 归档备忘录/便笺
	if len(request.Archive) > 0 {
		models.ArchiveBatchMemo(request.Archive)
	}
	// 删除备忘录/便笺
	if len(request.Delete) > 0 {
		models.DeleteBatchMemo(request.Delete)
	}
	memos := models.GetMemoList(util.UID)
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
