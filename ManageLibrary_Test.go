package main


import (
	"fmt"
	"testing"
	"math/rand"

)


func createTestRecord(){
	AllStudentsRecord = make(StudentS, 25)
	for i := 0; i < 25; i++ {
		AllStudentsRecord = append(AllStudentsRecord, HadCodeStudentRecord(&TotalStudentCount))
	}
}


func TestManagaLibrary(t *testing.T) {
	createTestRecord()
	index := rand.Intn(24)
	shoWAllBooks(index)
	addBooks(index)

	fmt.Println("Test Manage Library")
}

