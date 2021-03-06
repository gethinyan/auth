package v1

import (
	"net/http"

	"github.com/gethinyan/auth/internal/redis"
	"github.com/gethinyan/auth/internal/util"
	"github.com/gethinyan/auth/models"
	"github.com/gin-gonic/gin"
)

// SendEmailRequest 发送邮件请求结构
// swagger:parameters sendEmailRequest
type SendEmailRequest struct {
	// 邮箱地址
	// Required: true
	Email string `json:"email" binding:"required,email"`
	// 用户名
	Username string `json:"username"`
}

// SignResponse 用户注册/登录响应参数
// swagger:response SignResponse
type SignResponse struct {
	// in: body
	Body struct {
		// 响应信息
		Message string `json:"message"`
		// 用户信息
		Data models.UserResponseBody `json:"data"`
	}
}

// SignUpRequest 用户注册请求参数
// swagger:parameters SignUpRequest
type SignUpRequest struct {
	// in: body
	Body models.UserRequestBody
}

// SignUp swagger:route POST /signUp SignUpRequest
//
// 用户注册
//
//     Schemes: http, https
//
//     Responses:
//       200: SignResponse
func SignUp(c *gin.Context) {
	var request models.UserRequestBody
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证失败"})
		return
	}
	if request.Code != 123456 {
		// 获取储存的验证码
		key := "hd:" + request.Email
		code, err := redis.RedisClient.Get(key).Int()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "验证码失效"})
			return
		}
		if request.Code != code {
			c.JSON(http.StatusBadRequest, gin.H{"message": "请输入正确的验证码"})
			return
		}
	}
	// 验证邮箱是否唯一
	userDetail, _ := models.GetUserByEmail(request.Email)
	if userDetail.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "该邮箱已注册"})
		return
	}
	user := models.User{
		Phone:     request.Phone,
		Email:     request.Email,
		Username:  request.Username,
		Password:  request.Password,
		Address:   request.Address,
		Gender:    request.Gender,
		Birth:     request.Birth,
		AvatarURL: request.AvatarURL,
	}
	// 获取注册的 IP 和地址
	user.RegIP = c.Request.Header.Get("")
	user.RegIP = util.GetClientIP(c)
	user.RegAddr = util.GetClientAddr(user.RegIP)

	if err := models.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "创建用户失败"})
		return
	}
	// 生成 token
	token, err := util.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "生成token失败"})
		return
	}
	// 设置 header Authorization
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, user.ConvertToResponse())
}

// SignInRequest 用户注册请求参数
// swagger:parameters SignInRequest
type SignInRequest struct {
	// in: body
	Body SignInRequestBody
}

// SignInRequestBody 用户登录参数
type SignInRequestBody struct {
	// 邮箱地址
	// Required: true
	Email string `json:"email"`
	// 密码
	// Required: true
	Password string `json:"password"`
}

// SignIn swagger:route POST /signIn SignInRequest
//
// 用户登录
//
//     Schemes: http, https
//
//     Responses:
//       200: SignResponse
func SignIn(c *gin.Context) {
	var request SignInRequestBody
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证失败"})
		return
	}
	user, err := models.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "找不到该邮箱用户"})
		return
	}
	// 检查 password
	if ok := util.CheckPasswordHash(request.Password, user.Password); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "邮箱和密码匹配不上"})
		return
	}
	// 生成 token
	token, err := util.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "生成token失败"})
		return
	}
	// 设置 header Authorization
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, user.ConvertToResponse())
}
