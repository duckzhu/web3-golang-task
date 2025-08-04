package middlewares

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 路由中间件示例，类似JAVA的Filter或Interceptor
// 这些中间件可以在路由分组中使用，处理请求前后的逻辑
// 例如：打印请求信息、验证用户权限等
func InitMiddleware1(c *gin.Context) {
	// 中间件示例：打印请求方法
	println("Request Method:", c.Request.Method)
	// 使用这个函数后，这个中间件相当于JAVA的Filter，类似于环绕处理
	c.Next() // 调用下一个中间件或处理函数
	println("After request processing in initMiddleware1")
}

func InitMiddleware2(c *gin.Context) {
	// 中间件示例：打印请求头
	println("Request Headers:", c.Request.Header)
	// 使用这个函数后，这个中间件相当于JAVA的Filter，类似于环绕处理
	c.Next() // 调用下一个中间件或处理函数
	println("After request processing in initMiddleware2")
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从请求头中获取 Authorization 字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "未提供认证 Token"})
			c.Abort()
			return
		}

		// 2. 检查格式是否为 "Bearer <token>"
		var tokenString string
		_, err := fmt.Sscanf(authHeader, "Bearer %s", &tokenString)
		if err != nil || tokenString == "" {
			c.JSON(401, gin.H{"error": "Token 格式错误"})
			c.Abort()
			return
		}

		// 3. 解析并验证 Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法是否是 HS256
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
			}
			// 返回用于验证签名的密钥
			return []byte("123"), nil
		})

		if err != nil {
			c.JSON(401, gin.H{"error": "无效的 Token", "details": err.Error()})
			c.Abort()
			return
		}

		// 4. 检查 Token 是否有效（是否过期、签名是否通过）
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 可以在这里获取用户信息，比如用户名
			username := claims["username"].(string)
			idFloat := claims["id"].(float64)
			// id := int(idFloat)
			// 将用户信息存入 Gin 的上下文中，供后续 Handler 使用
			c.Set("username", username)
			c.Set("userid", idFloat)
			c.Next() // 继续处理请求
		} else {
			c.JSON(401, gin.H{"error": "无效的 Token"})
			c.Abort()
		}
	}
}
