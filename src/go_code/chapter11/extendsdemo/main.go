package main

import "fmt"

//编写一个学生考试系统
type Student struct {
	Name  string
	Age   int
	Score float64
}

//将 Pupil 和 Graduate 共有的方法也绑定到 *studens
//显示具体信息
func (s *Student) ShowInfo() {
	fmt.Printf("学生名=%v  年龄=%v  成绩=%v\n", s.Name, s.Age, s.Score)
}

//设置成绩
func (s *Student) SetScore(score float64) {
	//业务判断
	s.Score = score
}

//小学生
type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试..........")
}

type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试..........")
}

//大学生,研究生
func main() {
	//测试
	//当我们对结构体嵌入了匿名结构体,其使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "Tom"
	pupil.Student.Age = 10
	pupil.testing()
	pupil.Student.SetScore(89)
	pupil.Student.ShowInfo()

	//大学生
	graduate := &Graduate{}
	graduate.Student.Name = "kangnan"
	graduate.Student.Age = 20
	graduate.testing()
	graduate.Student.SetScore(98)
	graduate.Student.ShowInfo()

}
