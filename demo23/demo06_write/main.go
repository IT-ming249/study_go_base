package main

/*
二、写入文件（方法2） bufio 写入文件

		1、打开文件  file, err := os.OpenFile("C:/test.txt", os.O_CREATE|os.O_RDWR, 0666)

		2、创建writer对象  writer := bufio.NewWriter(file)

		3、将数据先写入缓存  writer.WriteString("你好golang\r\n")

		4、将缓存中的内容写入文件	writer.Flush()

		5、关闭文件流 file.Close()
*/
func main() {

}
