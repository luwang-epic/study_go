package api3

import "fmt"

func InternalApi() {
	fmt.Println("在internal文件夹里面的方法外部访问不到")
}
