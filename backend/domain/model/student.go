package model

import "golang.org/x/crypto/bcrypt"

// student ...
type Student struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

// StudentCoursesDetail ...
type StudentCoursesDetail struct {
	StudentID   int    `db:"student_id"`
	CourseID    int    `db:"course_id"`
	CourseName  string `db:"course_name"`
	Semester    int    `db:"semester"`
	Instructor  int    `db:"instructor"`
	Credits     int    `db:"credits"`
	TeacherName string `db:"teacher_name"`
	Position    string `db:"position"`
	DayOfWeek   int    `db:"day_of_week"`
	Period      int    `db:"period"`
}

// StudentCoursesWithMulitSchedules... Including associated teachers and mulit schedules.
type StudentCoursesWithMulitSchedules struct {
	StudentID    int
	CourseID     int
	CourseName   string
	Semester     int
	InstructorID int
	Credits      int
	TeacherName  string
	Position     string
	Description  string
	Schedules    []*CourseSchedule
}

// PasswordEncrypt ...
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CompareHashAndPassword ...
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
