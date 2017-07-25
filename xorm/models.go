package main

import (
	"log"

	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"`
}

var x *xorm.Engine

func init() {
	var err error
	x, err = xorm.NewEngine("mysql", "root:root@(127.0.0.1:3306)/xorm")
	if err != nil {
		log.Fatalf("fail to create engine: %v", err)
	}
	//自动同步表结构
	err = x.Sync(new(Account))
	if err != nil {
		log.Fatalf("fail to sync database : %v", err)
	}
}
func NewAccount(name string, balance float64) error {
	_, err := x.Insert(&Account{Name: name, Balance: balance})
	return err
}

//要求返回指针地址
func GetAccount(id int64) (*Account, error) {
	a := &Account{Id: id}
	has, err := x.Id(id).Get(a)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("Account not found")
	}
	return a, nil
}

func MakeDeposit(id int64, deposit float64) (*Account, error) {
	a, err := GetAccount(id)
	if err != nil {
		return nil, err
	}
	a.Balance += deposit
	_, err = x.Update(a)
	return a, nil
}
func MakeWithdraw(id int64, withdraw float64) (*Account, error) {
	a, err := GetAccount(id)
	if err != nil {
		return nil, err
	}
	if a.Balance <= withdraw {
		return nil, errors.New("Not Enough Balance")
	}
	a.Balance -= withdraw
	_, err = x.Update(a)
	return a, nil
}

// func MakeTranfer(id1, id2 int64, balance float64) error {
// 	a1, err := GetAccount(id1)
// 	if err != nil {
// 		return err
// 	}
// 	a2, err := GetAccount(id2)
// 	if err != nil {
// 		return err
// 	}
// 	if a1.Balance <= balance {
// 		return errors.New("Not Enough balance")
// 	}
// 	//代码存在问题，需要用事务回滚来改进
// 	a1.Balance -= balance
// 	a2.Balance += balance
// 	if _, err = x.Update(a1); err != nil {
// 		return err
// 	} else if _, err = x.Update(a2); err != nil {
// 		return err
// 	}
// 	return nil
// }
//出现差错采用事务回滚
func MakeTranfer(id1, id2 int64, balance float64) error {
	a1, err := GetAccount(id1)
	if err != nil {
		return err
	}
	a2, err := GetAccount(id2)
	if err != nil {
		return err
	}
	if a1.Balance <= balance {
		return errors.New("Not Enough balance")
	}
	a1.Balance -= balance
	a2.Balance += balance

	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Update(a1); err != nil {
		sess.Rollback()
		return err
	} else if _, err = sess.Update(a2); err != nil {
		sess.Rollback()
		return err
	}
	return sess.Commit()
}
func GetAccountsAscId() (as []*Account, err error) {
	err = x.Asc("id").Find(&as)
	return as, err
}
func GetAccountDescBalance() (as []*Account, err error) {
	err = x.Desc("balance").Find(&as)
	return as, err
}
func DeleteAccount(id int64) error {
	_, err := x.Delete(&Account{Id: id})
	return err
}
