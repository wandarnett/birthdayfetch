package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Welcome() (time.Time, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("__________  Welcome to birthdayfetch  __________")
	fmt.Println("")
	for {
		fmt.Print("Enter your age (dd.mm.yyyy): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if date, ok := isValidDate(input); ok {
			fmt.Println("Date accepted:", date.Format("02.01.2006"))
			return date, nil
		}
		fmt.Println("Error: invalid format or non-existent date. Please try again.")
	}

}

func isValidDate(s string) (time.Time, bool) {
	if len(s) != 10 || s[2] != '.' || s[5] != '.' {
		return time.Time{}, false
	}
	t, err := time.Parse("02.01.2006", s)
	if err != nil || t.After(time.Now()) {
		return time.Time{}, false
	}
	return t, true
}
