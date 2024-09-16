package main

/*
二、读取文件（方法2）bufio 读取文件

		1、只读方式打开文件 file,err := os.Open()

		2、创建reader对象  reader := bufio.NewReader(file)

		3、ReadString读取文件  line, err := reader.ReadString('\n')

		4、关闭文件流 defer file.Close()
*/

func main() {

	//bufio 读取文件

}
