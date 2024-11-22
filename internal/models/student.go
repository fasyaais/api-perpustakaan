package models

type Student struct {
	Name  string `json:"name"`
	NPM   string `json:"npm"`
	PRODI string `json:"prodi"`
}
type StudentImpl struct {
	Student
}
