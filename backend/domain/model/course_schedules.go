package model

type CourseSchedule struct {
	DayOfWeek int `db:"day_of_week"`
	Period    int `db:"period"`
}
