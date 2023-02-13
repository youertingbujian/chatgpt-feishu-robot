package biz

import (
	"chatgpt/config"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func ReceiveEvent(c *gin.Context) {
	var req *ReceiveEventEncrypt
	// 结构体和传进来的json绑定
	if bind := c.ShouldBindJSON(&req); bind != nil {
		logrus.WithError(bind).Errorf("failed to read request")
		c.JSON(http.StatusOK, gin.H{
			"message": "非法访问",
		})
		return
	}
	// 事件解密
	decryptStr, err := Decrypt(req.Encrypt, config.Conf.EncryptKey)
	if err != nil {
		logrus.WithError(err).Errorf("decrypt error")
		return
	}
	//logrus.Infof("decrypt event: %v", decryptStr)

	decryptToken := &DecryptToken{}
	err = json.Unmarshal([]byte(decryptStr), decryptToken)
	if err != nil {
		logrus.Errorf("Unmarshal failed again")
		return
	}

	// 返回该challenge值作为响应。
	if decryptToken.Challenge != "" {
		c.JSON(http.StatusOK, gin.H{
			"challenge": decryptToken.Challenge,
		})
		return
	}
	
	event := &larkim.P2MessageReceiveV1{}
	err = json.Unmarshal([]byte(decryptStr), event)
	if err != nil {
		logrus.Errorf("Unmarshal failed, maybe Challenge")
		return
	}

	eventType := event.EventV2Base.Header.EventType
	//logrus.Infof("eventType: %v", eventType)
	switch eventType {
	case "im.message.receive_v1":
		if err != nil {
			logrus.Errorf("Unmarshal failed, maybe Challenge")
			return
		}
		go func() {
			err = ReceiveMessageEventHandle(event)
			if err != nil {
				logrus.WithError(err).Errorf("handle receive message event failed")
			}
		}()
	default:
		logrus.Info("unhandled event")
	}

	c.JSON(200, gin.H{
		"message": "ok",
	})
	return
}

// Decrypt 事件解密
func Decrypt(encrypt string, key string) (string, error) {
	buf, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", fmt.Errorf("base64StdEncode Error[%v]", err)
	}
	if len(buf) < aes.BlockSize {
		return "", errors.New("cipher  too short")
	}
	keyBs := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(keyBs[:sha256.Size])
	if err != nil {
		return "", fmt.Errorf("AESNewCipher Error[%v]", err)
	}
	iv := buf[:aes.BlockSize]
	buf = buf[aes.BlockSize:]
	// CBC mode always works in whole blocks.
	if len(buf)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(buf, buf)
	n := strings.Index(string(buf), "{")
	if n == -1 {
		n = 0
	}
	m := strings.LastIndex(string(buf), "}")
	if m == -1 {
		m = len(buf) - 1
	}
	return string(buf[n : m+1]), nil
}
