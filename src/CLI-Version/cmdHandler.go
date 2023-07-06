package cliver

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	//"strconv"
	"strings"
	//"io"
	//"log"
	//"net/http"
)

type Subject struct {
	Name  string
	Grade float64
}

type Student struct {
	Id       int
	Forename string
	Surname  string
	Subjects []Subject
}

var students = []Student{} // temp cache database

// After the database is created, this function will increment the highest ID in the database.
// However, for now it will just assign 1 as the ID.
func idGenerator() int {
	return 1
}

func stdInput(reader *bufio.Reader) string {
	cmd, _ := reader.ReadString('\n')
	return strings.Replace(cmd, "\n", "", -1)
}

func cmdHandler(cmd string, reader *bufio.Reader) {
	if strings.Compare("help", cmd) == 0 {
		fmt.Println("hai")

	} else if strings.Compare("addstudent", cmd) == 0 {
		addStudentCmd(reader)

	} else if strings.Compare("addgrade", cmd) == 0 {
		addGradeCmd(reader)

	} else if strings.Compare("seestudent", cmd) == 0 {
		seeStudentCmd(reader)

	} else if strings.Compare("exit", cmd) == 0 {
		os.Exit(1)
	}
}

func CmdInput() {
	fmt.Println("Welcome to the SGMS v1!")
	fmt.Println("Please type 'help' to get a list of commands!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reflect.TypeOf(reader))

	for {
		fmt.Print("$ ")
		cmd, _ := reader.ReadString('\n')
		// convert CRLF to LF
		cmd = strings.Replace(cmd, "\n", "", -1)
		cmdHandler(cmd, reader)
	}

}
