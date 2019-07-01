package main

import (
	"fmt"
)


type Student struct {
	RollNo uint64
	FirstName string
	LastName string
	Gender string
	Stream string
	Subjects []string
	Books []string
	BestFriends map[int]*Student
}

type StudentS []Student
var TotalStudentCount uint64 = 0
var AllStudentsRecord StudentS
// Mutex for AllStudentsRecord and TotalStudentCount
//var StudentMutex = &sync.Mutex{}
func main() {
	var noOfStudents int
	fmt.Print(" How many Student Records you want to Create ? (Enter to Auto create):  ")
	fmt.Scanf("%d",&noOfStudents)
	//StudentMutex.Lock()

	if(noOfStudents < 1) {

		AllStudentsRecord = make(StudentS, 5)
		for i := 0; i < 15; i++ {
			AllStudentsRecord = append(AllStudentsRecord, HadCodeStudentRecord(&TotalStudentCount))
		}
	}else {

		AllStudentsRecord = make(StudentS, noOfStudents)
		for i := 0; i < noOfStudents; i++ {
			AllStudentsRecord = append(AllStudentsRecord, CreateStudentRecord(&TotalStudentCount))
		}
	}
	WriteStudentDateToFile(AllStudentsRecord)

	var reTry string = "Y"
	for reTry == "Y" || reTry == "y" {
		reTry="n"
		var rollnumber uint64
		fmt.Print("Please enter Student Roll Number to Manage record:")
		fmt.Scanf("%d",&rollnumber)
		present,index := IsRollNumberExist(rollnumber)
		if present != 0 && index >= 0 {

			fmt.Print("Please enter Choice. \n 1. Update Student Info \n 2. Delete Student \n 3. Manage Books \n 4. Swow all record.\n 5. Update BestFriend \n Enter n to exist : ")
			var choice int
			fmt.Scanf("%d",&choice)
			switch choice {
			case 1:
				UpdateStudentInfo(index)

			case 2:
				DeleteStudentRecord(index)
			case 3:
				ManagaLibrary(index)
			case 4:
				ReadFromStudentFile()
			case 5:
				UpdateBestFriend(index)
			default:
				fmt.Println("You have not Spesified any Choice")

			}
		}
		fmt.Printf("Do you wan to continue (y/n):")
		fmt.Scanf("%s",&reTry)
	}


	//StudentMutex.Unlock()

}
