package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"
)

type DayRelation int

const (
	Past DayRelation = iota
	Today
	Future
)

type DaysUntil struct {
	days     int
	date     time.Time
	relation DayRelation
}

func printSubCommands() {
	fmt.Println("Available subcommands:")
	fmt.Println("  days-until - Calculate days until a specific date")
	fmt.Println("  age        - Calculate age based on birthdate")
}

func convertInputToTime(dateString string) (time.Time, error) {
	const form = "2006-01-02"
	dateTime, err := time.Parse(form, dateString)
	return dateTime, err
}

func parseUserInputDate(date string) error {
	const form = "2006-01-02"
	_, err := time.Parse(form, date)
	if err != nil {
		return fmt.Errorf("invalid date format: %s. Please use YYYY-MM-DD", date)
	}
	return nil
}

// It is also an option to merge convertInputToTime and parseUserInputDate
// func convertInputToTime(date string) (time.Time, error) {
//     const form = "2006-01-02"
//     t, err := time.Parse(form, date)
//     if err != nil {
//         return t, fmt.Errorf("invalid date format: %s. Please use YYYY-MM-DD", date)
//     }
//     return t, nil
// }

func getDaysUntil(inputDate time.Time) DaysUntil {
	// Default from to todays Datetime when calculating daysuntil
	today := time.Now()

	// Calculate the days between today and input
	days := daysBetween(today, inputDate)

	var relation DayRelation
	switch {
	case days < 0:
		// Convert to positive integer
		days = -days
		relation = Past
	case days == 0:
		relation = Today
	default:
		relation = Future
	}

	return DaysUntil{days: days, date: inputDate, relation: relation}
}

func daysBetween(from, to time.Time) int {
	days := int(math.Floor(from.Sub(to).Hours() / 24))
	return days
}

func printDaysUntil(result DaysUntil) {
	//Format and print result
	switch {
	case result.relation == Today:
		fmt.Println("0 days, it is today!")
	case result.relation == Past:
		fmt.Printf("This date %s has already been. It was %d days ago", result.date.Format("2006-01-02"), result.days)
	default:
		fmt.Printf("There are %d days until %s \n", result.days, result.date.Format("2006-01-02"))
	}

}

func main() {
	start := time.Now()
	fmt.Println(start)

	daysUntilCmd := flag.NewFlagSet("days-until", flag.ExitOnError)
	ageCmd := flag.NewFlagSet("age", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected an argument subcommands, see the list of available subcommands below:")
		printSubCommands()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "days-until":
		daysUntilCmd.Parse(os.Args[2:])
		input := daysUntilCmd.Arg(0)

		if err := parseUserInputDate(input); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		inputDate, err := convertInputToTime(input)
		if err != nil {
			panic(err)
		}

		result := getDaysUntil(inputDate)
		printDaysUntil(result)

	case "age":
		ageCmd.Parse(os.Args[2:])
		fmt.Println("age command")
	}

}
