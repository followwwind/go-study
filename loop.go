package main


import "fmt"

func switchLoop(){
	/* 定义局部变量 */
	var grade  = "B"
	var marks  = 90

	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70 : grade = "C"
		default: grade = "D"
	}

	switch {
		case grade == "A" :
			fmt.Printf("优秀!\n" )
		case grade == "B", grade == "C" :
			fmt.Printf("良好\n" )
		case grade == "D" :
			fmt.Printf("及格\n" )
		case grade == "F":
			fmt.Printf("不及格\n" )
		default:
			fmt.Printf("差\n" )
	}
	fmt.Printf("你的等级是 %s\n", grade )

	//switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	var x interface{}

	switch i := x.(type) {
		case nil:
			fmt.Printf("x 的类型 :%T",i)
		case int:
			fmt.Printf("x 是 int 型")
		case float64:
			fmt.Printf("x 是 float64 型")
		case func(int) float64:
			fmt.Printf("x 是 func(int) 型")
		case bool, string:
			fmt.Printf("x 是 bool 或 string 型" )
		default:
			fmt.Printf("未知型")
	}
}

func forLoop(){
	arr := []int{1, 2, 3}

	for i,x := range arr{
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
		if i == 2 {
			fmt.Println(x)
		}
	}
}

func main(){
	forLoop()
	switchLoop()
}