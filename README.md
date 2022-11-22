# dongle  #
[![Carbon Release](https://img.shields.io/github/release/golang-module/dongle.svg)](https://github.com/golang-module/dongle/releases)
[![Go Build](https://github.com/golang-module/dongle/actions/workflows/test.yml/badge.svg)](https://github.com/golang-module/dongle/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/dongle)](https://goreportcard.com/report/github.com/golang-module/dongle)
[![codecov](https://codecov.io/gh/golang-module/dongle/branch/main/graph/badge.svg)](https://codecov.io/gh/golang-module/dongle)
[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/dongle)
![License](https://img.shields.io/github/license/golang-module/dongle)

English | [简体中文](README.cn.md)

### Introduction
A simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption

`Dongle` has been included by [awesome-go](https://github.com/avelino/awesome-go#security "awesome-go"), if you think
it is helpful, please give me a star

[github.com/golang-module/dongle](https://github.com/golang-module/dongle "github.com/golang-module/dongle")

[gitee.com/golang-module/dongle](https://gitee.com/golang-module/dongle "gitee.com/golang-module/dongle")

### Installation
```go
// By github
go get -u github.com/golang-module/dongle

import (
    "github.com/golang-module/dongle"
)

// By gitee
go get -u gitee.com/golang-module/dongle

import (
    "gitee.com/golang-module/dongle"
)
```

### Usage and example

#### Encode and decode

##### Encode and decode by hex

```go
// Encode by hex from string and output string
dongle.Encode.FromString("hello world").ByHex().ToString() // 68656c6c6f20776f726c64
// Decode by hex from string and output string
dongle.Decode.FromString("68656c6c6f20776f726c64").ByHex().ToString() // hello world

// Encode by hex from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByHex().ToBytes() // []byte("68656c6c6f20776f726c64")
// Decode by hex from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("68656c6c6f20776f726c64")).ByHex().ToBytes() // []byte("hello world")
```

##### Encode and decode by base16

```go
// Encode by base16 from string and output string
dongle.Encode.FromString("hello world").ByBase16().ToString() // 68656c6c6f20776f726c64
// Decode by base16 from string and output string
dongle.Decode.FromString("68656c6c6f20776f726c64").ByBase16().ToString() // hello world

// Encode by base16 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase16().ToBytes() // []byte("68656c6c6f20776f726c64")
// Decode by base16 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("68656c6c6f20776f726c64")).ByBase16().ToBytes() // []byte("hello world")
```

##### Encode and decode by base32

```go
// Encode by base32 from string and output string
dongle.Encode.FromString("hello world").ByBase32().ToString() // NBSWY3DPEB3W64TMMQ======
// Decode by base32 from string and output string
dongle.Decode.FromString("NBSWY3DPEB3W64TMMQ======").ByBase32().ToString() // hello world

// Encode by base32 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase32().ToBytes() // []byte("NBSWY3DPEB3W64TMMQ======")
// Decode by base32 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("NBSWY3DPEB3W64TMMQ======")).ByBase32().ToBytes() // []byte("hello world")
```

##### Encode and decode by base58

```go
// Encode by base58 from string and output string
dongle.Encode.FromString("hello world").ByBase32().ToString() // StV1DL6CwTryKyV
// Decode by base58 from string and output string
dongle.Decode.FromString("StV1DL6CwTryKyV").ByBase32().ToString() // hello world

// Encode by base58 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase32().ToBytes() // []byte("StV1DL6CwTryKyV")
// Decode by base58 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("StV1DL6CwTryKyV")).ByBase32().ToBytes() // []byte("hello world")
```

##### Encode and decode by base64

```go
// Encode by base64 from string and output string
dongle.Encode.FromString("hello world").ByBase64().ToString() // aGVsbG8gd29ybGQ=
// Decode by base64 from string and output string
dongle.Decode.FromString("aGVsbG8gd29ybGQ=").ByBase64().ToString() // hello world

// Encode by base64 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase64().ToBytes() // []byte("aGVsbG8gd29ybGQ=")
// Decode by base64 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("aGVsbG8gd29ybGQ=")).ByBase64().ToBytes() // []byte("hello world")

```

##### Encode and decode by base64URL
```go
// Encode by base64 from url string and output string
dongle.Encode.FromString("www.gouguoyin.cn").ByBase64URL().ToString() // d3d3LmdvdWd1b3lpbi5jbg==
// Decode by base64 from string and output url string
dongle.Decode.FromString("d3d3LmdvdWd1b3lpbi5jbg==").ByBase64URL().ToString() // www.gouguoyin.cn

// Encode by base64 from url byte slice and output byte slice
dongle.Encode.FromBytes([]byte("www.gouguoyin.cn")).ByBase64URL().ToBytes() // []byte("d3d3LmdvdWd1b3lpbi5jbg==")
// Decode by base64 from byte slice and output url byte slice
dongle.Decode.FromBytes([]byte("d3d3LmdvdWd1b3lpbi5jbg==")).ByBase64URL().ToBytes() // []byte("www.gouguoyin.cn")
```

##### Encode and decode by base85

```go
// Encode by base85 from string and output string
dongle.Encode.FromString("hello world").ByBase64().ToString() // BOu!rD]j7BEbo7
// Decode by base85 from string and output string
dongle.Decode.FromString("BOu!rD]j7BEbo7").ByBase64().ToString() // hello world

// Encode by base85 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase64().ToBytes() // []byte("BOu!rD]j7BEbo7")
// Decode by base85 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("BOu!rD]j7BEbo7")).ByBase64().ToBytes() // []byte("hello world")
```

##### Encode and decode by safeURL

```go
// Encode by escape from url string and output string
dongle.Encode.FromString("www.gouguoyin.cn?sex=男&age=18").BySafeURL().ToString() // www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18
// Decode by escape from string and output url string
dongle.Decode.FromString("www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18").BySafeURL().ToString() // www.gouguoyin.cn?sex=男&age=18

// Encode by escape from url byte slice and output byte slice
dongle.Encode.FromBytes([]byte("www.gouguoyin.cn?sex=男&age=18")).BySafeURL().ToBytes() // []byte("www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18")
// Decode by escape from byte slice and output url byte slice
dongle.Decode.FromBytes([]byte("www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18")).BySafeURL().ToBytes() // []byte("www.gouguoyin.cn?sex=男&age=18")
```

#### Encrypt and decrypt

##### Encrypt by md4

```go
// Encrypt by md4 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByMd4().ToHexString() // aa010fbc1d14c795d86ef98c95479d17
// Encrypt by md4 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByMd4().ToBase64String() // qgEPvB0Ux5XYbvmMlUedFw==

// Encrypt by md4 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd4().ToHexBytes() // []byte("aa010fbc1d14c795d86ef98c95479d17")
// Encrypt by md4 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd4().ToBase64Bytes() // []byte("qgEPvB0Ux5XYbvmMlUedFw==")

// Encrypt by md4 from file and output string with hex encoding
dongle.Encrypt.FromFile("./LICENSE").ByMd4().ToHexString() // 1240c5c0fb26b585999357915c56b511
// Encrypt by md4 from file and output string with base64 encoding
dongle.Encrypt.FromFile("./LICENSE").ByMd4().ToBase64String() // EkDFwPsmtYWZk1eRXFa1EQ==

// Encrypt by md4 from file and output byte slice with hex encoding
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd4().ToHexBytes() // []byte("1240c5c0fb26b585999357915c56b511")
// Encrypt by md4 from file and output byte slice with base64 encoding
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd4().ToBase64Bytes() // []byte("EkDFwPsmtYWZk1eRXFa1EQ==")
```

##### Encrypt by hmac-md4
```go
// Encrypt by hmac-md4 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacMd4("dongle").ToHexString() // 7a9df5247cbf76a8bc17c9c4f5a75b6b
// Encrypt by hmac-md4 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacMd4("dongle").ToBase64String() // ep31JHy/dqi8F8nE9adbaw==

// Encrypt by hmac-md4 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToHexBytes() // []byte("7a9df5247cbf76a8bc17c9c4f5a75b6b")
// Encrypt by hmac-md4 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToBase64Bytes() // []byte("ep31JHy/dqi8F8nE9adbaw==")
```

##### Encrypt by md5
```go
// Encrypt by md5 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByMd5().ToHexString() // 5eb63bbbe01eeed093cb22bb8f5acdc3
// Encrypt by md5 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByMd5().ToBase64String() // XrY7u+Ae7tCTyyK7j1rNww==

// Encrypt by md5 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd5().ToHexBytes() // []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")
// Encrypt by md5 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd5().ToBase64Bytes() // []byte("XrY7u+Ae7tCTyyK7j1rNww==")

// Encrypt by md5 from file and output string with hex encoding
dongle.Encrypt.FromFile("./LICENSE").ByMd5().ToHexString() // 014f03f9025ea81a8a0e9734be540c53
// Encrypt by md5 from file and output string with base64 encoding
dongle.Encrypt.FromFile("./LICENSE").ByMd5().ToBase64String() // AU8D+QJeqBqKDpc0vlQMUw==

// Encrypt by md5 from file and output byte slice with hex encoding
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd5().ToHexBytes() // []byte("014f03f9025ea81a8a0e9734be540c53")
// Encrypt by md5 from file and output byte slice with base64 encoding
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd5().ToBase64Bytes() // []byte("AU8D+QJeqBqKDpc0vlQMUw==")
```

##### Encrypt by hmac-md5
```go
// Encrypt by hmac-md5 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacMd5("dongle").ToHexString() // 4790626a275f776956386e5a3ea7b726
// Encrypt by hmac-md5 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacMd5("dongle").ToBase64String() // R5Biaidfd2lWOG5aPqe3Jg==

// Encrypt by hmac-md5 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToHexBytes() // []byte("4790626a275f776956386e5a3ea7b726")
// Encrypt by hmac-md5 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToBase64Bytes() // []byte("R5Biaidfd2lWOG5aPqe3Jg==")
```

##### Encrypt by sha1
```go
// Encrypt by sha1 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha1().ToHexString() // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
// Encrypt by sha1 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha1().ToBase64String() // Kq5sNclPz7QV2+lfQIuc6R7oRu0=

// Encrypt by sha1 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha1().ToHexBytes() // []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
// Encrypt by sha1 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha1().ToBase64Bytes() // []byte("Kq5sNclPz7QV2+lfQIuc6R7oRu0=")
```

##### Encrypt by hmac-sha1
```go
// Encrypt by hmac-sha1 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha1("dongle").ToHexString() // 91c103ef93ba7420902b0d1bf0903251c94b4a62
// Encrypt by hmac-sha1 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha1("dongle").ToBase64String() // kcED75O6dCCQKw0b8JAyUclLSmI=

// Encrypt by hmac-sha1 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToHexBytes() // []byte("91c103ef93ba7420902b0d1bf0903251c94b4a62")
// Encrypt by hmac-sha1 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToBase64Bytes() // []byte("kcED75O6dCCQKw0b8JAyUclLSmI=")
```

##### Encrypt by sha224

```go
// Encrypt by sha224 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha224().ToHexString() // 2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b
// Encrypt by sha224 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha224().ToBase64String() // LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==

// Encrypt by sha224 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha224().ToHexBytes() // []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
// Encrypt by sha224 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha224().ToBase64Bytes() // []byte("LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==")
```

##### Encrypt by hmac-sha224
```go
// Encrypt by hmac-sha224 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha224("dongle").ToHexString() // e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec
// Encrypt by hmac-sha224 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha224("dongle").ToBase64String() // 4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A====

// Encrypt by hmac-sha224 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToHexBytes() // []byte("e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec")
// Encrypt by hmac-sha224 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToBase64Bytes() // []byte("4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A==")
```

##### Encrypt by sha256

```go
// Encrypt by sha256 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha256().ToHexString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// Encrypt by sha256 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha256().ToBase64String() // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// Encrypt by sha256 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha256().ToHexBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// Encrypt by sha256 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha256().ToBase64Bytes() // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")
```

##### Encrypt by hmac-sha256
```go
// Encrypt by hmac-sha256 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha256("dongle").ToHexString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// Encrypt by hmac-sha256 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha256("dongle").ToBase64String() // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// Encrypt by hmac-sha256 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToHexBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// Encrypt by hmac-sha256 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToBase64Bytes() // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")
```

##### Encrypt by sha384

```go
// Encrypt by sha384 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha384().ToHexString() // fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
// Encrypt by sha384 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha384().ToBase64String() // /b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9

// Encrypt by sha384 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha384().ToHexBytes() // []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")
// Encrypt by sha384 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha384().ToBase64Bytes() // []byte("/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9")
```

##### Encrypt by hmac-sha384
```go
// Encrypt by hmac-sha384 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha384("dongle").ToHexString() // 421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8
// Encrypt by hmac-sha384 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha384("dongle").ToBase64String() // Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o

// Encrypt by hmac-sha384 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToHexBytes() // []byte("421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8")
// Encrypt by hmac-sha384 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToBase64Bytes() // []byte("Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o")
```

##### Encrypt by sha512
```go
// Encrypt by sha512 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha512().ToHexBytes() // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
// Encrypt by sha512 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha512().ToBase64String() // MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==

// Encrypt by sha512 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToHexBytes() // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
// Encrypt by sha512 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToBase64Bytes() // []byte("MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==")
```

##### Encrypt by hmac-sha512
```go
// Encrypt by hmac-sha512 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha512("dongle").ToHexString() // d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1
// Encrypt by hmac-sha512 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha512("dongle").ToBase64String() // 2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q==

// Encrypt by hmac-sha512 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToHexBytes() // []byte("d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1")
// Encrypt by hmac-sha512 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToBase64Bytes() // []byte("2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q==")
```

##### Encrypt by rc4
```go
// Encrypt by rc4 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByRc4("dongle").ToHexString() // eba154b4cb5a9038dbbf9d
// Encrypt by rc4 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByRc4("dongle").ToBase64String() // 66FUtMtakDjbv50=

// Encrypt by rc4 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4("dongle").ToHexBytes() // []byte("eba154b4cb5a9038dbbf9d")
// Encrypt by rc4 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4("dongle").ToBase64Bytes() // []byte("66FUtMtakDjbv50=")
```

##### Encrypt and decrypt by rsa

```go
pkcs1PublicKey := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCtdjE3fOKpAmc6eIi1I/GElIJW
hTLBZb/SpSC+Pl7bONDyt/sG7FzxMHLoxnlJnYDHmYiu0iy0s/EHhNL+1bPnBiuA
bOKi7TBgojusBLsrddCrGn08gZLRZ3ZP/oAS6+wf2/vfh2jCCKCVHbBZLHQIN3MU
0qcFQKXvfJaBTUX6gwIDAQAB
-----END PUBLIC KEY-----`

pkcs8PublicKey := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCqzZNa9VrcewyU6wDoV7Y9kAHq
X1VK0B3Rb6GNmQe4zLEfce7cVTaLrc4VGTKl35tADG1cRHqtaG4S/WttpiGZBhxJ
y4MpOXb6eIPiVLsn2lL+rJo5XdbSr3gyjxEOQQ97ihtw4lDd5wMo4bIOuw1LtMez
HC1outlM6x+/BB0BSQIDAQAB
-----END PUBLIC KEY-----`

pkcs1PrivateKey := `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCtdjE3fOKpAmc6eIi1I/GElIJWhTLBZb/SpSC+Pl7bONDyt/sG
7FzxMHLoxnlJnYDHmYiu0iy0s/EHhNL+1bPnBiuAbOKi7TBgojusBLsrddCrGn08
gZLRZ3ZP/oAS6+wf2/vfh2jCCKCVHbBZLHQIN3MU0qcFQKXvfJaBTUX6gwIDAQAB
AoGAFwAfEo56t5JcAcLNzccQVVYj2jkbO820G8hNiSxYA5WLD0QaAxcAU/Lqqbb3
ii1aUB0ppJS13NgnU6nnGGdZzUYBG1Hai6EkVyCGrI4amQ93AaVdKncL8gJ4RZAm
YzPPUwSMEESsu24pS1NF1G1Y8C+28b/Wr0oqOsCvL6PhsMECQQDwsPJJoWRx7ZJw
E1K5KLT0cXKyrIpyXY3I6tyA5imCzOzccf3d1vDgB0L9sdSO7bG3ceSwpAeiWEbg
5jGZemPzAkEAuH6U4pEI4AMbWnatpK55Rc235NDgmT3VyIuRaKC02YXAZ+jznFep
XMd4DTli4R9r3j2YVhUpyDVbdQpFH98DMQJAQpOvcU6DSkA80WOG7lCkPTlkUKgJ
Y7kdDwZoF/+SW+vzWMbvQf3CgzV/Ak2+TgrRrbyDVZkJw45HjM4fyiRgoQJBALH/
/qlxgPyQQs3O/s2KQBsm1auAE5IF5MLuVUZ69sF/mBko2hEXSqHnGV645TuKU0pC
Zz12ga9WO3z6gaK0SaECQQDah1pKt9ViBBy4USXK3OWXEloHuTwmyr9AbLqqI5tQ
2eNuH0NkuJYQmnXmHLbKOELoYocldEBXmkzPXSN+X9kV
-----END RSA PRIVATE KEY-----`

pkcs8PrivateKey := `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAKrNk1r1Wtx7DJTr
AOhXtj2QAepfVUrQHdFvoY2ZB7jMsR9x7txVNoutzhUZMqXfm0AMbVxEeq1obhL9
a22mIZkGHEnLgyk5dvp4g+JUuyfaUv6smjld1tKveDKPEQ5BD3uKG3DiUN3nAyjh
sg67DUu0x7McLWi62UzrH78EHQFJAgMBAAECgYAeo3nHWzPNURVUsUMcan96U5bE
YA2AugxfQVMNf2HvOGidZ2adh3udWrQY/MglERNcTd5gKriG2rDEH0liBecIrNKs
BL4lV+qHEGRUcnDDdtUBdGInEU8lve5keDgmX+/huXSRJ+3tYA5u9j+32RquVczv
Idtb5XnBLUl61k0osQJBAON5+eJjtw6xpn+pveU92BSHvaJYVyrLHwUjR07aNKb7
GlGVM3MGf1FCa8WQUo9uUzYxGLtg5Qf3sqwOrwPd5UsCQQDAOF/zWqGuY3HfV/1w
giXiWp8rc+S8tanMj5M37QQbYW5YLjUmJImoklVahv3qlgLZdEN5ZSueM5jfoSFt
Nts7AkBKoRDvSiGbi4MBbTHkzLZgfewkH/FxE7S4nctePk553fXTgCyh9ya8BRuQ
dHnxnpNkOxVPHEnnpEcVFbgrf5gjAkB7KmRI4VTiEfRgINhTJAG0VU7SH/N7+4cu
fPzfA+7ywG5c8Fa79wOB0SoB1KeUjcSLo5Ssj2fwea1F9dAeU90LAkBJQFofveaD
a3YlN4EQZOcCvJKmg7xwWuGxFVTZDVVEws7UCQbEOEEXZrNd9x0IF5kpPLR+rxua
RPgUNaDGIh5o
-----END PRIVATE KEY-----`

// Use public key to encrypt by rsa from string
cipherText := dongle.Encrypt.FromString("hello world").ByRsa(pkcs1PublicKey)
// Use pkcs1 private key to decrypt by rsa from string with hex encoding
dongle.Decrypt.FromHexString(cipherText.ToHexString()).ByRsa(pkcs1PrivateKey, dongle.PKCS1).ToString() // hello world
// Use pkcs1 private key to decrypt by rsa from string with base64 encoding
dongle.Decrypt.FromBase64String(cipherText.ToBase64String()).ByRsa(pkcs1PrivateKey, dongle.PKCS1).ToString() // hello world

// Use public key to encrypt by rsa from byte
cipherText := dongle.Encrypt.FromBytes([]byte("hello world")).ByRsa([]byte(pkcs1PublicKey))
// Use pkcs1 private key to decrypt by rsa from byte with hex encoding
dongle.Decrypt.FromHexBytes(cipherText.ToHexBytes()).ByRsa([]byte(pkcs1PrivateKey), dongle.PKCS1).ToByte() // []bytes("hello world)
// Use pkcs1 private key to decrypt by rsa from byte with base64 encoding
dongle.Decrypt.FromBase64Bytes(cipherText.ToBase64Bytes()).ByRsa([]byte(pkcs1PrivateKey), dongle.PKCS1).ToByte() // []bytes("hello world)

// Use public key to encrypt by rsa from string
cipherText := dongle.Encrypt.FromString("hello world").ByRsa(pkcs8PublicKey)
// Use pkcs8 private key to decrypt by rsa from string with hex encoding
dongle.Decrypt.FromHexString(cipherText.ToHexString()).ByRsa(pkcs8PrivateKey, dongle.PKCS8).ToString() // hello world
// Use pkcs8 private key to decrypt by rsa from string with base64 encoding
dongle.Decrypt.FromBase64String(cipherText.ToBase64String()).ByRsa(pkcs8PrivateKey, dongle.PKCS8).ToString() // hello world

// Use public key to encrypt by rsa from byte
cipherText := dongle.Encrypt.FromBytes([]byte("hello world")).ByRsa([]byte(pkcs8PublicKey))
// Use pkcs8 private key to decrypt by rsa from byte with hex encoding
dongle.Decrypt.FromHexBytes(cipherText).ByRsa(pkcs8PrivateKey, dongle.PKCS8).ToByte() // []bytes("hello world)
// Use pkcs8 private key to decrypt by rsa from byte with base64 encoding
dongle.Decrypt.FromBase64Bytes(cipherText.ToBase64Bytes()).ByRsa(pkcs8PrivateKey, dongle.PKCS8).ToByte() // []bytes("hello world)
```

##### Encrypt and decrypt by aes

```go
cipher := NewCipher()
cipher.SetMode(dongle.CBC) // CBC、CFB、OFB、CTR、ECB、GCM
cipher.SetPadding(dongle.PKCS7) // No、Zero、PKCS5、PKCS7
cipher.SetKey("0123456789abcdef") // key must be 16, 24 or 32 bytes
cipher.SetIV("0123456789abcdef")  // iv must be 16 bytes (ECB mode doesn't require setting iv)

// Encrypt by aes from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToHexString() // c1e9b4529aac9793010f4677f6358efe
// Decrypt by aes from string with hex encoding and output string
dongle.Decrypt.FromHexString("c1e9b4529aac9793010f4677f6358efe").ByAes(cipher).ToString() // hello world

// Encrypt by aes from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToBase64String() // wem0Upqsl5MBD0Z39jWO/g==
// Decrypt by aes from string with base64 encoding and output string
dongle.Decrypt.FromBase64String("wem0Upqsl5MBD0Z39jWO/g==").ByAes(cipher).ToString() // hello world

// Encrypt by aes from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToHexBytes() // []byte("c1e9b4529aac9793010f4677f6358efe")
// Decrypt by aes from byte slice with hex encoding and output byte slice
dongle.Decrypt.FromHexBytes([]byte("c1e9b4529aac9793010f4677f6358efe")).ByAes(cipher).ToBytes() // []byte("hello world")

// Encrypt by aes from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToBase64Bytes() // []byte("wem0Upqsl5MBD0Z39jWO/g==")
// Decrypt by aes from byte slice with base64 encoding and output byte slice
dongle.Decrypt.FromBase64Bytes(()byte("wem0Upqsl5MBD0Z39jWO/g==")).ByAes(cipher).ToBytes() // []byte("hello world")

```

##### Encrypt and decrypt by des

```go
cipher := NewCipher()
cipher.SetMode(dongle.CBC) // CBC、ECB、CFB、OFB、CTR、GCM
cipher.SetPadding(dongle.PKCS7) // No、Zero、PKCS5、PKCS7
cipher.SetKey("12345678")       // key must be 8 bytes
cipher.SetIV("123456788") // iv must be 8 bytes

// Encrypt by des from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByDes(cipher).ToHexString() // 0b2a92e81fb49ce1a43266aacaea7b81
// Decrypt by des from string with hex encoding and output string
dongle.Decrypt.FromHexString("0b2a92e81fb49ce1a43266aacaea7b81").ByDes(cipher).ToString() // hello world

// Encrypt by des from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByDes(cipher).ToBase64String() // CyqS6B+0nOGkMmaqyup7gQ==
// Decrypt by des from string with base64 encoding and output string
dongle.Decrypt.FromBase64String("CyqS6B+0nOGkMmaqyup7gQ==").ByDes(cipher).ToString() // hello world

// Encrypt by des from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByDes(cipher).ToHexBytes() // []byte("0b2a92e81fb49ce1a43266aacaea7b81")
// Decrypt by des from byte slice with hex encoding and output byte slice
dongle.Decrypt.FromHexBytes([]byte("0b2a92e81fb49ce1a43266aacaea7b81")).ByDes(cipher).ToBytes() // []byte("hello world")

// Encrypt by des from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByDes(cipher).ToBase64Bytes() // []byte("CyqS6B+0nOGkMmaqyup7gQ==")
// Decrypt by des from byte slice with base64 encoding and output byte slice
dongle.Decrypt.FromBase64Bytes(()byte("CyqS6B+0nOGkMmaqyup7gQ==")).ByDes(cipher).ToBytes() // []byte("hello world")
```

##### Encrypt and decrypt by 3des

```go
cipher := NewCipher()
cipher.SetMode(dongle.CBC) // CBC、ECB、CFB、OFB、CTR、GCM
cipher.SetPadding(dongle.PKCS7) // No、Zero、PKCS5、PKCS7
cipher.SetKey("123456781234567812345678") // key must be 24 bytes
cipher.SetIV("123456788") // iv must be 8 bytes

// Encrypt by 3des from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").By3Des(cipher).ToHexString() // 0b2a92e81fb49ce1a43266aacaea7b81
// Decrypt by des from string with hex encoding and output string
dongle.Decrypt.FromHexString("0b2a92e81fb49ce1a43266aacaea7b81").By3Des(cipher).ToString() // hello world

// Encrypt by 3des from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").By3Des(cipher).ToBase64String() // CyqS6B+0nOGkMmaqyup7gQ==
// Decrypt by 3des from string with base64 encoding and output string
dongle.Decrypt.FromBase64String("CyqS6B+0nOGkMmaqyup7gQ==").By3Des(cipher).ToString() // hello world

// Encrypt by 3des from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).By3Des(cipher).ToHexBytes() // []byte("0b2a92e81fb49ce1a43266aacaea7b81")
// Decrypt by 3des from byte slice with hex encoding and output byte slice
dongle.Decrypt.FromHexBytes([]byte("0b2a92e81fb49ce1a43266aacaea7b81")).By3Des(cipher).ToBytes() // []byte("hello world")

// Encrypt by 3des from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).By3Des(cipher).ToBase64Bytes() // []byte("CyqS6B+0nOGkMmaqyup7gQ==")
// Decrypt by 3des from byte slice with base64 encoding and output byte slice
dongle.Decrypt.FromBase64Bytes(()byte("CyqS6B+0nOGkMmaqyup7gQ==")).By3Des(cipher).ToBytes() // []byte("hello world")
```

### Error handling

> If more than one error occurs, only the first error is returned

```go
e := dongle.Encrypy.FromString("hello world").ByRsa("xxxx")
if e.Error != nil {
// 错误处理...
log.Fatal(e.Error)
}
fmt.Println(e.ToString())
// Output
invalid public key, please make sure the public key is valid
```

### Feature list

- [x] Encoding and decoding by Hex
- [x] Encoding and decoding by Base16
- [x] Encoding and decoding by Base32
- [x] Encoding and decoding by Base58
- [ ] Encoding and decoding by Base62
- [x] Encoding and decoding by Base64
- [x] Encoding and decoding by Base64URL
- [x] Encoding and decoding by SafeURL
- [x] Encoding and decoding by Base85
- [ ] Encoding and decoding by Base91
- [ ] Encoding and decoding by Base92
- [ ] Encoding and decoding by Base100
- [x] Encryption by Md4
- [x] Encryption by Hmac-md4
- [x] Encryption by Md5
- [x] Encryption by Hmac-md5
- [x] Encryption by Sha1
- [x] Encryption by Hmac-sha1
- [x] Encryption by Sha224
- [x] Encryption by Hmac-sha224
- [x] Encryption by Sha256
- [x] Encryption by Hmac-sha256
- [x] Encryption by Sha384
- [x] Encryption by Hmac-sha384
- [x] Encryption by Sha512
- [x] Encryption by Hmac-sha512
- [ ] Encryption by Rc2
- [x] Encryption by Rc4
- [ ] Encryption by Rc5
- [ ] Encryption by Rc6
- [ ] Encryption and decryption by Sm2
- [ ] Encryption and decryption by Sm3
- [ ] Encryption and decryption by Sm4
- [ ] Encryption and decryption by Sm7
- [ ] Encryption and decryption by Sm9
- [ ] Encryption and decryption by DSA
- [ ] Encryption and decryption by Tea
- [ ] Encryption and decryption by Xtea
- [x] Encryption and decryption by RSA with PKCS1Pem/PKCS8Pem
- [x] Encryption and decryption by AES with ECB/CBC/CTR/CFB/OFB mode and No/Zero/PKCS5/PKCS7 padding
- [x] Encryption and decryption by DES with ECB/CBC/CTR/CFB/OFB mode and No/Zero/PKCS5/PKCS7 padding
- [x] Encryption and decryption by 3DES with ECB/CBC/CTR/CFB/OFB mode and No/Zero/PKCS5/PKCS7 padding

### References

* [javascript/crypto-js](https://github.com/brix/crypto-js)
* [nodejs/crypto](https://nodejs.org/api/crypto.html)
* [java/jasypt](https://github.com/jasypt/jasypt)
* [python/pycryptodome](https://github.com/Legrandin/pycryptodome)

### Online website

* [www.ssleye.com](https://www.ssleye.com/ssltool)
* [www.sojson.com](https://www.sojson.com/encrypt.html)
* [tool.chacuo.net](https://tool.chacuo.net/cryptaes)
* [www.oktools.net](https://oktools.net/aes)

### Sponsors

`Dongle` is a non-commercial open source project. If you want to support `Dongle`, you
can [buy a cup of coffee](https://www.gouguoyin.cn/zanzhu.html) for developer.

### Thanks

`Dongle` had been being developed with GoLand under the free JetBrains Open Source license, I would like to express my
thanks here.

<a href="https://www.jetbrains.com"><img src="https://github-oss.oss-cn-beijing.aliyuncs.com/jetbrains.png" height="100" alt="JetBrains"/></a>