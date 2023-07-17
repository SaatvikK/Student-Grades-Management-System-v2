package cliver

import (
	"bufio"
	"fmt"
	"os"

	//"os"
	//"reflect"
	"strconv"
	//"strings"
	//"io"
	//"log"
	//"net/http"
)

func createDBIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir)
	}
	return nil
}

func createDBFolders(path string) error {
	path = "./database" + path
	status := createDBIfNotExists(path)
	if status != nil {
		return status
	}
	return nil
}

func saveStudent()
