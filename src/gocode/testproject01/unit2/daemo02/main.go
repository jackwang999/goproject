package main
import "fmt"
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
		fmt.Println("默认")
	}
}