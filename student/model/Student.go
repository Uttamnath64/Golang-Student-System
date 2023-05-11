package model

type Student struct {
	ID     string `json:"Id"`
	Name   string `json:"Student_Name"`
	Course string `json:"Course_Name"`
	Eno    string `json:"Enrollment_Number"`
	DOB    string `json:"DOB"`
	Gender string `json:"Gender"`
	Sem    string `json:"Semeter"`
}
