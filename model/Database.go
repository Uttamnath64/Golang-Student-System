package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const USERNAME = "root"
const PASSWORD = ""
const DBNAME = "bvschool"

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", USERNAME+":"+PASSWORD+"@tcp(localhost)/"+DBNAME)
	if err != nil {
		fmt.Print("Database error!")
	}
	DB = db
}

func GetStudents() []Student {
	result, err := DB.Query("Select * From student")

	if err != nil {
		return nil
	}

	Students := []Student{}

	for result.Next() {
		var stu Student
		err := result.Scan(&stu.ID, &stu.Name, &stu.Course, &stu.Eno, &stu.DOB, &stu.Gender, &stu.Sem)
		if err != nil {
			return nil
		}
		Students = append(Students, stu)
	}
	return Students
}

func GetStudent(id string) *Student {

	stu := &Student{}
	result, err := DB.Query("Select * From student Where Id = ?", id)
	if err != nil {
		return nil
	}

	if result.Next() {

		err := result.Scan(&stu.ID, &stu.Name, &stu.Course, &stu.Eno, &stu.DOB, &stu.Gender, &stu.Sem)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}
	return stu
}

func AddStudent(stu Student) bool {
	insert, err := DB.Query("Insert Into student(Id,Student_Name,Course_Name,Enrollment_Number,DOB,Gender,Semeter) Values (?,?,?,?,?,?,?)",
		stu.ID, stu.Name, stu.Course, stu.Eno, stu.DOB, stu.Gender, stu.Sem)
	if err != nil {
		return false
	}
	defer insert.Close()
	return true
}

func UpdateStudent(stu Student, id string) bool {
	update, err := DB.Query("Update student SET Id = ?,Student_Name = ?,Course_Name = ?,Enrollment_Number = ?,DOB = ?,Gender = ?,Semeter = ? Where Id = ?",
		stu.ID, stu.Name, stu.Course, stu.Eno, stu.DOB, stu.Gender, stu.Sem, id)
	if err != nil {
		return false
	}
	update.Close()
	return true
}

func DeleteStudent(id string) bool {
	delete, err := DB.Query("Delete From student Where Id = '" + id + "'")
	if err != nil {
		return false
	}
	delete.Close()
	return true
}
