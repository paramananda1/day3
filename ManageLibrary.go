package main

import (
	"bufio"
	"fmt"
	"os"
)

func ManagaLibrary(index int) {
	fmt.Print("Spesify the Oparetion you want to do. 1. Add new books 2. Delete Books 3. Show All Books (Enter to NoN) :")
	var choice int
	fmt.Scanf("%d",&choice)
	switch choice {
	case 1:
		shoWAllBooks(index)
		addBooks(index)
	case 2:
		shoWAllBooks(index)
		deleteBooks(index)
	case 3:
		shoWAllBooks(index)
	default:
		fmt.Println("You have not Spesified any Choice")
	}

}

func addBooks(index int){
	consoleReader := bufio.NewReader(os.Stdin)
	reTry := "Y"

				for reTry=="Y" || reTry=="y" {
					fmt.Print("Enter Book name you want to Add:")
					newBook, _ := consoleReader.ReadString('\n')
					AllStudentsRecord[index].Books = append(AllStudentsRecord[index].Books, newBook)
					reTry="n"
					fmt.Print("Do you want to Add more book (y/n)? :")
					fmt.Scanf("%s",&reTry)
				}

		if reTry != "Y" || reTry != "y"{
			WriteStudentDateToFile(AllStudentsRecord)
		}


}

func deleteBooks(index int){

	consoleReader := bufio.NewReader(os.Stdin)
	reTry := "Y"

			for reTry=="Y" || reTry=="y" {
				reTry="n"
				fmt.Print("Enter Book name you want to Delete:")
				oldBook, _ := consoleReader.ReadString('\n')
				present,bookindex := isBookEnrolled(index, oldBook)
				if present == 1{
					// Remove the element at index bookKey from Books[].
					// Shift [bookKey+1:] left one index.
					copy(AllStudentsRecord[index].Books[bookindex:],AllStudentsRecord[index].Books[bookindex+1:])
					// Erase last element (write zero value).
					AllStudentsRecord[index].Books[len(AllStudentsRecord[index].Books)-1] = ""
					// Truncate slice.
					AllStudentsRecord[index].Books = AllStudentsRecord[index].Books[:len(AllStudentsRecord[index].Books)-1]

				}

				fmt.Print("Do you want to Delete more book (y/n)? :")
				fmt.Scanf("%s",&reTry)
			}

			if reTry != "Y" || reTry != "y"{
				WriteStudentDateToFile(AllStudentsRecord)
			}



}

func  isBookEnrolled(studentindex int, oldBook string ) (present,index int) {
	present = 0
	for bookKey:= range AllStudentsRecord[studentindex].Books {
		if AllStudentsRecord[studentindex].Books[bookKey] == oldBook{
			present = 1
			index = bookKey
			break
		}
	}
	if present == 0{
		fmt.Println("No Books EnRolled in this name for Student .")

	}

	return
}

func shoWAllBooks(index int){
			fmt.Print("Books in Record: ")
			fmt.Println(AllStudentsRecord[index].Books)

}