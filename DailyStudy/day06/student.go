package main

import "fmt"

type student struct {
	Id    int
	Name  string
	Class string
}

//newStudent Student类型的构造函数
func newStudent(id int, name string, class string) *student {
	return &student{
		Id:    id,
		Name:  name,
		Class: class,
	}
}

type studentManage struct {
	students []*student
}

func newStudentManage() *studentManage {
	return &studentManage{
		students: make([]*student, 0, 100),
	}
}

func (s *studentManage) addStudent(newStu *student) {
	s.students = append(s.students, newStu)
}

func (s *studentManage) updateStudent(newStu *student) {
	for index, stu := range s.students {
		if stu.Id == newStu.Id {
			s.students[index] = newStu
		}
	}
}

func (s *studentManage) showStudent() {
	for _, stu := range s.students {
		fmt.Printf("学号：%d 姓名: %s 班级：%s", stu.Id, stu.Name, stu.Class)
	}
}
