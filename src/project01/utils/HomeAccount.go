package utils

import (
	"fmt"
	"strings"
)

type HomeAccount struct {
	//菜单序号 1-4
	key string
	//账单总金额
	balance float64
	//收支金额
	money float64
	//收支明细
	note string
	//收支明细  拼接得到
	details string
}

//工厂模式
func NewHomeAccount() *HomeAccount {
	return &HomeAccount{
		key: "",
		details : "收支\t账户金额\t收支金额\t收支说明",
		balance : 10000.0,
		money : 0.0,
		note : "",
	}
}

//退出程序
func (this *HomeAccount) exit() {
	isExit := ""
	for {
		fmt.Println("是否退出程序（Y/N）")
		fmt.Scanln(&isExit)
		if strings.EqualFold(isExit , "Y") {  //== 
			fmt.Println("---您已经退出程序---")
			break
		}
	}
}

//收入明细
func (this *HomeAccount) showDetails(flag bool){
	fmt.Println("-----------------当前收支明细-------------------")
	if !flag {
		fmt.Println("还没收支，快来添加一笔吧。。。。")
	} else {
		fmt.Println(this.details)
	}	

}
func (this *HomeAccount) income(){
   fmt.Println("本次收入金额:")
   fmt.Scanln(&this.money)
   this.balance += this.money
   fmt.Println("本次收入说明:")
   fmt.Scanln(&this.note)
   this.details += fmt.Sprintf("\n收入\t %v \t\t %v \t\t %v",this.balance,this.money,this.note)
}

func (this *HomeAccount) pay(){
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足。。。")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t %v \t\t %v \t\t %v",this.balance,this.money,this.note)
}

//显示主菜单
func (this *HomeAccount) MainMenu() {
	loop := true
	flag := false //是否有收支
	for {
		fmt.Println("\n-----------------家庭收支记账软件-------------------")
		fmt.Println("                   1.收支明细")
		fmt.Println("                   2.登记收入")
		fmt.Println("                   3.登记支出")
		fmt.Println("                   4.退    出")
		fmt.Print("                   请选择(1-4):")
		fmt.Scanln(&this.key)
		switch this.key {
			case "1":
				this.showDetails(flag)
			case "2":
				flag = true
				this.income()
			case "3":
				flag = true
				this.pay()
			case "4":
				loop = false
			default:
				fmt.Println("输入菜单有误！！！！")
		}
		//退出程序
		if !loop {
			this.exit()
			break
		}
	}

}
