package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	var (
		builder   strings.Builder
		reader    *strings.Reader
		replacer  *strings.Replacer
		readBytes []byte
		text      string
	)

	//判断两个字符串的大小写敏感的字节是否相等
	fmt.Println(strings.EqualFold("abc", "ABC"))
	//判断字符串是否以指定的字符串开头
	fmt.Println(strings.HasPrefix("abc", "a"))
	//判断字符串是否以指定的字符串结尾
	fmt.Println(strings.HasSuffix("abc", "c"))
	//判断字符串是否包含指定的字符串
	fmt.Println(strings.Contains("abc", "b"))
	//判断字符串是否包含指定的字符
	fmt.Println(strings.ContainsRune("abc", 'b'))
	//判断字符串是否包含指定的任意字符
	fmt.Println(strings.ContainsAny("abc", "bcd"))
	//统计字符串中指定的字符串出现的次数
	fmt.Println(strings.Count("abcabc", "bc"))
	//查找字符串中指定的字符串第一次出现的位置
	fmt.Println(strings.Index("abcabc", "bc"))
	//查找字符串中指定的字符第一次出现的位置
	fmt.Println(strings.IndexByte("abcabc", 'b'))
	//查找字符串中指定的字符第一次出现的位置
	fmt.Println(strings.IndexRune("abcabc", 'b'))
	//查找字符串中指定的任意字符第一次出现的位置
	fmt.Println(strings.IndexAny("abcabc", "cbd"))
	//查找字符串满足指定函数的第一个字符的位置
	fmt.Println(strings.IndexFunc("abc", func(r rune) bool {
		return r == 'c'
	}))
	//查找字符串中指定的字符串最后一次出现的位置
	fmt.Println(strings.LastIndex("abcabc", "bc"))
	//查找字符串中指定的字符最后一次出现的位置
	fmt.Println(strings.LastIndexByte("abcabc", 'b'))
	//查找字符串中指定的任意字符最后一次出现的位置
	fmt.Println(strings.LastIndexAny("abcabc", "cfs"))
	//查找字符串满足指定函数的最后一个字符的位置
	fmt.Println(strings.LastIndexFunc("abc", func(r rune) bool {
		return r == 'c'
	}))
	//大写转小写
	fmt.Println(strings.ToLower("ABC"))
	fmt.Println(strings.ToLowerSpecial(unicode.SpecialCase{}, "ABC"))
	//小写转大写
	fmt.Println(strings.ToUpper("abc"))
	fmt.Println(strings.ToUpperSpecial(unicode.SpecialCase{}, "abc"))
	//
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitleSpecial(unicode.SpecialCase{}, "loud noises"))
	//将字符串重复指定的次数
	fmt.Println(strings.Repeat("ABC", 5))
	//将字符串中的指定字符替换为指定的字符串
	fmt.Println(strings.Replace("ABCABCABC", "A", "B", -1))
	fmt.Println(strings.ReplaceAll("ABCABCABC", "A", "B"))
	//将字符串的每个字符替换为指定的函数返回的字符
	fmt.Println(strings.Map(func(r rune) rune {
		if r == 'A' {
			return 'B'
		} else if r == 'B' {
			return 'A'
		}
		return r
	}, "ABCABCABC"))
	//将字符串前后指定的字符去掉
	fmt.Println(strings.Trim(" ABC ", " "))
	//将字符串前后的空白字符去掉
	fmt.Println(strings.TrimSpace(" ABC "))
	//将字符串前的指定的字符去掉
	fmt.Println(strings.TrimLeft(" ABC ", " "))
	//将字符串后的指定的字符去掉
	fmt.Println(strings.TrimRight(" ABC ", " "))
	//将字符串指定的前缀去掉
	fmt.Println(strings.TrimPrefix("ABCABCABC", "ABC"))
	//将字符串指定的后缀去掉
	fmt.Println(strings.TrimSuffix("ABCABCABC", "ABC"))
	//将字符串前后满足指定函数的字符去掉
	fmt.Println(strings.TrimFunc("abcABCabc", func(r rune) bool {
		return r > 'Z'
	}))
	//将字符串前满足指定函数的字符去掉
	fmt.Println(strings.TrimLeftFunc("abcABCabc", func(r rune) bool {
		return r > 'Z'
	}))
	//将字符串后满足指定函数的字符去掉
	fmt.Println(strings.TrimRightFunc("abcABCabc", func(r rune) bool {
		return r > 'Z'
	}))
	//将字符串按unicode.IsSpace空白字符分割
	fmt.Println(strings.Fields("I am jfxy"))
	//将字符串按指定函数返回的字符分割
	fmt.Println(strings.FieldsFunc("I am jfxy", unicode.IsSpace))
	//将字符串按指定的字符分割
	fmt.Println(strings.Split("I_am_jfxy", "_"))
	//将字符串按指定的字符分割，返回切片中最多包含n个元素
	fmt.Println(strings.SplitN("I_am_jfxy", "_", 2))
	//将字符串按指定的字符分割，分割的字符串中包含分割字符
	fmt.Println(strings.SplitAfter("I__am_jfxy", "_"))
	//将字符串按指定的字符分割，分割的字符串中包含分割字符，返回切片中最多包含n个元素
	fmt.Println(strings.SplitAfterN("I__am_jfxy", "_", 2))
	//将一系列字符串连接为一个字符串，之间用指定的字符分割
	fmt.Println(strings.Join([]string{"I", "am", "jfxy"}, "_"))

	//strings.Builder
	builder = strings.Builder{}
	//向builder中写入[]byte
	builder.Write([]byte("hello world"))
	//向builder中写入byte
	builder.WriteByte('!')
	//向builder中写入string
	builder.WriteString("hello world")
	//向builder中写入rune
	builder.WriteRune('！')
	//将内部的[]byte转换为string
	text = builder.String()
	//清空builder中的内容
	builder.Reset()
	fmt.Println(text, builder.String())

	//strings.Reader
	reader = strings.NewReader("你好 世界！")
	//读取一个byte
	fmt.Println(reader.ReadByte())
	//ReadByte的回退方法
	fmt.Println(reader.UnreadByte())
	//读取一个rune
	fmt.Println(reader.ReadRune())
	//ReadRune的回退方法
	fmt.Println(reader.UnreadRune())
	//读取len([]byte)长度的字节
	readBytes = make([]byte, 7)
	reader.Read(readBytes)
	fmt.Println(len(readBytes), string(readBytes))
	//从指定的位置读取len([]byte)长度的字节
	fmt.Println(reader.ReadAt(readBytes, 3))
	fmt.Println(len(readBytes), string(readBytes))
	//返回未读取的字节长度和总字节长度
	fmt.Println(reader.Len(), reader.Size())
	//修改读取位置
	fmt.Println(reader.Seek(0, io.SeekStart))
	//重置reader
	reader.Reset("hello world!")
	//将reader的内容输出
	reader.WriteTo(os.Stdout)

	//strings.Replacer
	replacer = strings.NewReplacer("A", "B", "B", "A")
	//使用strings.Replacer的规则去替换字符串
	fmt.Println(replacer.Replace("ABCABCABC"))
	//使用strings.Replacer的规则去替换字符串并输出
	replacer.WriteString(os.Stdout, "ABCABCABC")
}
