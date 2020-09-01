package util

import (
	"math/rand"
	"net"
	"strings"
	"time"

	"github.com/gethinyan/enterprise/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString 生成随机字符串
func RandomString(n int) string {
	str := make([]rune, n)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	return string(str)
}

// RandomCode 生成随机验证码
func RandomCode() int {
	min := setting.Code.Min
	max := setting.Code.Max + 1
	return rand.Intn(max-min) + min
}

// GeneratePassword 生成加密密码
func GeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash 检查密码是否正确
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetClientIP 获取客户端 IP 地址
func GetClientIP(c *gin.Context) string {
	xForwardedFor := c.Request.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(c.Request.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// GetClientAddr 获取客户端地址
func GetClientAddr(ip string) string {
	db, err := geoip2.Open("GeoIP2-City.mmdb")
	if err != nil {
		return ""
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	parseIP := net.ParseIP(ip)
	record, err := db.City(parseIP)
	if err != nil {
		return ""
	}

	return record.City.Names["pt-BR"]
}
