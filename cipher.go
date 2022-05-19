package dongle

import (
	"bytes"
	"crypto/cipher"
)

// mode constant
// 模式常量
const (
	CBC = "cbc"
	ECB = "ecb"
	CFB = "cfb"
	OFB = "ofb"
	CTR = "ctr"
	GCM = "gcm"
)

// padding constant
// 填充常量
const (
	NO    = "no"
	ZERO  = "zero"
	PKCS5 = "pkcs5"
	PKCS7 = "pkcs7"
)

// Cipher defines a Cipher struct.
// 定义 Cipher 结构体
type Cipher struct {
	mode       string // 分组模式
	padding    string // 填充模式
	key        []byte // 密钥
	iv         []byte // 偏移向量
	publicKey  []byte // 公钥
	privateKey []byte // 私钥
}

// NewCipher returns a new Cipher instance.
// 初始化 Cipher 结构体
func NewCipher() *Cipher {
	return &Cipher{
		mode:    CBC,
		padding: PKCS7,
	}
}

// SetMode sets mode.
func (c *Cipher) SetMode(mode string) {
	c.mode = mode
}

// SetPadding sets padding.
func (c *Cipher) SetPadding(padding string) {
	c.padding = padding
}

// SetKey sets key, key size must be 16, 24 or 32 bytes.
func (c *Cipher) SetKey(key interface{}) {
	switch v := key.(type) {
	case string:
		c.key = string2bytes(v)
	case []byte:
		c.key = v
	}
}

// SetIV sets iv.
func (c *Cipher) SetIV(iv interface{}) {
	switch v := iv.(type) {
	case string:
		c.iv = string2bytes(v)
	case []byte:
		c.iv = v
	}
}

// ZeroPadding padding with ZERO mode.
func (c *Cipher) ZeroPadding(src []byte, size int) []byte {
	paddingCount := size - len(src)%size
	if paddingCount == 0 {
		return src
	}
	paddingText := bytes.Repeat([]byte{byte(0)}, paddingCount)
	return append(src, paddingText...)
}

// ZeroUnPadding removes padding with ZERO mode.
func (c *Cipher) ZeroUnPadding(src []byte) []byte {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
}

// PKCS5Padding padding with PKCS5 mode.
func (c *Cipher) PKCS5Padding(src []byte) []byte {
	return c.PKCS7Padding(src, 8)
}

// PKCS5UnPadding removes padding with PKCS5 mode.
func (c *Cipher) PKCS5UnPadding(src []byte) []byte {
	return c.PKCS7UnPadding(src)
}

// PKCS7Padding padding with PKCS7 mode.
func (c *Cipher) PKCS7Padding(src []byte, size int) []byte {
	paddingCount := size - len(src)%size
	paddingText := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	return append(src, paddingText...)
}

// PKCS7UnPadding removes padding with PKCS7 mode.
func (c *Cipher) PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	trim := length - int(src[length-1])
	if trim < 0 {
		return nil
	}
	return src[:trim]
}

// CBCEncrypt encrypts with CBC mode.
func (c *Cipher) CBCEncrypt(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCBCEncrypter(block, c.iv).CryptBlocks(dst, src)
	return
}

// CBCDecrypt decrypts with CBC mode.
func (c *Cipher) CBCDecrypt(dst []byte, block cipher.Block) (src []byte) {
	src = make([]byte, len(dst))
	cipher.NewCBCDecrypter(block, c.iv).CryptBlocks(src, dst)
	return
}

// CFBEncrypt encrypts with CFB mode.
func (c *Cipher) CFBEncrypt(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCFBEncrypter(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// CFBDecrypt encrypts with CFB mode.
func (c *Cipher) CFBDecrypt(dst []byte, block cipher.Block) (src []byte) {
	src = make([]byte, len(dst))
	cipher.NewCFBDecrypter(block, c.iv[:block.BlockSize()]).XORKeyStream(src, dst)
	return
}

// CTREncrypt encrypts with CTR mode.
func (c *Cipher) CTREncrypt(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCTR(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// CTRDecrypt encrypts with CTR mode.
func (c *Cipher) CTRDecrypt(dst []byte, block cipher.Block) (src []byte) {
	src = make([]byte, len(dst))
	cipher.NewCTR(block, c.iv[:block.BlockSize()]).XORKeyStream(src, dst)
	return
}
