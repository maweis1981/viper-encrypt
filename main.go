package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"github.com/spf13/viper"
)

// 解密函数
func decrypt(ciphertext string, key string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(data) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)

	return string(data), nil
}

func encrypt(plaintext string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func main() {
	// 读取配置文件
	viper.SetConfigName("config") // 配置文件名称（不包含扩展名）
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(".")      // 配置文件路径
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// 从配置文件中读取加密的密码
	encryptedPassword := viper.GetString("db.password")
	fmt.Println("Encrypted Password is ", encryptedPassword)

	if encryptedPassword == "" {
		log.Fatal("Encrypted password not found in config file")
	}

	// 定义解密密钥
	decryptionKey := "mysecretencryptionkey" // 确保这个密钥安全存储
	password, err := decrypt(encryptedPassword, decryptionKey)
	if err != nil {
		log.Fatalf("Error decrypting password: %s", err)
	}

	// 输出解密后的密码
	fmt.Println("Decrypted password:", password)
}
