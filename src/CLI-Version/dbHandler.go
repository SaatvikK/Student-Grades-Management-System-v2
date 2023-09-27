package cliver

import (
	//"bufio"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"os"
	"strconv"
	//"os"
	//"reflect"
	//"strings"
	//"io"
	//"log"
	//"net/http"
)

var dbPath string = "./database"

func createDBIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir)
	}
	return nil
}

func createDBFolders(path string) error {
	path = dbPath + path
	status := createDBIfNotExists(path)
	if status != nil {
		return status
	}
	return nil
}

func saveStudent(student Student, collectionPath string) error {
	studentJson, err := json.Marshal(student)
	if err == nil {
		return err
	}

	studentID := strconv.Itoa(student.Id)
	err = ioutil.WriteFile(dbPath+collectionPath+studentID+".json", studentJson, 0644)

	if err == nil {
		return err
	}

	return nil
}

func loadStudent(studentID int, collectionPath string) interface{} {
	var student Student
	byteValue, err := ioutil.ReadFile(dbPath + collectionPath + strconv.Itoa(studentID) + ".json")
	if err == nil {
		return err
	}

	json.Unmarshal(byteValue, &student)
	return student
}

func createSchoolInDB(schoolID string, name string, noYearGroups int) error {
	status := createDBFolders("/" + schoolID)
	if status != nil {
		return status
	}

	schoolPath := "/" + schoolID
	collections := [6]string{"years", "teachers", "SchoolData", "MetaData", "ValidTeachers", "ValidAdmins"}

	for i := 0; i < len(collections); i++ {
		if collections[i] == "years" {
			for j := 0; j < noYearGroups+1; j++ {
				status = createDBFolders(schoolPath + "/year" + strconv.Itoa(j+1))
				if status != nil {
					return status
				}
			}
		}

		status = createDBFolders(schoolPath + "/" + collections[i])
		if status != nil {
			return status
		}
	}
	return nil
}
