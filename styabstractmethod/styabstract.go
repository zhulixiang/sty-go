package styabstractmethod

import "fmt"

type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}


func OperateAccount(){

	var account = Account{"11111111","666666",100,}
	account.Deposit(100,"666666")
	account.Query("666666")
	account.WithDraw(80,"666666")
	account.Query("666666")
}




//存款
func (account *Account) Deposit(money float64,pwd string){

	if pwd != account.Pwd {
		fmt.Print("输入密码不对")
		return
	}

	if money < 0 {
		fmt.Println("存款金额不合法")
		return
	}
	account.Balance +=money
}
//取款
func (account *Account)WithDraw(money float64,pwd string){
	if pwd != account.Pwd {
		fmt.Print("输入密码不对")
		return
	}

	if money < 0  {
		fmt.Println("存款金额不合法")
		return
	}
	if money > account.Balance {
		 fmt.Println("余额不足")
		return
	}
	account.Balance -= money
}

func (account *Account)Query(pwd string){

	if pwd != account.Pwd {
		fmt.Print("输入密码不对")
		return
	}
	fmt.Println("余额:",account.Balance)
}



