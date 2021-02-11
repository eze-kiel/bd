package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func (bd birthdays) getNames(pfx string) []string {
	var namesInList []string
	var namesToComplete []string

	for _, n := range bd.Birthdays {
		namesInList = append(namesInList, n.Name)
	}

	for _, n := range namesInList {
		if strings.HasPrefix(n, pfx) {
			namesToComplete = append(namesToComplete, n)
		}
	}
	return namesToComplete
}

func (bd *birthdays) removeFromDatabase(i int) {
	bd.Birthdays[i] = bd.Birthdays[len(bd.Birthdays)-1]
	bd.Birthdays = bd.Birthdays[:len(bd.Birthdays)-1]
}

func (bd birthdays) updateDatabase(file string) {
	data, err := json.Marshal(bd)
	if err != nil {
		logrus.Fatalf("cannot write to JSON file: %s", err)
	}

	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		logrus.Fatalf("cannot open %s: %s", file, err)
	}
	defer f.Close()

	err = f.Truncate(0)
	_, err = f.Write(data)
	if err != nil {
		logrus.Fatalf("cannot write new entry to %s: %s", file, err)
	}
}

func getRemainingDays(date string) (string, error) {
	dob, err := time.Parse(dateLayout, date)
	if err != nil {
		return "", err
	}
	curr := time.Now().Format(dateLayout)
	t2, err := time.Parse(dateLayout, curr)
	if err != nil {
		return "", err
	}

	var t1 time.Time
	t1 = dob.AddDate(t2.Year()-dob.Year(), 0, 0)
	days := t1.Sub(t2).Hours() / 24
	if days < 0 {
		days = 365 + days
	}

	return fmt.Sprintf("%0.f", days), nil
}
