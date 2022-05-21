package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	hmacSha256Input, hmacSha256Key = "hello world", "dongle"
	hmacSha256HexExpected          = "77f5c8ce4147600543e70b12701e7b78b5b95172332ebbb06de65fcea7112179"
	hmacSha256Base32Expected       = "O724RTSBI5QAKQ7HBMJHAHT3PC23SULSGMXLXMDN4ZP45JYREF4Q===="
	hmacSha256Base64Expected       = "d/XIzkFHYAVD5wsScB57eLW5UXIzLruwbeZfzqcRIXk="
)

func TestEncrypt_ByHmacSha256_FromStringToString(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{hmacSha256Input, hmacSha256Key, hmacSha256HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input1).ByHmacSha256(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{hmacSha256Input, hmacSha256Key, hmacSha256Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input1).ByHmacSha256(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{hmacSha256Input, hmacSha256Key, hmacSha256Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input1).ByHmacSha256(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByHmacSha256_FromBytesToBytes(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(hmacSha256Input), []byte(hmacSha256Key), []byte(hmacSha256HexExpected)},
	}

	for index, test := range hexTests {
		e := Encrypt.FromBytes(test.input1).ByHmacSha256(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(hmacSha256Input), []byte(hmacSha256Key), []byte(hmacSha256Base32Expected)},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromBytes(test.input1).ByHmacSha256(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(hmacSha256Input), []byte(hmacSha256Key), []byte(hmacSha256Base64Expected)},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromBytes(test.input1).ByHmacSha256(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}