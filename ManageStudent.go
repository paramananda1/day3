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
		//newStudent.BestFriends[i+1] = &AllStudentsRecord[bestFreiendRollNo]

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

	AllStudentsRecord[len(AllStudentsRecord)-1].Subjects=nil
	AllStudentsRecord[len(AllStudentsRecord)-1].RollNo=0
	for key := range AllStudentsRecord[len(AllStudentsRecord)-1].BestFriends {
		delete(AllStudentsRecord[index].BestFriends, key)
	}
	removeFromBestFriendMap(AllStudentsRecord[index].RollNo)
	AllStudentsRecord[len(AllStudentsRecord)-1].BestFriends=nil
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


func UpdateBestFriend(index int) {

	fmt.Print("\nEnter 1. Map best frinds 2. To Un map Best friends (Enter to NoN) :")
	var choice int
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		addBestFriend(index)
	case 2:
		deleteBestFriend(index)
	default:
		fmt.Println("You have not Spesified any Choice")
		return
	}
	fmt.Println("New Value Updated")
	showRecordValue(index)
	WriteStudentDateToFile(AllStudentsRecord)
}

func addBestFriend(index int) {
	fmt.Println(" All Best Friends Mapped :  ")
	showAllBfRecord(index )
	fmt.Print(" Enter Number of Best Friends you want to Add ( Hit Enter for NiLL) :  ")
	var noOfBestFriends int

	fmt.Scanf("%d",&noOfBestFriends)
	for i:=0; i<noOfBestFriends; i++{
		var bestFreiendRollNo uint64
		fmt.Printf(" Enter (%d ) Best Friend Roll Number : ",i+1)
		fmt.Scanf("%d",&bestFreiendRollNo)
		present,bfindex  := IsRollNumberExist(bestFreiendRollNo)
		if present != 0{
			fmt.Println(bfindex,index,len(AllStudentsRecord[index].BestFriends),bestFreiendRollNo)
			AllStudentsRecord[index].BestFriends[len(AllStudentsRecord[index].BestFriends)] = &AllStudentsRecord[bfindex]
		}
	}
	fmt.Println(" All Best Friends Mapped after update:  ")
	showAllBfRecord(index )
}


func deleteBestFriend(index int) {
	var bestFreiendRollNo uint64
	fmt.Println(" All Mapped Best Friends record :  ")
	showAllBfRecord(index )
	fmt.Printf(" Enter Best Friend Roll Number to unmap: ")
	fmt.Scanf("%d",&bestFreiendRollNo)
	present,bfindex :=  IsBestFriendMapped(index,bestFreiendRollNo)

	if present == 0{
		return
	}

	delete(AllStudentsRecord[index].BestFriends, bfindex)
	fmt.Println(" All Best Friends Mapped after Delete:  ",)
	showAllBfRecord(index )
}


func  IsRollNumberExist(rollnumber uint64) (present,index int) {
	present = 0
	for key:= range AllStudentsRecord{
		if AllStudentsRecord[key].RollNo == rollnumber{
			present = 1
			index = key
			return
		}
	}
	if present == 0{
		fmt.Println("Student Roll number not Exist")

	}

	return
}

func  IsBestFriendMapped(index int ,rollnumber uint64) (present,bfindex int) {
	present = 0
	for key := range AllStudentsRecord[index].BestFriends {
		if AllStudentsRecord[index].BestFriends[key].RollNo == rollnumber{
			present = 1
			bfindex = key
			return
		}
	}
	if present == 0{
		fmt.Println("Student is not mapped to ",index," as best friend ")

	}

	return
}

func  removeFromBestFriendMap( rollnumber uint64)  {

	for key := range AllStudentsRecord {
		for bfkey := range AllStudentsRecord[key].BestFriends {
			if AllStudentsRecord[key].BestFriends[bfkey].RollNo == rollnumber{
				delete(AllStudentsRecord[key].BestFriends, bfkey)
				//AllStudentsRecord[key].BestFriends[bfkey]=nil
			}
		}
	}
}


func showAllBfRecord(index int){
	fmt.Println("Best Friend Record: ")
	for _,value := range AllStudentsRecord[index].BestFriends {
		fmt.Printf("%+v ",value )
		fmt.Println()
	}
}

func showRecordValue(index int){
	fmt.Print("Record: ")
	fmt.Printf("%+v",AllStudentsRecord[index])
	fmt.Println()

}