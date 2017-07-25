package main

import (
	"fmt"
)

const prompt = `please enter number of operation:
1.Create new account
2.Show detail of account
3.Deposit
4.Withdraw
5.Make transfer
6.List Account by id
7.List Account by balance
8.Delete account
9.Exit
`

func main() {
	fmt.Println("welcome bank of xorm")

Exit:
	for {
		fmt.Println(prompt)
		var num int
		fmt.Scanf("%d\n", &num)
		switch num {
		case 1:
			fmt.Println("Please enter <name> <balance>:")
			var name string
			var balance float64
			fmt.Scanf("%s %f \n", &name, &balance)
			if err := NewAccount(name, balance); err != nil {
				fmt.Println(err)
			}
		case 2:
			fmt.Println("Please enter <id>:")
			var id int64
			fmt.Scanf("%d\n", &id)
			a, err := GetAccount(id)
			if err != nil {
				fmt.Println(err)
			} else {
				//打印结构体
				fmt.Printf("%#v\n", a)
			}
		case 3:
			fmt.Println("Please enter <id> <deposit>:")
			var id int64
			var deposit float64
			fmt.Scanf("%d %f\n", &id, &deposit)
			a, err := MakeDeposit(id, deposit)
			if err != nil {
				fmt.Println(err)
			} else {
				//打印结构体
				fmt.Printf("%#v\n", a)
			}
		case 4:
			fmt.Println("Please enter <id> <withdraw>:")
			var id int64
			var withdraw float64
			fmt.Scanf("%d %f\n", &id, &withdraw)
			a, err := MakeWithdraw(id, withdraw)
			if err != nil {
				fmt.Println(err)
			} else {
				//打印结构体
				fmt.Printf("%#v\n", a)
			}
		case 5:
			fmt.Println("Please enter <id> <balance> <id>:")
			var id1, id2 int64
			var balance float64
			fmt.Scanf("%d %f %d\n", &id1, &balance, &id2)
			if err := MakeTranfer(id1, id2, balance); err != nil {
				fmt.Println(err)
			}
		case 6:
			as, err := GetAccountsAscId()
			if err != nil {
				fmt.Println(err)
			} else {
				for i, a := range as {
					fmt.Printf("%d: %#v\n", i, a)
				}
			}
		case 7:
			as, err := GetAccountDescBalance()
			if err != nil {
				fmt.Println(err)
			} else {
				for i, a := range as {
					fmt.Printf("%d: %#v\n", i, a)
				}
			}
		case 8:
			fmt.Println("Please enter <id> :")
			var id int64
			fmt.Scanf("%d\n", &id)
			if err := DeleteAccount(id); err != nil {
				fmt.Println(err)
			}
		case 9:
			break Exit
		}
	}

}
