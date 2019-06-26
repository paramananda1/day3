package main

import (
	"fmt"
	"io/ioutil"
	//"log"
	"os"
)


func DoFileOperation(AllStudentsRecord StudentS){
	WriteStudentDateToFile(AllStudentsRecord)
	ReadFromStudentFile()
}



func WriteStudentDateToFile(allStudentsRecord StudentS) {

	emptyFile, err := os.Create("/tmp/allStudentsRecord.txt")
	check(err)
	//log.Println(emptyFile)
	defer emptyFile.Close()

	_, err = fmt.Fprint(emptyFile, allStudentsRecord)

	check(err)
	fmt.Println("Data written to file...")

}


func check(e error) {
	if e != nil {
		panic(e)
	}
}


func ReadFromStudentFile(){
	dat, err := ioutil.ReadFile("/tmp/allStudentsRecord.txt")
	check(err)
	fmt.Println(string(dat))
}