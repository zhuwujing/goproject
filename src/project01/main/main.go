package main

import (
	"fmt"
	"strconv"
	// "project01/utils"
	//   D:\goProject\src\ project01\utils   注意斜线的方向  路径是src（默认加上）以后的 加载GOPATH/src/模块
)

func main() {
	/*
	  变量的使用：1.声明/定义变量  2.变量赋值  3.使用变量
	  变量三要素  变量= 变量名 + 值 + 数据类型
	*/
	var  a int = 10
	fmt.Printf("a = %v\n",a)
	var num1,num2,num3 int 
	fmt.Printf("num1=%d,num2=%d,num3=%d\n",num1,num2,num3)
	var b ,c  ,d = 12,"tom",15.0
	fmt.Printf("b %T  c %T d %T \n",b,c,d)

	//十进制 --> 2进制
	b1 := strconv.FormatInt(127,2)
	fmt.Printf("b1=%v",b1)
	// var a int  //基本
	// var b = 10  //类型推导
	// c := 20
	// str := "hello"   //不同的类型不能做运算
	// fmt.Println(a + b + c)
	// var str string = "hello world! 社长"
	// len := utils.Len(str)//包名.函数名（）
	//   fmt.Println("main strlen= ",len)
	//   fmt.Println(str)

	//   utils.Traverse(str);
	// utils.TraverseByRune(str)
	// a,err := utils.Atoi("123p")
	// a := utils.Itoa(13384141)
	// a := utils.Byte2Str([]byte{97,98,99})
	// a := utils.Str2Byte("abc")
	// a := utils.FormatInt(64,2)
	// b := utils.FormatInt(64,8)
	// fmt.Println("a = ",a)
	// fmt.Printf("b = %v b的类型 %T",b,b)

	// utils.NewHomeAccount().MainMenu()
}
