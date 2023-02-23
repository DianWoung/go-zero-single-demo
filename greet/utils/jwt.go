package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

/**
 * jwt版本
 * 1）版本<v4.0.0 https://github.com/dgrijalva/jwt-go
 * 2）版本>v4.0.0 https://github.com/golang-jwt/jwtv4
 * xxx
 *
claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
      "user": "zhangshan",
   })
newWithClaims会返回一个Token结构体，而这个token结构体有以下属性

type Token struct {
   Raw       string        //原始令牌
   Method    SigningMethod   // 加密方法 比如sha256加密
   Header    map[string]interface{} // token头信息
   Claims    Claims  // 加密配置，比如超时时间等
   Signature string  // 加密后的字符串
   Valid     bool   // 是否校验
}
我们可以通过该结构体获取到加密后的字符串信息。

接下来我们需要讲解一下Claims该结构体存储了token字符串的超时时间等信息以及在解析时的Token校验工作。

type Claims interface {
   Valid() error
}
//实现类有MapClaims、RegisteredClaims、StandardClaims(舍弃)
//其实后两个结构体都是根据MapClaims编写而来，所以我们只需要掌握MapClaims即可
type MapClaims map[string]interface{}
————————————————
版权声明：本文为CSDN博主「迷茫路人」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/a1023934860/article/details/125161365
*/

/*
 *  @Description: jwt加密
 *  @param jwtSecret "jwt密钥"
 *  @param data "待加密数据"
 *  @return string "jwt token"
 *  @return error
 */
func GetJwtToken(jwtSecret string, data map[string]interface{}) (string, error) {
	//var iat int64 = time.Now().Unix()

	// 创建Token结构体；MapClaims是一个map[string]interface{},用来保存一些数据，和Payload对应。
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": data, //数据
		//"exp":  iat + c.Auth.AccessExpire, //过期时间
		//"iat":  iat,                       //当前时间
	})

	/**
	 *1) newWithClaims会返回一个Token结构体，而这个token结构体有以下属性
	 *2) 我们可以通过该结构体获取到加密后的字符串信息。
		type Token struct {
		   Raw       string        //原始令牌
		   Method    SigningMethod   // 加密方法 比如sha256加密
		   Header    map[string]interface{} // token头信息
		   Claims    Claims  // 加密配置，比如超时时间等
		   Signature string  // 加密后的字符串
		   Valid     bool   // 是否校验
		}
	*/

	// 调用加密方法，发挥Token字符串
	tokenSecret, err := tokens.SignedString([]byte(jwtSecret))
	return tokenSecret, err
}

/*
 *  @Description: jwt解密
 *  @param jwtSecret "jwt密钥"
 *  @param encryptToken "jwt token"
 *  @return jwt.MapClaims
 *  @return error
 */
func ParseToken(jwtSecret string, encryptToken string) (jwt.MapClaims, error) {
	//ParseWithClaims用来解析加密Token 返回token结构体
	token, err := jwt.ParseWithClaims(encryptToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	fmt.Println("是否正确")
	fmt.Println(token.Valid)
	fmt.Printf("%#v", token.Claims)
	/**
	 * 1) ParseWithClaims 会返回一个Token结构体
	 * 2) token.Claims 实现了上面Valid()方法，该方法里面实现了对token过期日期校验、发布时间、生效时间的校验工作。
	 * 3) 所以在map里面有三个固定的键我们可以根据需要进行设置，exp 过期时间、iat 发布时间、nbf 生效时间
	 */
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("couldn't handle this token")
}
