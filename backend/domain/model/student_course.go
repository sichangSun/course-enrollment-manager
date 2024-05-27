package model

type StudentCourse struct {
	StudentID int `db:"student_id"`
	CourseID  int `db:"course_id"`
}
