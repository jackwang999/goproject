package main
import (
	"fmt"
	"time"
	"strconv"
	"strings"
)
func main(){
	var age int
	age = 18
	fmt.Println("age = ", age)
	var n1 = "abc" + "ddf"
	var n2 int = +11
	fmt.Println(n1)
	fmt.Printf("%T\n",n1)
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
		fmt.Println(10,"wowo")
	case 20:
		fmt.Println(20,"haha")
	default:
		fmt.Println("默认1")
	}

	var dates [3]time.Time
	dates[0] = time.Unix(1250000000, 0)
	dates[1] = time.Unix(1250000000, 0)
	dates[2] = time.Unix(1250000000, 0)
	fmt.Println(dates[0])
    
	var arr [3]int16
	fmt.Printf("dates数组的地址是：%p  %p  %p\n",&arr,&arr[1],&arr[2])

	// var age1 int
	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age1)
	// fmt.Println("年龄为：", age1)

	var c0 byte = 'a'
	var c1 byte = 'b'
	var c2 int = '中'
	fmt.Println(c0,c1,c2)
	fmt.Printf("%c   %c\n",c0,c2)

	var n11 int = 10
	var n12 float32 = float32(n11)
	fmt.Println(n12)

	var n22 int = 18
	var n23 string = strconv.FormatInt(int64(n22),16)
    fmt.Printf("%T, %q",n23,n23)

	fmt.Println("ceshi")

	fmt.Println(strings.Index("helloll","ll"))
	fmt.Println(strings.Replace("helloll","ll","好学生",-1))

	shuzu1 := strings.Split("hello,world,ok",",")
	fmt.Println(shuzu1,shuzu1[0],shuzu1[1],shuzu1[2])

	str1 := strings.ToLower("HelloWorld")
	str2 := strings.ToUpper("HelloWorld")
	fmt.Println(str1)
	fmt.Println(str2)

	fmt.Println(strings.TrimSpace("   hello   "))
	fmt.Println(strings.Trim("!!!hello!!!","!"))
	fmt.Println(strings.TrimLeft("!!!hello!!!","!"))
	fmt.Println(strings.TrimRight("!!!hello!!!","!"))
    fmt.Println(strings.HasPrefix("http://www.baidu.com","http"))
	fmt.Println(strings.HasSuffix("http://www.baidu.com","com"))

	fmt.Println(time.Now())
	now := time.Now()
	fmt.Printf("%T\n",time.Now())
	fmt.Println(now.Year())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	str201 := now.Format("2006/01/02 15:04:05")
	fmt.Println(str201)

	var arr1 [5]int = [5]int{1,2,3,4,5}
	fmt.Println(arr1)

	var arr2 = [...]int{1,2,3,4,5,6}
	fmt.Println(arr2)

	var arr3 = [...]int{2:100,5:200}
	fmt.Println(arr3)
}