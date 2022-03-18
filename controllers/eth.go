package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"math/big"
	"myproject/controllers/utils"
	"strings"
)

//myTestStructSession
var myTestStructSession = utils.GetContract()
var accountNum, _ = myTestStructSession.AccoutNum()
var accountList = UpdateAccount(myTestStructSession)
var voteInfo = GetInfo(myTestStructSession)

func GetInfo(myTestStructSession *utils.TestStructSession) []map[string]string {
	var voteInfo []map[string]string
	for i := 0; i < int(accountNum); i++ {
		temp, _ := myTestStructSession.GetHaoMa(accountList[i])
		if temp != "-1" {
			temp2 := map[string]string{
				"account": accountList[i],
				"votenum": temp,
			}
			voteInfo = append(voteInfo, temp2)
		}
	}
	return voteInfo
}
func UpdateAccount(myTestStructSession *utils.TestStructSession) []string {
	var accountList []string
	for i := 0; i < int(accountNum); i++ {
		temp, _ := myTestStructSession.AccountList(big.NewInt(int64(i)))
		accountList = append(accountList, temp)
	}
	return accountList
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
func (c *MainController) Register() {
	fmt.Println(accountNum, accountList)
	c.TplName = "register.html"
}
func (c *MainController) RegisterSubmit() {
	var msg string
	username := c.GetString("username")
	password := c.GetString("password")
	temp, _ := myTestStructSession.ValidName(username)
	if temp {
		msg = "用户名已存在"
	} else {
		transaction, _ := myTestStructSession.Register(username, password)
		fmt.Println("用户已注册，交易hash为：", transaction.Hash().Hex()) //打印交易hash值
		accountNum, _ = myTestStructSession.AccoutNum()
		accountList = UpdateAccount(myTestStructSession)
		msg = "success"
	}
	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"msg":  msg,
	}
	c.ServeJSON()

}
func (c *MainController) Submit() {
	var msg string
	username := c.GetString("username")
	password := c.GetString("password")
	onenum := c.GetString("onenum")
	twonum := c.GetString("twonum")
	threenum := c.GetString("threenum")
	fournum := c.GetString("fournum")
	fivenum := c.GetString("fivenum")
	numList := []string{onenum, twonum, threenum, fournum, fivenum}
	numString := strings.Join(numList, "")
	temp, _ := myTestStructSession.ValidName(username)
	if temp {
		temp2, _ := myTestStructSession.ValidPassword(username, password)
		if temp2 {
			temp3, _ := myTestStructSession.GetHaoMa(username)
			if temp3 == "-1" {
				transaction, _ := myTestStructSession.BuyBallot(username, numString)
				fmt.Println("投注成功，交易hash为：", transaction.Hash().Hex()) //打印交易hash值
				voteInfo = GetInfo(myTestStructSession)
				msg = "success"
			} else {
				msg = "该用户已经投注"
			}
		} else {
			msg = "密码错误"
		}

	} else {
		msg = "用户名不存在"
	}
	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"msg":  msg,
	}
	c.ServeJSON()

}
func (c *MainController) GetAccounts() {

	retMap := map[string]interface{}{
		"accounts": voteInfo,
	}
	c.Data["json"] = retMap
	c.ServeJSON()
}
