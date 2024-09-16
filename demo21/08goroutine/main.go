package main

// 循环遍历管道数据
func main() {
	//1、使用for range遍历通道，当通道被关闭的时候就会退出for range,如果没有关闭管道就会报个错误fatal error: all goroutines are asleep - deadlock!

	//2、通过for循环遍历管道的时候管道可以不关闭

}
