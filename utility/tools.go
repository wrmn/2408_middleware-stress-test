package utility

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

func getRandomNum(len int, max int) string {
	source := rand.NewSource(uint64(time.Now().UnixNano()))
	randomGenerator := rand.New(source)
	num := randomGenerator.Intn(max)
	return PadLeft(len, '0', strconv.Itoa(num))
}

func PadLeft(x int, y rune, z string) string {
	if len(z) >= x {
		return z
	}

	paddingLength := x - len(z)
	padding := strings.Repeat(string(y), paddingLength)
	return padding + z
}

func timeSince(t time.Time) string {
	duration := time.Since(t)

	switch {
	case duration < time.Minute:
		return fmt.Sprintf("%d seconds ago", int(duration.Seconds()))
	case duration < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	case duration < time.Hour*24:
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	default:
		days := int(duration.Hours() / 24)
		return fmt.Sprintf("%d days ago", days)
	}
}
