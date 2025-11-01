package data

import "time"

type Zodiac struct {
	Name       string
	StartMonth int
	StartDay   int
	EndMonth   int
	EndDay     int
}

func GetZodiac(birth time.Time) string {
	zodiacs := []Zodiac{
		{"Aries", 3, 21, 4, 19},
		{"Taurus", 4, 20, 5, 20},
		{"Gemini", 5, 21, 6, 20},
		{"Cancer", 6, 21, 7, 22},
		{"Leo", 7, 23, 8, 22},
		{"Virgo", 8, 23, 9, 22},
		{"Libra", 9, 23, 10, 22},
		{"Scorpio", 10, 23, 11, 21},
		{"Sagittarius", 11, 22, 12, 21},
		{"Capricorn", 12, 22, 1, 19},
		{"Aquarius", 1, 20, 2, 18},
		{"Pisces", 2, 19, 3, 20},
	}

	day := birth.Day()
	month := int(birth.Month())

	for _, z := range zodiacs {
		if (month == z.StartMonth && day >= z.StartDay) ||
			(month == z.EndMonth && day <= z.EndDay) ||
			(z.StartMonth > z.EndMonth && ((month == z.StartMonth && day >= z.StartDay) || (month == z.EndMonth && day <= z.EndDay))) {
			return z.Name
		}
	}

	return "Unknown"
}
