package view

import (
	"fmt"
	"time"
)

func FormatDateRange(startDate, endDate string) (string, error) {
	start, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		return "", fmt.Errorf("error parsing start date: %w", err)
	}

	end, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		return "", fmt.Errorf("error parsing end date: %w", err)
	}

	if start.Year() == end.Year() && start.Month() == end.Month() {
		// Same month and year: "Aug 10-12, 2025"
		return fmt.Sprintf("%s %d-%d, %d",
			start.Format("Jan"),
			start.Day(),
			end.Day(),
			start.Year()), nil
	}

	if start.Year() == end.Year() {
		// Same year, different months: "Aug 10-Sep 12, 2025"
		return fmt.Sprintf("%s %d-%s %d, %d",
			start.Format("Jan"),
			start.Day(),
			end.Format("Jan"),
			end.Day(),
			start.Year()), nil
	}

	// Different years: "Aug 10, 2025-Sep 12, 2026"
	return fmt.Sprintf("%s %d, %d-%s %d, %d",
		start.Format("Jan"),
		start.Day(),
		start.Year(),
		end.Format("Jan"),
		end.Day(),
		end.Year()), nil
}

func AvatarURL(avatar string) string {
	if len(avatar) >= 4 && avatar[:4] == "http" {
		return avatar
	}
	return "/static/images/speakers/" + avatar
}
