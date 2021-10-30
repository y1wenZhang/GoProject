package main

import (
	"fmt"
	"os"
)

// 学员信息管理系统

// 添加学员信息
// 编辑学员信息
// 展示学员信息

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1. 添加学员")
	fmt.Println("2. 编辑学员信息")
	fmt.Println("3. 展示学员信息")
	fmt.Println("4. 退出")
}

func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请输入学员学号:")
	fmt.Scanf("%d", &id)
	fmt.Println("请输入学员姓名:")
	fmt.Scanf("%s", &name)
	fmt.Println("请输入学员班级:")
	fmt.Scanf("%s", &class)
	stu := newStudent(id, name, class)
	return stu
}

func main() {
	sm := newStudentManage()
	for {
		// 展示菜单
		showMenu()
		fmt.Println("请选择要执行的操作序号：")
		var input int
		fmt.Scanf("%d \n", &input)

		switch input {
		case 1:
			// 添加学员信息
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			// 编辑学员信息
			stu := getInput()
			sm.updateStudent(stu)
		case 3:
			// 展示学员信息
			sm.showStudent()
		case 4:
			// 退出
			os.Exit(0)

		}
	}
}
