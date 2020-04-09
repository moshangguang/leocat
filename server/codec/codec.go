package codec

import (
	"io"
)

// Encoder 编码器
type Encoder interface {
	Encode(obj interface{}) error
}

// Decoder 解码器
type Decoder interface {
	Decode(obj interface{}) error
}

// Codec 编解码集规范接口，可被替换为自实现编解码集
type Codec interface {
	// NewEncoder 新建解码器，需要池化
	NewEncoder(writer io.Writer) Encoder
	// NewDecoder 新建编码器，需要池化
	NewDecoder(reader io.Reader) Decoder
}
