package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

// 日志功能
// log为go提供的基础包，但是只支持简单的日志功能，无法满足记录不同级别日志的情况以及写入文件等
// 如果需要实现复杂的功能，需要使用第三方包，如logrus
func LogDemo() {
	// 使用全局的log对象打印日志
	// 2022/10/27 09:07:43 这是一条正常的日志
	log.Println("这是一条正常的日志")

	// 设置日志的输出格式 会带上文件全路径，使用微秒时间 以及日期， 下面的样子
	// 2022/10/27 09:05:44.864228 D:/go_project/src/study_go/library/library.go:443: 这是一条正常的日志
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// 设置日志前缀
	log.SetPrefix("<prefix>")

	// 设置日志的输出
	// log.SetOutput(os.Stdout)
	// logFile, _ := os.OpenFile("log.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	// defer logFile.Close()
	// 设置输出到文件
	// log.SetOutput(logFile)

	log.Println("这是一条正常的日志")
	// 这个会直接导致程序退出，exit status 1
	// log.Fatalln("这是一条会触发fatal的日志")
	// 这个会导致程序抛出异常
	// log.Panicln("这是一条会触发panic的日志")

	// 也可以创建新的log对象
	mylog := log.New(os.Stdout, "", log.Ldate|log.Lshortfile|log.Ltime)
	mylog.Println("新创建的日志对象")

	// 设置logrus的配置
	// Log as JSON instead of the default ASCII formatter.
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{})
	// 输出stdout而不是默认的stderr，也可以是一个文件
	logrus.SetOutput(os.Stdout)
	// 设置日志级别
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("Something very low level.")
	logrus.Debug("Useful debugging information.")
	logrus.Info("Something noteworthy happened!")
	logrus.Warn("You should probably take a look at this.")
	logrus.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	// logrus.Fatal("Bye.")
	// Calls panic() after logging
	// logrus.Panic("I'm bailing.")

	// Field机制：logrus鼓励通过Field机制进行精细化的、结构化的日志记录，而不是通过冗长的消息来记录日志。
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := logrus.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")

	// 使用新的日志对象
	var newlog = logrus.New()
	newlog.Info("new log")
}

// zapLogDemo zap日志演示
func zapLogDemo() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Panicf("fail to new logger, err=%v", err)
	}

	defer func() {
		_ = logger.Sync() // flushes buffer, if any
	}()

	logger.Info("hello go", zap.String("a", "a1"))
}
