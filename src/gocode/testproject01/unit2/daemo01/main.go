package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
	"bufio"
	"io"
)

type Nanren struct {
	Person
}

func (nr *Nanren) smoke() {
	fmt.Println("男人抽烟")
}

type Person struct {
	Name string
	Age  int
}

func (p Person) sayHello() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}
func main() {
	var age int
	age = 18
	fmt.Println("age = ", age)
	var n1 = "abc" + "ddf"
	var n2 int = +11
	fmt.Println(n1)
	fmt.Printf("%T\n", n1)
	fmt.Println(&n2)
	var ptr *int = &n2
	fmt.Println("ptr地址代表的值：", *ptr)

	if n2 < 5 {
		fmt.Println("true,hahaha")
	} else {
		fmt.Println("okok")
	}

	switch n2 {
	case 10:
		fmt.Println(10, "wowo")
	case 20:
		fmt.Println(20, "haha")
	default:
		fmt.Println("默认1")
	}

	var dates [3]time.Time
	dates[0] = time.Unix(1250000000, 0)
	dates[1] = time.Unix(1250000000, 0)
	dates[2] = time.Unix(1250000000, 0)
	fmt.Println(dates[0])

	var arr [3]int16
	fmt.Printf("dates数组的地址是：%p  %p  %p\n", &arr, &arr[1], &arr[2])

	// var age1 int
	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age1)
	// fmt.Println("年龄为：", age1)

	var c0 byte = 'a'
	var c1 byte = 'b'
	var c2 int = '中'
	fmt.Println(c0, c1, c2)
	fmt.Printf("%c   %c\n", c0, c2)

	var n11 int = 10
	var n12 float32 = float32(n11)
	fmt.Println(n12)

	var n22 int = 18
	var n23 string = strconv.FormatInt(int64(n22), 16)
	fmt.Printf("%T, %q", n23, n23)

	fmt.Println("ceshi")

	fmt.Println(strings.Index("helloll", "ll"))
	fmt.Println(strings.Replace("helloll", "ll", "好学生", -1))

	shuzu1 := strings.Split("hello,world,ok", ",")
	fmt.Println(shuzu1, shuzu1[0], shuzu1[1], shuzu1[2])

	str1 := strings.ToLower("HelloWorld")
	str2 := strings.ToUpper("HelloWorld")
	fmt.Println(str1)
	fmt.Println(str2)

	fmt.Println(strings.TrimSpace("   hello   "))
	fmt.Println(strings.Trim("!!!hello!!!", "!"))
	fmt.Println(strings.TrimLeft("!!!hello!!!", "!"))
	fmt.Println(strings.TrimRight("!!!hello!!!", "!"))
	fmt.Println(strings.HasPrefix("http://www.baidu.com", "http"))
	fmt.Println(strings.HasSuffix("http://www.baidu.com", "com"))

	fmt.Println(time.Now())
	now := time.Now()
	fmt.Printf("%T\n", time.Now())
	fmt.Println(now.Year())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	str201 := now.Format("2006/01/02 15:04:05")
	fmt.Println(str201)

	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	var arr2 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr2)

	var arr3 = [...]int{2: 100, 5: 200}
	fmt.Println(arr3)

	var arr4 []int = []int{1, 2, 3}
	var slice []int = arr4[1:3]
	fmt.Println(slice)
	fmt.Printf("%T,slice长度：%d,slice容量：%d\n", slice, len(slice), cap(slice))

	var slice2 []int = make([]int, 5, 10)
	fmt.Println(slice2)
	fmt.Printf("slice2的长度是：%d，容量是：%d\n", len(slice2), cap(slice2))

	slice3 := []int{11, 22, 33, 44, 55}
	fmt.Println(slice3)
	slice4 := slice3[1:4]
	fmt.Println(slice4)
	slice4[0] = 222
	fmt.Println(slice3)
	fmt.Println(slice4)

	for k, v := range slice3 {
		fmt.Println(k, v)
	}

	var arr11 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", arr11)
	slice21 := arr11[1:]
	fmt.Printf("%v", slice21)
	arr12 := append(slice21, 6, 7, 8)
	fmt.Printf("%v", arr12)

	var slice31 = make([]int, 10)
	copy(slice31, slice21)
	fmt.Printf("slice31=%v\n", slice31)

	fmt.Println("-----map的使用-----")
	var a map[int]string //定义一个map
	a = make(map[int]string, 10)
	a[1] = "hello"
	a[2] = "world"
	fmt.Println(a)
	fmt.Printf("a的类型是：%T\n", a)

	b := make(map[int]string)
	b[10] = "ok"
	b[20] = "good"
	fmt.Println(b)

	c := map[int]string{
		100: "java",
		200: "go",
		300: "python",
	}
	fmt.Println(c)
	delete(c, 200)
	fmt.Println(c)

	value31, bool1 := c[100]
	fmt.Println("值是", value31, "是否查询到", bool1)

	for k, _ := range c {
		delete(c, k)
	}
	fmt.Println("C map:", c)

	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{"Bob", 25}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.Name, p1.Age)

	var p3 = new(Person)
	p3.Name = "Charlie" //本来应该是*p3
	p3.Age = 28
	fmt.Println(*p3)

	p1.sayHello()
	p2.sayHello()

	xiaonan := Nanren{}
	xiaonan.Name = "男人"
	xiaonan.Age = 35
	xiaonan.sayHello()
	xiaonan.smoke()

	// file, err := os.Create("ceshi1.txt")
	// if err != nil {
	// 	fmt.Println("创建文件失败:", err)
	// 	// 	return
	// 	// } else if _, err := file.WriteString("hello,王纪樟"); err != nil {
	// 	// 	fmt.Printf("写入文件失败: %v\n", err)
	// }

	// fmt.Printf("文件=%v\n", file)

	//err2 := file.Close()
	// if err2 != nil {
	// 	fmt.Println("关闭文件失败:", err2)
	// }
	content, err := ioutil.ReadFile("ceshi1.txt")
	if err != nil {
		fmt.Println("读取文件失败:", err, "\n")
	}
	fmt.Printf("%v", string(content))
	// for i, v := range content {
	// 	fmt.Printf("索引：%d，值：%v\n", i, string(v))
	// }
	//cd src/gocode/testproject01/unit2/daemo01

	file, err := os.Open("ceshi.txt")
	defer file.Close()    //语句压入栈，函数结束的时候执行
    reader := bufio.NewReader(file)       //带缓存区的读取文件
	for {
		line, err := reader.ReadString('\n')  //读取一行内容，直到遇到换行符
		if err == io.EOF {          //读取到了结尾
			break
		}
        fmt.Print(line) //输出读取到的一行内容
	}
	fmt.Println("文件读取成功，全部读取完毕")	


	// 写文件
	file1,err := os.OpenFile("D:/goproject/src/gocode/testproject01/unit2/daemo01/ceshidemo.txt", 
	os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
    defer file.Close()
	writer := bufio.NewWriter(file1)
	for i := 0; i < 10; i++ {
		writer.WriteString("这是第" + strconv.Itoa(i) + "行数据。\n")
	}
    writer.WriteString("Hello, World!\n")
    writer.Flush() //确保写入内容被刷新到文件中

	mode1 := os.FileMode(0777).String()
	fmt.Println("文件权限模式：", mode1)    //文件权限模式： -rwxrwxrwx
}