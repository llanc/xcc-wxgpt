package ability

import "fmt"

// SetPeriodicReminds 提供设置固定周期任务的能力
// time 固定周期,是一个大于0的整数秒数
// message 剥离掉时间后的提醒内容，需要适当调整以让实参看起来更像是一个来自助理的提醒
func SetPeriodicReminds(args ...string) {
	time := args[0]
	message := args[1]
	fmt.Printf("%v - %v", time, message)
}

// SetDisposableReminds 提供设置一次性提醒的能力
// timeType time的类型，取值1时代表time是相对0点的偏移秒数,取值2时代表time是相对当前时间的偏移秒数
// time 与type相对应，与type相对应，是相对0点的偏移秒数(要注意非今日概念,如明天8点应该是115200)或相对当前时间的偏移秒数
// message 剥离掉时间后的提醒内容，需要适当调整以让实参看起来更像是一个来自助理的提醒
func SetDisposableReminds(args ...string) {
	timeType := args[0]
	time := args[1]
	message := args[2]
	fmt.Printf("%v - %v - %v", timeType, time, message)
}
