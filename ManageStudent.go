package main

import (
	"bufio"
	"fmt"
	"os"
)

func CreateStudentRecord(TotalStudentCount *uint64) Student {

	rollnum := *TotalStudentCount+1
	fmt.Println("You are about to create a student record for Roll Number",rollnum)
	newStudent := new(Student)
	newStudent.RollNo = rollnum

	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter FirstName  : ")
	newStudent.FirstName, _ = consoleReader.ReadString('\n')
	fmt.Print("Enter LastName : ")
	newStudent.LastName, _ = consoleReader.ReadString('\n')
	//fmt.Print(newStudent.FirstName, newStudent.LastName)
	fmt.Print("Enter 	Gender :")
	fmt.Scanf("%s",&newStudent.Gender)
	fmt.Print("Enter 	Stream :")
	fmt.Scanf("%s",&newStudent.Stream)

	fmt.Println(" Enter Subjects :  ")
	var noOfSubjects int
	fmt.Print(" How many Subjects you want to enter ( Hit Enter for NiLL):  ")
	fmt.Scanf("%d",&noOfSubjects)
	newStudent.Subjects = make([]string,noOfSubjects)
	for i:=0; i<noOfSubjects; i++{
		fmt.Printf("Sub %d : ",i+1)
		newStudent.Subjects[i],_= consoleReader.ReadString('\n')
	}

	fmt.Println(" Enter Books :  ")
	var noOfBooks int
	fmt.Print(" How many Books you want to Issue ( Hit Enter for NiLL):  ")
	fmt.Scanf("%d",&noOfBooks)
	newStudent.Books = make([]string,noOfBooks)
	for i:=0; i<noOfBooks; i++{
		fmt.Printf("Book %d : ",i+1)
		newStudent.Books[i],_ = consoleReader.ReadString('\n')
	}

	fmt.Print(" Enter Number of Best Friends ( Hit Enter for NiLL) :  ")
	var noOfBestFriends int

	fmt.Scanf("%d",&noOfBestFriends)
	for i:=0; i<noOfBestFriends; i++{
		var bestFreiendRollNo uint64
		fmt.Printf(" Enter (%d ) Best Friend Roll Number : ",i+1)
		fmt.Scanf("%d",&bestFreiendRollNo)
		if(bestFreiendRollNo > rollnum ){
			fmt.Println("Roll number ",bestFreiendRollNo, " does not Exist..Do you want to continue (y/n)?")
			var conti string
			fmt.Scanf("%s",&conti)
			if conti != "Y" || conti != "y" {
				break
			}
			i--  // re- try for old
			continue
		}
		// if Roll number exist Map to student record
		newStudent.BestFriends = make(map[int]*Student)
		newStudent.BestFriends[i+1] = &AllStudentsRecord[bestFreiendRollNo]

	}
	*TotalStudentCount = rollnum
	return *newStudent
}


func DeleteStudentRecord(index int)  {

	// Remove the element at index bookKey from Books[].
	// Shift [index+1:] left one index.
	copy(AllStudentsRecord[index:],AllStudentsRecord[index+1:])
	// Erase last element (write zero value).
	AllStudentsRecord[len(AllStudentsRecord)-1].FirstName=""
	AllStudentsRecord[len(AllStudentsRecord)-1].LastName=""
	AllStudentsRecord[len(AllStudentsRecord)-1].Gender=""
	AllStudentsRecord[len(AllStudentsRecord)-1].Stream=""
	AllStudentsRecord[len(AllStudentsRecord)-1].Books=nil
	AllStudentsRecord[len(AllStudentsRecord)-1].BestFriends=nil
	AllStudentsRecord[len(AllStudentsRecord)-1].Subjects=nil
	AllStudentsRecord[len(AllStudentsRecord)-1].RollNo=0
	//AllStudentsRecord[len(AllStudentsRecord)-1] = nil
	// Truncate slice.
	AllStudentsRecord = AllStudentsRecord[:len(AllStudentsRecord)-1]
	WriteStudentDateToFile(AllStudentsRecord)

}

func UpdateStudentInfo(index int) {
	showRecordValue(index)
	fmt.Print("\nEnter field to Update. 1. FirstName 2. LastName 3. Gender 4. Stream (Enter to NoN) :")
	var choice int
	var newValue string
	fmt.Scanf("%d",&choice)
	fmt.Print("Enter New Value :")
	fmt.Scanf("%s",&newValue)
	switch choice {
	case 1:
		AllStudentsRecord[index].FirstName = newValue
	case 2:
		AllStudentsRecord[index].LastName = newValue
	case 3:
		AllStudentsRecord[index].Gender = newValue
	case 4:
		AllStudentsRecord[index].Stream = newValue
	default:
		fmt.Println("You have not Spesified any Choice")
		return
	}
	fmt.Println("New Value Updated")
	showRecordValue(index)
	WriteStudentDateToFile(AllStudentsRecord)
}


func  IsRollNumberExist(rollnumber uint64) (present,index int) {
	present = 0
	for key:= range AllStudentsRecord{
		if AllStudentsRecord[key].RollNo == rollnumber{
			present = 1
			index = key
			break
		}
	}
	if present == 0{
		fmt.Println("Student Roll number not Exist")

	}

	return
}

func showRecordValue(index int){
	fmt.Print("Record: ")
	fmt.Printf("%+v",AllStudentsRecord[index])
	fmt.Println()

}