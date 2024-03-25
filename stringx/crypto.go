package stringx

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func UUID() string {
	return uuid.New().String()
}

func Md5(x string) string {
	m := md5.New()
	m.Write([]byte(x))
	a := m.Sum(nil)
	return fmt.Sprintf("%x", a)
}

func Sha1(x string) string {
	m := sha1.New()
	m.Write([]byte(x))
	a := m.Sum(nil)
	return fmt.Sprintf("%x", a)
}

func Password(s string) string {
	password, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(password)
}

func ComparePassword(password string, encrypt string) bool {
	if bcrypt.CompareHashAndPassword([]byte(encrypt), []byte(password)) != nil {
		return false
	}
	return true
}

// AESEncrypt 使用 AES 加密算法对明文进行加密
func AESEncrypt(key []byte, plaintext string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建随机向量
	iv := make([]byte, aes.BlockSize)

	// 创建 AES-CBC 模式加密器
	mode := cipher.NewCBCEncrypter(block, iv)

	// 对明文进行填充
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	paddedText := append([]byte(plaintext), bytes.Repeat([]byte{byte(padding)}, padding)...)

	// 加密数据
	ciphertext := make([]byte, len(paddedText))
	mode.CryptBlocks(ciphertext, paddedText)

	return ciphertext, nil
}

// AESDecrypt 使用 AES 解密算法对密文进行解密
func AESDecrypt(key []byte, ciphertext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 创建随机向量
	iv := make([]byte, aes.BlockSize)

	// 创建 AES-CBC 模式解密器
	mode := cipher.NewCBCDecrypter(block, iv)

	// 解密数据
	decrypted := make([]byte, len(ciphertext))
	mode.CryptBlocks(decrypted, ciphertext)

	// 去除填充
	padding := int(decrypted[len(decrypted)-1])
	if padding < 1 || padding > aes.BlockSize {
		return "", errors.New("无效的填充")
	}
	decrypted = decrypted[:len(decrypted)-padding]

	return string(decrypted), nil
}
