package styextends

import "fmt"


type Hobby struct {
	Name string
	Hot string
}

type Student struct {
	Name  string
	Age   int
	Score float64
	sex string
}

func (student *Student) ShowInfo() {
	fmt.Printf("学生姓名 %v,年龄 %v,成绩 %v \n", student.Name, student.Age, student.Score)
}

func (student *Student) SetScore(score float64) {
	student.Score = score
}

func (student *Student) Say() {
	fmt.Println("匿名结构体say()")
}

//小学生
type Pupil struct {
	*Student
	*Hobby
	int
}

func (p *Pupil) testing() {
	fmt.Println("小学生考试。。。")
}

func (p *Pupil) Say() {
	fmt.Println("结构体say()")
}

//大学生
type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生考试。。。")
}

func Demo() {

	var p = &Pupil{}
	p.Student.Name = "jack"
	p.Student.Age = 12
	p.Student.SetScore(80)
	p.testing()
	p.Student.ShowInfo()

}
func Demo1() {
	var p = &Pupil{}
	p.Student = &Student{}
	p.Age = 13
	p.testing()
	p.SetScore(90)
	p.ShowInfo()
	p.Say()
	p.Student.Say()

	var p1 = &Pupil{&Student{Name: "ji",Age: 12,Score: 13},&Hobby{Name: "足球",Hot: "特别"},10}
	fmt.Println(p1)
}
