package v1

import (
	"fmt"
	"net/http"

	"e.coding.net/handnote/handnote/models"
	"e.coding.net/handnote/handnote/pkg/util"
	"github.com/gin-gonic/gin"
)

// ListMemo 备忘录/便笺列表
// ListMemo swagger:route GET /memos ListMemoRequest
//
// 备忘录/便笺列表
//
//     Schemes: http, https
//
//     Responses:
//       200: ListMemoResponse
func ListMemo(c *gin.Context) {
	memos := models.GetMemoList(util.UID)
	c.JSON(http.StatusOK, gin.H{"data": memos})
}

// ListMemoResponse 备忘录/便笺列表响应参数
// swagger:response ListMemoResponse
type ListMemoResponse struct {
	// in: body
	Body struct {
		// 响应信息
		Message string `json:"message"`
		// 备忘录/便笺列表
		Data []models.Memo `json:"data"`
	}
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
	// in: body
	Body struct {
		// 响应信息
		Message string `json:"message"`
		// 备忘录/便笺列表
		Data []models.Memo `json:"data"`
	}
}

// SyncMemo 同步备忘录/便笺
// SyncMemo swagger:route POST /syncMemo SyncMemoRequest
//
// 同步备忘录/便笺
//
//     Schemes: http, https
//
//     Responses:
//       200: SyncMemoResponse
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
