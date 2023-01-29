
go mock的使用
    1. 生成 mock 文件，一般会mock接口，例如mock person.go中的person接口和方法，命令如下：
        mockgen -source=./gomock/person.go -destination=./gomock/mock_person.go -package=gomock
            -source：设置需要模拟（mock）的接口文件
            -destination：设置 mock 文件输出的地方，若不设置则打印到标准输出中
            -package：设置 mock 文件的包名，若不设置则为 mock_ 前缀加上文件名（如本文的包名会为 mock_person）
      在执行完毕后，可以发现 gomock/ 目录下多出了 /mock_person.go 文件，这就是 mock 文件。

    2. 常用 mock 方法
       调用方法
           Call.Do()：声明在匹配时要运行的操作
           Call.DoAndReturn()：声明在匹配调用时要运行的操作，并且模拟返回该函数的返回值
           Call.MaxTimes()：设置最大的调用次数为 n 次
           Call.MinTimes()：设置最小的调用次数为 n 次
           Call.AnyTimes()：允许调用次数为 0 次或更多次
           Call.Times()：设置调用次数为 n 次
       参数匹配
           gomock.Any()：匹配任意值
           gomock.Eq()：通过反射匹配到指定的类型值，而不需要手动设置
           gomock.Nil()：返回 nil
       建议更多的方法可参见 官方文档: https://pkg.go.dev/github.com/golang/mock/gomock#pkg-index