package main

import "fmt"

// SDS 接口
type SDS interface {
	SetString(str string)
	String() string
	AppendString(str string)
}

// sdshdrRaw 结构体
type sdshdrRaw struct {
	len  uint32 // 当前字符串的长度
	free uint32 // 剩余的空闲空间
	data []byte // 存储字符串数据的切片
}

// sdshdrEmbstr 结构体
type sdshdrEmbstr struct {
	len  uint8  // 当前字符串的长度
	free uint8  // 剩余的空闲空间
	data []byte // 存储字符串数据的切片
}

// NewSDS 创建一个新的 SDS
func NewSDS(str string) SDS {
	strLen := len(str)
	if strLen <= 44 { // 假设 embstr 编码的最大长度为 44 字节
		return &sdshdrEmbstr{
			len:  uint8(strLen),
			free: 0,
			data: []byte(str),
		}
	} else {
		return &sdshdrRaw{
			len:  uint32(strLen),
			free: 0,
			data: []byte(str),
		}
	}
}

// SetString 设置 SDS 中的字符串
func (s *sdshdrRaw) SetString(str string) {
	strLen := len(str)
	s.data = make([]byte, strLen)
	copy(s.data, str)
	s.len = uint32(strLen)
	s.free = 0
}

func (s *sdshdrEmbstr) SetString(str string) {
	strLen := len(str)
	s.data = make([]byte, strLen)
	copy(s.data, str)
	s.len = uint8(strLen)
	s.free = 0
}

// String 获取 SDS 中的字符串
func (s *sdshdrRaw) String() string {
	return string(s.data[:s.len])
}

func (s *sdshdrEmbstr) String() string {
	return string(s.data[:s.len])
}

// AppendString 追加字符串
func (s *sdshdrRaw) AppendString(str string) {
	strLen := len(str)
	requiredLen := int(s.len) + strLen
	if requiredLen > len(s.data) {
		s.data = append(s.data, make([]byte, requiredLen-len(s.data))...)
	}
	copy(s.data[s.len:], str)
	s.len = uint32(requiredLen)
	s.free = uint32(len(s.data)) - s.len
}

func (s *sdshdrEmbstr) AppendString(str string) {
	strLen := len(str)
	requiredLen := int(s.len) + strLen
	if requiredLen > len(s.data) {
		s.data = append(s.data, make([]byte, requiredLen-len(s.data))...)
	}
	copy(s.data[s.len:], str)
	s.len = uint8(requiredLen)
	s.free = uint8(len(s.data)) - s.len
}

// Allocate1MB 分配 1MB 内存
func (s *sdshdrRaw) Allocate1MB() {
	s.data = append(s.data, make([]byte, 1024*1024)...)
	s.free = uint32(len(s.data)) - s.len
}

// Allocate1MB 分配 1MB 内存
func (s *sdshdrEmbstr) Allocate1MB() {
	s.data = append(s.data, make([]byte, 1024*1024)...)
	s.free = uint8(len(s.data)) - s.len
}

func main() {
	// 创建一个新的 SDS
	sds := NewSDS("Hello, SDS!")

	// 追加字符串
	sds.AppendString(" This is a longer string.")
	fmt.Println(sds.String()) // 输出: Hello, SDS! This is a longer string.

	// 分配 1MB 内存
	if rawSDS, ok := sds.(*sdshdrRaw); ok {
		rawSDS.Allocate1MB()
	} else if embstrSDS, ok := sds.(*sdshdrEmbstr); ok {
		embstrSDS.Allocate1MB()
	}

	// 追加更长的字符串
	sds.AppendString("This is another longer string.")
	fmt.Println(sds.String()) // 输出: Hello, SDS! This is a longer string.This is another longer string.
}
