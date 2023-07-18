package cliver

import (
	"bufio"
	"fmt"
	//"os"
	//"reflect"
	"strconv"
	//"strings"
	//"io"w
	//"log"
	//"net/http"
)

func createStudent(forename string, surname string, id int) Student {
	newStudent := Student{Id: id, Forename: forename, Surname: surname}
	return newStudent
}

func addGradeCmd(reader *bufio.Reader) {
	fmt.Println("== Adding a Grade to a Student ==")
	fmt.Print("| Student ID: ")
	id, _ := strconv.ParseInt(stdInput(reader), 10, 64)
	for i := 0; i < len(students); i++ {
		if students[i].Id == int(id) {
			student := students[i]
			fmt.Print("| Subject Name: ")
			subjectName := stdInput(reader)
			fmt.Print("| Grade: ")
			grade, _ := strconv.ParseFloat(stdInput(reader), 64)

			newSubject := &Subject{
				Name:  subjectName,
				Grade: grade,
			}

			student.Subjects = append(student.Subjects, *newSubject)
			students[i] = student
			break
		}
	}

	fmt.Println("| Added subject & grade to student.")
}

func addStudentCmd(reader *bufio.Reader) {
	fmt.Println("== Adding a New Student ==")
	fmt.Print("| Forename: ")
	forename := stdInput(reader)
	fmt.Print("| Surname: ")
	surname := stdInput(reader)

	createdStudent := createStudent(forename, surname, idGenerator())
	students = append(students, createdStudent)
}

func seeStudentCmd(reader *bufio.Reader) {
	fmt.Println("== See a Student's Transcript ==")
	fmt.Print("| Student ID: ")
	id, _ := strconv.ParseInt(stdInput(reader), 10, 64)
	for i := 0; i < len(students); i++ {
		if students[i].Id == int(id) {
			student := students[i]
			fmt.Println("--- " + student.Forename + " " + student.Surname + " ---")
			fmt.Println("| Student ID: " + strconv.Itoa(student.Id))
			fmt.Println("| Student Name: " + student.Forename + " " + student.Surname)
			fmt.Println("| Subject(s) Enrolled: ")
			for j := 0; j < len(student.Subjects); j++ {
				subject := student.Subjects[j]
				fmt.Println("|| " + subject.Name + ": " + strconv.FormatFloat(subject.Grade, 'f', 0, 64))
			}
			fmt.Println("-----------------")
		}
	}
}
