package main

import "fmt"

func main() {

	switch finger := 3; finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	default:
		fmt.Println("无效输入")
	}
	testSwitch3()
	switchDemo4()
	switchDemo5()
	gotoDemo()
	gotoDemo2()
	continueDemo()
}

//case 分支可以使用多个值
func testSwitch3() {
	switch n := 7; n {
	case 1, 3, 5, 7:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

//省略condition写法
func switchDemo4() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受生活")
	default:
		fmt.Println("活着真好")
	}

}

//fallthrough可以执行满足条件语句的下一个case
func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

//跳到指定标签

func gotoDemo() {
	var breakflag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				breakflag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		if breakflag {
			break
		}
	}

}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
