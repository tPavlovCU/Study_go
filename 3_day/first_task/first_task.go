package main

import "fmt"

var msu University

func mean(data []int) float64 {
	sum := float64(0)
	length := 0
	for _, v := range data {
		sum += float64(v)
		length++
	}
	return sum / float64(length)
}

type CourseInfo struct {
	Number         int
	Specialization string
}

type Student struct {
	ID     int
	Name   string
	Info   CourseInfo
	Grades []int
}
type University struct {
	Students []Student
}

// пока я помню, напишу в комментарии вопрос, я не придумал как иначе сохранять последний ID
func (u *University) AddStudent(name string, course int, spec string, grades []int) {
	students := u.Students
	Id := 0
	if len(students) == 0 {
		Id = 1
	} else {
		Id = students[len(students)-1].ID + 1
	}

	courseInfo := CourseInfo{course, spec}
	student := Student{Id, name, courseInfo, grades}
	students = append(students, student)
}

func (u University) FindBadStudents() (badStudents []Student, count int) {
	students := u.Students

	for _, student := range students {
		grades := student.Grades
		if mean(grades) < 3.5 {
			badStudents = append(badStudents, student)
			count++
		}
	}
	return badStudents, count
}

func (u *University) PromoteAll() {
	students := u.Students

	for _, student := range students {
		if student.Info.Number < 4 {
			student.Info.Number++
		} else {
			fmt.Printf("Студент %s успешно выпустился!", student.Name)
		}
	}
}

func initTest() {
	msu = University{
		Students: []Student{
			{
				ID:     1,
				Name:   "Иван Иванов",
				Info:   CourseInfo{Number: 1, Specialization: "Computer Science"},
				Grades: []int{5, 5, 4, 5}, // Отличник (Ср. балл: 4.75)
			},
			{
				ID:     2,
				Name:   "Мария Петрова",
				Info:   CourseInfo{Number: 2, Specialization: "Data Science"},
				Grades: []int{3, 2, 3, 4, 2}, // Двоечник (Ср. балл: 2.8) -> попадет в FindBadStudents
			},
			{
				ID:     3,
				Name:   "Алексей Сидоров",
				Info:   CourseInfo{Number: 4, Specialization: "Computer Science"},
				Grades: []int{4, 4, 5, 4}, // Выпускник (4 курс) -> отработает условие в PromoteAll
			},
			{
				ID:     4,
				Name:   "Елена Козлова",
				Info:   CourseInfo{Number: 3, Specialization: "Design"},
				Grades: []int{3, 3, 4, 3, 3}, // На грани (Ср. балл: 3.2) -> тоже попадет в FindBadStudents
			},
		},
	}

	fmt.Println("[Запущена функция init()]: Тестовые данные успешно загружены в систему.")
}

func main() {
	initTest()
	fmt.Println(msu)
	grades := []int{2, 3, 5, 4}
	msu.AddStudent("Tim", 1, "math", grades)
	fmt.Println(msu)

	msu.FindBadStudents()

	msu.PromoteAll()
	fmt.Println(msu)

}
