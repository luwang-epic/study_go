package main

import (
	"fmt"
	"time"
)

// 时间相关的库
func TimeDemo() {
	now := time.Now()
	fmt.Println("当前  秒时间为:", now.Unix())
	fmt.Println("当前毫秒时间为:", now.UnixMilli())
	fmt.Println("当前微秒时间为:", now.UnixMicro())
	fmt.Println("当前纳秒时间为:", now.UnixNano())

	// 星期
	fmt.Println("当前星期几:", now.Weekday().String())
	fmt.Println("当前小时:", now.Hour())
	fmt.Println("当前分钟:", now.Minute())
	fmt.Println("当前秒数:", now.Second())
	fmt.Println("当前月份:", now.Month())
	// Month底层就是int类型：type Month int
	fmt.Println("当前月份:", int(now.Month()))
	fmt.Println("当前年:", now.Year())
	fmt.Println("当前年中第几天:", now.YearDay())
	fmt.Println("当前月中第几天:", now.Day())

	// 获取时间间隔
	duration := time.Since(now)
	// 这样也可以获取，推荐用上面的
	// duration := time.Now().Sub(now)
	fmt.Println("时间间隔秒为:", duration.Seconds())
	fmt.Println("时间间隔毫秒为:", duration.Microseconds())
	fmt.Println("时间间隔纳秒为:", duration.Nanoseconds())

	// Time 时刻  Duration 时间段
	// 时刻 + 时间段 = 另一个时刻
	// 时刻 - 时刻 = 时间段
	// 时刻 + 时刻 不行
	duration2Hour := time.Duration( 2 * time.Hour)
	after2Hour := now.Add(duration2Hour)
	fmt.Println("2小时后的时间为:", after2Hour.Unix())

	// 时间格式化，必须是2006-01-02 15:04:05，不能是其他数字
	fmt.Println("时间格式化:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("时间格式化:", now.Format("2006-01-02"))
	fmt.Println("时间格式化:", now.Format("20060102"))
	
	if t, err :=time.Parse("2006-01-02", "2022-08-10"); err == nil {
		fmt.Println("2022-08-10解析的时间为:", t.Year(), t.Month(), t.Day())
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	// 推荐使用这个，这样就和时区无关了
	if t, err :=time.ParseInLocation("2006-01-02", "2022-08-10", loc); err == nil {
		fmt.Println("2022-08-10解析的时间为:", t.Year(), t.Month(), t.Day())
	}

	// 每个多长时间执行一次
	tc := time.NewTicker(3 * time.Second)
	defer tc.Stop()
	for i := 0; i< 6; i++ {
		// 利用管道来实现的，每3s写入一次，没到时间会阻塞在这里
		t := <- tc.C
		fmt.Println("每个3s执行:", t)
	}

	// 多长时间后执行
	tm := time.NewTimer(3 * time.Second) // Second本质是Duration类型，所以可以直接传入
	defer tm.Stop()
	// 内部也是用管道来实现的，但是只能用一次，再使用会报错
	t := <- tm.C
	fmt.Println("3s后执行:", t)
	// 这种方式也可以实现 多长时间后执行
	t = <- time.After(3 * time.Second)
	fmt.Println("3s后执行:", t)

	// time可以通过reset来重置
	for i := 0; i < 6; i++ {
		tm.Reset(1 * time.Second)
		t = <- tm.C
		fmt.Println("每隔1s执行:", t)
	}

	// 睡眠
	time.Sleep(3 * time.Second)

	// 其他更复杂的定时功能，可以通过第三方的cron库来实现
}