package main

import (
	"errors"
	"fmt"
	"task3/model"

	"gorm.io/gorm"
)

type APIUser struct {
	Age int
}

func main() {
	/*
		// 新增
		user := model.User{
			Name:    "增加的人",
			Age:     16,
			Sex:     "男",
			Address: "地址",
		}

		model.DB.Create(&user)

		//修改
		model.DB.Model(&user).Where("name = ?", "增加的人").Update("address", "修改地址")

		// 删除
		model.DB.Where("name = ?", "增加的人").Delete(&user)
		//删除--执行原生SQL
		// model.DB.Exec("DELETE from user where name=?", "增加的人")

		//查询
		//查询第一条
		user1 := model.User{}
		model.DB.First(&user1)
		fmt.Printf("查询到的user%v \n", user1)

		//按照条件查询 等于
		users1 := []model.User{}
		model.DB.Where("name=?", "张三").Find(&users1)
		fmt.Printf("name = 条件查询到的user列表%#v \n", users1)
		//条件查询原生写法
		users2 := []model.User{}
		model.DB.Raw("select * from user where name=?", "张三").Scan(&users2)
		fmt.Printf("原生sql查询name ==条件查询到的user列表%#v \n", users2)

		//IN条件查询
		users3 := []model.User{}
		model.DB.Where("name IN ?", []string{"张三", "李四"}).Find(&users3)
		fmt.Printf("name IN条件查询到的user列表%#v \n", users3)
	*/

	//SQL语句练习 题目一 基础CRUD操作
	/*假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
	要求 ：
	编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。*/

	//新增学生
	// student := model.Students{Name: "张三", Age: 20, Grade: "三年级"}
	// model.DB.Create(&student)
	// student2 := model.Students{Name: "王五", Age: 13, Grade: "一年级"}
	// model.DB.Create(&student2)
	// model.DB.Exec("INSERT INTO students (`name`, `age`, `grade`) VALUES ('李四', 18, '三年级')")

	students := []model.Students{}
	model.DB.Where("age > ?", 18).Find(&students)
	fmt.Printf("学生大于18岁列表%#v\n", students)

	model.DB.Model(&model.Students{}).Where("name = ?", "张三").Update("grade", "四年级")

	model.DB.Where("age < ?", 15).Delete(&model.Students{})

	//SQL语句练习 题目2 事务操作
	/*假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
	要求 ：
	编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。*/

	accountA := model.Accounts{Balance: 50}
	model.DB.Create(&accountA)
	accountA1 := model.Accounts{Balance: 200}
	model.DB.Create(&accountA1)
	accountB := model.Accounts{Balance: 200}
	model.DB.Create(&accountB)
	var amount int = 100

	//A账户余额不足，回滚事务
	errCommit := TransferMoney(model.DB, accountA, accountB, amount)
	if errCommit != nil {
		fmt.Println("errCommit执行报错,事务回滚")
	}
	// A1账户余额足够，转账成功
	errcommit2 := TransferMoney(model.DB, accountA1, accountB, amount)
	if errcommit2 != nil {
		fmt.Println("errcommit2执行报错,事务回滚")
	}

	//sqlx入门  题目1使用SQL扩展库进行查询
	/*假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
	要求 ：
	编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。*/

	//sqlx入门  实现类型安全映射
	/*假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	要求 ：
	定义一个 Book 结构体，包含与 books 表对应的字段。
	编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。*/

}

func TransferMoney(db *gorm.DB, accountA model.Accounts, accountB model.Accounts, amount int) error {
	// 开启事务
	err := db.Transaction(func(tx *gorm.DB) error {
		// 1. 查询转出账户 A，同时锁定该行（防止并发修改，使用 SELECT ... FOR UPDATE）
		// if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&accountA, fromAccountID).Error; err != nil {
		// 	fmt.Println(err) // 账户不存在或其他错误
		// }

		// 检查余额是否足够
		if accountA.Balance < amount {
			fmt.Println("余额不足")
			return errors.New("余额不足")
		}

		// 更新账户 A：扣除 amount
		newBalanceA := accountA.Balance - amount
		err3 := tx.Model(&model.Accounts{}).Where("id = ?", accountA.Id).Update("balance", newBalanceA).Error
		if err3 != nil {
			return err3
		}

		// 更新账户 B：增加 amount
		newBalanceB := accountB.Balance + amount
		err4 := tx.Model(&model.Accounts{}).Where("id = ?", accountB.Id).Update("balance", newBalanceB).Error
		if err4 != nil {
			return err4
		}
		// 全部成功，返回 nil 提交事务
		fmt.Println("事务执行完成")
		return nil
	})
	return err
}
