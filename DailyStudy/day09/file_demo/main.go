package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作

// os buferio ioutil

func osFileRead() {
	// os
	fileObj, err := os.Open("./text.txt")
	if err != nil {
		fmt.Printf("err:%s", err)
		return
	}
	fmt.Println(fileObj)
	var b = make([]byte, 128)
	num, err := fileObj.Read(b)

	if err == io.EOF {
		fmt.Println("文件读取完毕")
	}

	if err != nil {
		fmt.Printf("err:%s", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println(num)

	defer fileObj.Close()
}

//OsReadAll os.open file read
func OsReadAll() {
	// os
	fileObj, err := os.Open("./text.txt")
	if err != nil {
		fmt.Printf("err:%s", err)
		return
	}

	defer fileObj.Close()

	var content []string
	var b = make([]byte, 128)
	for {
		_, err := fileObj.Read(b)

		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}

		if err != nil {
			fmt.Printf("err:%s", err)
			return
		}
		content = append(content, string(b))
		//fmt.Println(string(b))
		//fmt.Println(num)
	}

	fmt.Println(content)

}

// buferio

func bufioRead() {
	fileObj, err := os.Open("./text.txt")
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(line)
			fmt.Println("文件读取完毕")
			break
		}
		if err != nil {
			fmt.Printf("err:%v", err)
			return
		}
		fmt.Println(line)

	}
}


// ioReadAll ioutil read all
func ioReadAll(){
	content, err := ioutil.ReadFile("text.txt")
	if err != nil{
		fmt.Printf("err:%v",err)
		return
	}
	fmt.Println(string(content))
}

// 文件写入操作
func osWrite(){
	file, err := os.OpenFile("./text.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("err:%s", err)
		return
	}
	defer file.Close()

	str := "hello world"
	num, err := file.Write([]byte(str))
	if err != nil{
		fmt.Printf("err:%s", err)
		return
	}
	println(num)
	n, err := file.WriteString(str)
	if err != nil{
		fmt.Printf("err:%s", err)
		return
	}
	fmt.Println(n)
}

// bufio wirte
func bufioWrite(){
	file, err := os.OpenFile("./text.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("err:%s", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i:=0; i<=2;i++ {
		str := "hello world"
		writer.WriteString(str)
	}
	writer.Flush()
}

// ioutl write
func ioutlilWrite(){
	str := "hello world"
	err :=ioutil.WriteFile("./test.txt", []byte(str), 0644)
	if err != nil{
		fmt.Printf("write file failed:%v", err)
		return
	}

}

func CopyFile(srcFile string, dstFile string)(written int64, err error){
	src,err := os.Open(srcFile)
	if err != nil{
		fmt.Printf("open src file failed:%s", err)
	}
	defer src.Close()
	dst, err := os.OpenFile(dstFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		fmt.Printf("write dst file failed: %s", dst)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func main() {
	_, err := CopyFile("text.txt", "src.txt")
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}
	fmt.Println("copy done!")
}
