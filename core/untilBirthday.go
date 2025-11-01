package core

import (
	"fmt"
	"time"
)

func DaysUntilBirthday(birthdayDate time.Time, currentTime time.Time) {
	nextBirthday := time.Date(currentTime.Year(), birthdayDate.Month(), birthdayDate.Day(), 0, 0, 0, 0, time.Local)

	if nextBirthday.Before(currentTime) {
		nextBirthday = nextBirthday.AddDate(1, 0, 0)
	}

	days := nextBirthday.Sub(currentTime).Hours() / 24
	fmt.Println(days)
}

func BirthdayWeekday(birthdayDate time.Time, currentTime time.Time) {
	nextBirthday := time.Date(
		currentTime.Year(),
		birthdayDate.Month(),
		birthdayDate.Day(),
		0, 0, 0, 0,
		time.Local,
	)

	if nextBirthday.Before(currentTime) {
		nextBirthday = nextBirthday.AddDate(1, 0, 0)
	}

	weekday := nextBirthday.Weekday()
	fmt.Println(weekday)
}

func NextUserAge(birthdayDate time.Time, currentTime time.Time) {
	nextBirthday := time.Date(
		currentTime.Year(),
		birthdayDate.Month(),
		birthdayDate.Day(),
		0, 0, 0, 0,
		time.Local,
	)

	if nextBirthday.Before(currentTime) {
		nextBirthday = nextBirthday.AddDate(1, 0, 0)
	}
	birthdayYear := birthdayDate.Year()

	userAge := nextBirthday.Year() - birthdayYear

	fmt.Println(userAge)
}
