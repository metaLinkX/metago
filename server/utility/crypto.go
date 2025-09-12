package utility

import (
	"github.com/forgoer/openssl"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/crypto/gsha1"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// AesECBEncrypt 加密
func AesECBEncrypt(src, key []byte) (dst []byte, err error) {
	return openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
}

func AesCBCEncrypt(src, key, iv []byte) (dst []byte, err error) {
	return openssl.AesCBCEncrypt(src, key, iv, openssl.PKCS7_PADDING)
}

// AesECBDecrypt 解密
func AesECBDecrypt(src, key []byte) (dst []byte, err error) {
	return openssl.AesECBDecrypt(src, key, openssl.PKCS7_PADDING)
}

func EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(gsha1.Encrypt(password + "@" + salt))
}

func CheckPassword(inputPassword, dbPassword, salt string) (err error) {
	inputPassword = EncryptPassword(inputPassword, salt)
	if inputPassword != dbPassword {
		err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "密码或账号不正确")
		return err
	}
	return nil
}
