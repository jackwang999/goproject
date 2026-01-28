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
	fmt.Println(strings.Replace("helloll","ll","zz",-1))
}