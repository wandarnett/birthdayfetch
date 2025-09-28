package data

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type BirthdayData struct {
	Year        int
	Birthday    string
	CurrentDate string
	DaysUntil   int
	GiftIdea    string
	IsBirthday  bool
	Theme       string
}

func GetBirthdayDate(currentTime time.Time) time.Time {
	var UserBirthdayDay int
	var UserBirthdayMonth int
	var UserBirthdayYear int

	fmt.Print("__ Write your birthday Day(dd): ")
	fmt.Scan(&UserBirthdayDay)

	fmt.Print("__ Write your birthday Month(mm): ")
	fmt.Scan(&UserBirthdayMonth)

	fmt.Print("__ Write your birthday Year(yyyy): ")
	fmt.Scan(&UserBirthdayYear)

	CurrentYear := currentTime.Year()

	today := time.Date(CurrentYear, currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local)
	BirthdayDate := time.Date(CurrentYear, time.Month(UserBirthdayMonth), UserBirthdayDay, 0, 0, 0, 0, time.Local)

	if today.After(BirthdayDate) {
		BirthdayDate = time.Date(CurrentYear+1, time.Month(UserBirthdayMonth), UserBirthdayDay, 0, 0, 0, 0, time.Local)
	}

	return BirthdayDate
}

func DaysUntilBirthday(currentTime time.Time, BirthdayTime time.Time) int {
	days := BirthdayTime.Sub(currentTime).Hours() / 24

	daysInt := math.Floor(days)

	if daysInt == -1 {
		daysInt = 0
	}

	return int(daysInt)
}

func GiftIdea() string {
	giftsList := []string{}
	giftsList = append(giftsList,
		"a laptop",
		"a computer",
		"a headphones",
		"a land",
		"a car",
		"a computer mouse",
		"a powerbank",
		"a phone",
		"a watch",
		"the Red Hat Enterprise Linux",
		"a dog",
		"a cat",
		"a house",
		"a horse",
		"a toy",
		"a mug",
		"a cap",
		"a T-shirt",
		"a phone case",
		"a key ring",
		"a funny socks",
		"a puzzle",
		"a postcard with money",
		"a bicycle",
		"a monowheel",
		"a skateboard",
		"a software license",
		"a service subscription",
		"a pad",
	)

	return giftsList[rand.Intn(len(giftsList))]
}
