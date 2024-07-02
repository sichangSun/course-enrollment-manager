package model

// Course ...
type Course struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Semester    int    `db:"semester"`
	Instructor  int    `db:"instructor"`
	Credits     int    `db:"credits"`
	Description string `db:"description"`
}

// Course_schedules
type CourseScheduleDetail struct {
	ID          int    `db:"id"`
	CourseName  string `db:"course_name"`
	Semester    int    `db:"semester"`
	Instructor  int    `db:"instructor"`
	Credits     int    `db:"credits"`
	TeacherName string `db:"teacher_name"`
	Position    string `db:"position"`
	DayOfWeek   int    `db:"day_of_week"`
	Period      int    `db:"period"`
	Description string `db:"description"`
}

// Including associated teachers and mulit schedules.
type CourseWithMulitSchedules struct {
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
