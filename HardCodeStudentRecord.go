package main

import (
	"strconv"
)

func HadCodeStudentRecord(TotalStudentCount *uint64) Student {

	rollnum := *TotalStudentCount+1

	newStudent := new(Student)
	newStudent.RollNo = rollnum

	newStudent.FirstName = "Paramananda" + strconv.Itoa(int(rollnum))
	newStudent.LastName = "Pati" + strconv.Itoa(int(rollnum))

	newStudent.Gender = "M" + strconv.Itoa(int(rollnum))
	newStudent.Stream ="BTech" + strconv.Itoa(int(rollnum))


	newStudent.Subjects = make([]string,2)
	newStudent.Subjects[0] = "C" + strconv.Itoa(int(rollnum))
	newStudent.Subjects[1] = "C++" + strconv.Itoa(int(rollnum))
	newStudent.Subjects[1] = "Java" + strconv.Itoa(int(rollnum))
	newStudent.Subjects[1] = "Perl" + strconv.Itoa(int(rollnum))
	newStudent.Subjects[1] = "OS" + strconv.Itoa(int(rollnum))

	newStudent.Books = make([]string,2)
	newStudent.Books[0] = "Book-1" + strconv.Itoa(int(rollnum))
	newStudent.Books[1] = "Book-2" + strconv.Itoa(int(rollnum))
	newStudent.Books[1] = "OS-2" + strconv.Itoa(int(rollnum))
	newStudent.Books[1] = "JAVA-2" + strconv.Itoa(int(rollnum))
	newStudent.Books[1] = "C-2" + strconv.Itoa(int(rollnum))

	// if Roll number exist Map to student record
	newStudent.BestFriends = make(map[int]*Student)
	//newStudent.BestFriends[0] = &AllStudentsRecord[rollnum]

	*TotalStudentCount = rollnum
	return *newStudent
}
