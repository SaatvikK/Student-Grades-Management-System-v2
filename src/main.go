package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	//"io"
	//"log"
	//"net/http"
)

type Subject struct {
	Code  string
	Name  string
	Grade float32
}

type Student struct {
	Id       int
	Forename string
	Surname  string
	Subjects []Subject
}

var students = []Student{}

// After the database is created, this function will increment the highest ID in the database.
// However, for now it will just assign 1 as the ID.
func idGenerator() int {
	return 1
}

func stdInput(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}

func createStudent(forename string, surname string, id int) Student {
	newStudent := Student{Id: id, Forename: forename, Surname: surname}
	return newStudent
}

func main() {
	fmt.Println("Welcome to the SGMS v1!")
	fmt.Println("Please type 'help' to get a list of commands!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reflect.TypeOf(reader))

	for {
		fmt.Print("$ ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("help", text) == 0 {
			fmt.Println("hai")
		} else if strings.Compare("addstudent", text) == 0 {
			fmt.Println("== Adding a New Student ==")
			fmt.Print("| Forename: ")
			forename := stdInput(reader)
			fmt.Print("| Surname: ")
			surname := stdInput(reader)

			createdStudent := createStudent(forename, surname, idGenerator())
			students = append(students, createdStudent)

		} else if strings.Compare("addgrade", text) == 0 {
			fmt.Println("== Adding a Grade to a Student ==")
			fmt.Print("| Student ID: ")
			id, _ := strconv.ParseInt(stdInput(reader), 10, 64)
			student := Student{}
			for i := 0; i < len(students); i++ {
				if students[i].Id == int(id) {
					student = students[i]
					break
				}

			}

			fmt.Print("| Subject: ")
		}

	}

}
