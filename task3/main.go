package main

import (
	"errors"
	"fmt"
	"log"
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

	// students := []model.Students{}
	// model.DB.Where("age > ?", 18).Find(&students)
	// fmt.Printf("学生大于18岁列表%#v\n", students)

	// model.DB.Model(&model.Students{}).Where("name = ?", "张三").Update("grade", "四年级")

	// model.DB.Where("age < ?", 15).Delete(&model.Students{})

	//SQL语句练习 题目2 事务操作
	/*假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
	要求 ：
	编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。*/

	// accountA := model.Accounts{Balance: 50}
	// model.DB.Create(&accountA)
	// accountA1 := model.Accounts{Balance: 200}
	// model.DB.Create(&accountA1)
	// accountB := model.Accounts{Balance: 200}
	// model.DB.Create(&accountB)
	// var amount int = 100

	//A账户余额不足，回滚事务
	// errCommit := TransferMoney(model.DB, accountA, accountB, amount)
	// if errCommit != nil {
	// 	fmt.Println("errCommit执行报错,事务回滚")
	// }
	// A1账户余额足够，转账成功
	// errcommit2 := TransferMoney(model.DB, accountA1, accountB, amount)
	// if errcommit2 != nil {
	// 	fmt.Println("errcommit2执行报错,事务回滚")
	// }

	//sqlx入门  题目1使用SQL扩展库进行查询
	/*假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
	要求 ：
	编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。*/

	employees := []model.Employees{}
	err6 := model.SqlxDb.Select(&employees, `
	SELECT id, name, salary, department
	FROM employees
	WHERE department = ?
`, "技术部")
	if err6 != nil {
		fmt.Println("查询失败", err6)
	} else {
		fmt.Printf("技术部的员工%#v \n", employees)
	}
	topEmployees := []model.Employees{}
	err7 := model.SqlxDb.Select(&topEmployees, `
	SELECT id, name, salary, department
	FROM employees
	WHERE salary = (SELECT MAX(salary) FROM employees)
`)

	if err7 != nil {
		fmt.Println("查询失败", err7)
	} else {
		fmt.Printf("工资最高的员工%#v \n", topEmployees)
	}

	//sqlx入门  实现类型安全映射
	/*假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	要求 ：
	定义一个 Book 结构体，包含与 books 表对应的字段。
	编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。*/

	books := []model.Books{}
	err8 := model.SqlxDb.Select(&books, `
	SELECT id, title, author, price
	FROM books
	WHERE price > ?`, 50)

	if err8 != nil {
		fmt.Println("查询失败", err8)
	} else {
		fmt.Printf("价格大于50的书籍%#v \n", books)
	}

	/*
			进阶gorm
		题目1：模型定义
		假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
		要求 ：
		使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
		编写Go代码，使用Gorm创建这些模型对应的数据库表。
	*/
	/*
		题目2：关联查询
		基于上述博客系统的模型定义。
		要求 ：
		编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
		编写Go代码，使用Gorm查询评论数量最多的文章信息。*/
	queryUserPostsWithComments(model.Blogdb, 1)
	queryPostWithMostComments(model.Blogdb)

	/*
		题目3：钩子函数
		继续使用博客系统的模型。
		要求 ：
		为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
		为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。*/

	userTest := model.User{Name: "刘德华2", Email: "12啊#@", PostCount: 0}
	model.Blogdb.Create(&userTest)
	post1 := model.Posts{Title: "标题刘德华2", Content: "s啊啊dgdas", AuthorID: userTest.ID, CommentStatus: ""}
	post2 := model.Posts{Title: "标题刘德华2", Content: "sdfsfd", AuthorID: userTest.ID, CommentStatus: ""}
	model.Blogdb.Create(&post1)
	model.Blogdb.Create(&post2)
	comment1 := model.Comments{Content: "评论8", PostID: post1.ID}
	comment2 := model.Comments{Content: "评论8", PostID: post1.ID}
	comment3 := model.Comments{Content: "评论9", PostID: post2.ID}
	model.Blogdb.Create(&comment1)
	model.Blogdb.Create(&comment2)
	model.Blogdb.Create(&comment3)

	model.Blogdb.Where("id = ?", comment3.ID).Delete(&comment3)

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

func queryUserPostsWithComments(db *gorm.DB, userID uint) {
	var user model.User

	// 使用 Preload 一次性加载该用户的所有文章，以及每篇文章的所有评论
	err := db.Preload("Posts").Preload("Posts.Comments").First(&user, userID).Error
	if err != nil {
		if err.Error() == "record not found" {
			fmt.Printf("未找到 ID 为 %d 的用户\n", userID)
		} else {
			log.Fatalf("查询失败: %v", err)
		}
		return
	}

	fmt.Printf("用户信息: ID=%d, 姓名=%s, 邮箱=%s\n", user.ID, user.Name, user.Email)

	// 遍历该用户的所有文章
	for _, post := range user.Posts {
		fmt.Printf("\n文章 ID: %d, 标题: %s, 内容预览: %.50s...\n", post.ID, post.Title, post.Content)

		// 遍历该文章的所有评论
		commentCount := len(post.Comments)
		fmt.Printf("该文章共有 %d 条评论:\n", commentCount)
		for _, c := range post.Comments {
			fmt.Printf("评论ID: %d, 内容: %.50s...\n", c.ID, c.Content)
		}
	}
}

func queryPostWithMostComments(db *gorm.DB) {
	type PostWithCommentCount struct {
		Id           uint   `gorm:"column:id"`
		Title        string `gorm:"column:title"`
		Content      string `gorm:"column:content"`
		AuthorID     uint
		CommentCount int // 评论数量
	}

	var result PostWithCommentCount

	err := db.Table("posts").
		Select(`
			posts.id as id,
			posts.title,
			posts.content,
			posts.author_id,
			COUNT(comments.id) as comment_count
		`).
		Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&result).Error

	if err != nil {
		if err.Error() == "record not found" || err == gorm.ErrRecordNotFound {
			fmt.Println("暂无文章或评论数据")
		} else {
			log.Fatalf("查询评论最多的文章失败: %v", err)
		}
		return
	}

	fmt.Printf("\n评论数量最多的文章:\n")
	fmt.Printf("文章ID: %d\n", result.Id)
	fmt.Printf("标题: %s\n", result.Title)
	fmt.Printf("内容预览: %.50s...\n", result.Content)
	fmt.Printf("评论总数: %d\n", result.CommentCount)
	fmt.Printf("作者ID: %d\n", result.AuthorID)
}
