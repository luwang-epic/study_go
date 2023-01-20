// 一般包名应该和文件夹名一致，这里为了演示不一致外部如何使用
package api2

import "fmt"

func ExternalApi() {
	fmt.Println("不在internal文件夹里面的方法外部都可以访问")
}
