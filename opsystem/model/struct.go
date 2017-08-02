package model

import (
	"log"

	"fmt"

	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Juris struct {
	Jid  int    `xorm:"not null pk INT(11)"`
	Name string `xorm:"VARCHAR(255)"`
}

type Product struct {
	Pid  int    `xorm:"not null pk autoincr INT(11)"`
	Name string `xorm:"VARCHAR(255)"`
}

type Sell struct {
	Id    int    `xorm:"not null pk autoincr INT(11)"`
	Pid   int    `xorm:"index INT(11)"`
	Place string `xorm:"VARCHAR(255)"`
	Num   int    `xorm:"INT(11)"`
	Time  string `xorm:"VARCHAR(255)"`
}

type User struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	Username string `xorm:"VARCHAR(255)"`
	Password string `xorm:"VARCHAR(255)"`
	Nickname string `xorm:"VARCHAR(255)"`
	Sex      int    `xorm:"INT(1)"`
	Age      string `xorm:"VARCHAR(8)"`
	Jid      int    `xorm:"index INT(1)"`
}

var x *xorm.Engine

func init() {
	var err error
	x, err = xorm.NewEngine("mysql", "root:root@(127.0.0.1:3306)/opsystem")
	if err != nil {
		log.Fatalf("fail to create engine: %v", err)
	}
	//自动同步表结构
	err = x.Sync2(new(Juris), new(Product), new(Sell), new(User))
	if err != nil {
		log.Fatalf("fail to sync database : %v", err)
	}
}

//user开始
func AddUser(m *User) error {
	fmt.Printf("the answer is : %v\n", m)
	id, err := x.Insert(m)
	fmt.Printf("the id is :%d\n", id)
	return err
}

func DeleteUSer(id int) error {
	_, err := x.Delete(&User{Id: id})
	return err
}

func UpdateUser(m *User) error {
	_, err := x.Id(m.Id).Update(m)
	return err
}

func GetUserByUsername(username string) (*User, error) {
	user := &User{Username: username}
	has, err := x.Get(user)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("no data")
	} else {
		return user, nil
	}
}

func GetUserInfo(id int) (*User, error) {
	user := &User{Id: id}
	has, err := x.Cols("nickname", "age", "sex", "jid", "username").Get(user)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("no data")
	} else {
		return user, nil
	}
}

//user结束

//product开始
func AddProduct(m *Product) error {
	_, err := x.Insert(m)
	return err
}

func GetAllProds() ([]*Product, error) {
	var prods []*Product
	err := x.Find(&prods)
	if err != nil {
		return nil, err
	}
	return prods, nil
}

func GetProdsCount() (int64, error) {
	prod := new(Product)
	total, err := x.Count(prod)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetProdsByPage(current int, pagesize int) ([]*Product, error) {
	var prods []*Product
	err := x.Limit(pagesize, (current-1)*pagesize).Find(&prods)
	if err != nil {
		return nil, err
	}
	return prods, nil
}

func UpdateProd(prod *Product) error {
	_, err := x.Id(prod.Pid).Update(prod)
	return err
}

func DeleteProd(id int) error {
	_, err := x.Delete(&Product{Pid: id})
	return err
}

//product结束

//sell开始
func AddSell(m *Sell) error {
	_, err := x.Insert(m)
	return err
}
func CheckAlreadySet(m *Sell) (bool, error) {
	v := &Sell{Pid: m.Pid, Time: m.Time, Place: m.Place}
	has, err := x.Get(v)
	if err != nil {
		return false, err
	} else if !has {
		return false, nil
	}
	return true, nil
}

func GetSellCount() (int64, error) {
	sell := new(Sell)
	total, err := x.Count(sell)
	if err != nil {
		return 0, err
	}
	return total, nil
}

type SellInfo struct {
	Sell `xorm:"extends"`
	Name string
}

//获取带商品名的销售
func GetSellByPage(current, pagesize int) ([]*SellInfo, error) {
	sells := make([]*SellInfo, 0)
	err := x.Sql("SELECT s.*,p.name FROM sell s,product p WHERE s.pid = p.pid limit ?,?", (current-1)*pagesize, pagesize).Find(&sells)
	if err != nil {
		return nil, err
	}
	return sells, nil
}

func DeleteSell(id int) error {
	_, err := x.Delete(&Sell{Id: id})
	return err
}

func UpdateSell(m *Sell) error {
	_, err := x.Id(m.Id).Update(m)
	return err
}

type SellNums struct {
	Name  string
	Value int64
}

func GetTotalSellNums() ([]*SellNums, error) {
	prods, err := GetAllProds()
	nums := make([]*SellNums, 0)
	if err == nil {
		for _, v := range prods {
			sell := new(Sell)
			total, err := x.Where("pid = ?", v.Pid).SumInt(sell, "num")
			if err == nil {
				num := &SellNums{Name: v.Name, Value: total}
				nums = append(nums, num)
			} else {
				return nil, err
			}
		}
		return nums, nil
	}
	return nil, err
}

type PlaceSell struct {
	Place string
	Total int64
}

func GetTotalSellPlace(id int, date string) ([]*PlaceSell, error) {
	nums := make([]*PlaceSell, 0)
	err := x.Sql("SELECT place,SUM(num) total FROM sell WHERE place IN ( SELECT place FROM sell GROUP BY place HAVING COUNT(place) > 0 ) AND pid = ? AND time LIKE ? GROUP BY place", id, "%"+date+"%").Find(&nums)
	if err == nil {
		return nums, nil
	}
	return nil, err
}

type DateSell struct {
	Time  string
	Total int64
}

func GetProdSellMonthly(id int, date string) ([]*DateSell, error) {
	sells := make([]*DateSell, 0)
	err := x.Sql("SELECT time, SUM(num) total FROM sell WHERE time IN ( SELECT time FROM sell WHERE pid = ? GROUP BY time HAVING COUNT(time) > 0 ) AND pid = ? AND time LIKE ? GROUP BY time", id, id, "%"+date+"%").Find(&sells)
	if err == nil {
		return sells, nil
	}
	return nil, err
}

//sell结束

//juris开始
func GetJuris() ([]*Juris, error) {
	jurieses := make([]*Juris, 0)
	err := x.Find(&jurieses)
	if err != nil {
		return nil, err
	}
	return jurieses, nil
}

//juris结束

//获取带权限名称的user列表
type UserGroup struct {
	User  `xorm:"extends"`
	Juris `xorm:"extends"`
}

func (UserGroup) TableName() string {
	return "user"
}

func GetTotalUserInfo() ([]*UserGroup, error) {
	users := make([]*UserGroup, 0)
	fmt.Println("method going 1")
	err := x.Join("inner", "juris", "juris.jid = user.jid").Find(&users)
	fmt.Println("method going 2")
	if err != nil {
		return nil, err
	}
	return users, nil
}

//获取带权限名称的user列表结束
