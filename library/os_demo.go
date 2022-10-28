package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// os库相关
func OsDemo() {
	// 创建名称为test的目录，权限设置是perm，例如0777
	os.Mkdir("test", 0777)
	// 根据path创建多级子目录，例如test/test1/test2。
    os.MkdirAll("test/test1/test2", 0777)
    
	// 删除名称为test的目录，当目录下有文件或者其他目录是会出错
	os.Remove("test")
	// 根据path删除多级子目录，如果path是单个名称，那么该目录下的子目录全部删除。
	os.RemoveAll("test")

	userFile := "test.txt"
	// Go语言里面删除文件和删除文件夹是同一个函数
	defer os.Remove(userFile)

	// 根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的。
	// 如果有该文件，会覆盖这个文件，生成一个新的文件
    fout, _ := os.Create(userFile)
    defer fout.Close()
    for i := 0; i < 1; i++ {
		// 写入string信息到文件
        fout.WriteString("Just a test!\r\n")
		// 写入byte类型的信息到文件
        fout.Write([]byte("Just a test!\r\n"))
    }
	// 缓冲流写入数据
	write := bufio.NewWriter(fout)
	for i := 0; i < 2; i++ {
		// 把内容写入到缓冲区
		write.WriteString("Just a test!\r\n")
		// 强行把缓冲区的内容刷到磁盘去
		write.Flush() 
	}

	// 该方法打开一个名称为name的文件，但是是只读方式，内部实现其实调用了OpenFile。
	// func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
	// 		打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限
    fl, _ := os.Open(userFile)        
    defer fl.Close()
    buf := make([]byte, 1024)
    for {
        n, err := fl.Read(buf)
        if err == io.EOF {
			fmt.Println("读取到文件末尾了...")
            break
        }
        os.Stdout.Write(buf[:n])
    }

	// 从开始位置重新读
	fl.Seek(0, 0)
	// 带缓冲区的
	reader := bufio.NewReader(fl)
	for {
		// 读到\n停止，说明读取了一行数据
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取到文件末尾了...")
            break
        } else {
			// line中包含换行符
			fmt.Print(line)
		}
	}
}
