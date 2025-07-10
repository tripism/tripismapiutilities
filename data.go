package tripismapiutilities

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// CleanWebsiteQueryString removes unused / tracking query string parameters from website URLs
//
//	for example, it deletes all "utm_source", "utm_campaign" and "utm_medium" Google Analytcs parameters
func CleanWebsiteQueryString(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return uri
	}

	q := u.Query()
	q.Del("scid")
	q.Del("cid")
	q.Del("y_source")
	q.Del("utm_source")
	q.Del("utm_campaign")
	q.Del("utm_medium")
	q.Del("utm_term")
	q.Del("WT.mc_id")
	q.Del("SEO_id")
	q.Del("s_cid")
	q.Del("cm_mmc")
	q.Del("sourceid")

	u.RawQuery = q.Encode()
	return u.String()
}

type NormalizedEntry struct {
	Original string // Original string before normalization
	Type     string // "time", "period", "empty" or "unknown"
	Value    string // Normalized value like "6:00 pm" or "24 hours"
}

// NormalizeCancellationPolicy normalizes various formats of cancellation policies
// to a standard format for easier comparison and storage.
//
// It handles time formats and periods, such as "6 PM", "12:00 noon", "24 hours", etc.
// It returns a NormalizedEntry with the original string, type, and normalized value.
// It also handles empty strings and standardizes punctuation.
func NormalizeCancellationPolicy(input string) NormalizedEntry {
	original := input
	raw := strings.ToLower(strings.TrimSpace(input))

	if raw == "" {
		return NormalizedEntry{original, "empty", ""}
	}
	// Standardize punctuation
	raw = strings.ReplaceAll(raw, ";", ":")
	raw = strings.ReplaceAll(raw, ".", "")
	raw = strings.TrimSpace(raw)

	// Normalize "noon"
	if raw == "noon" || raw == "12noon" {
		return NormalizedEntry{original, "time", "12:00 pm"}
	}

	// Normalize space between digit and AM/PM e.g. "6 PM" -> "6PM"
	raw = regexp.MustCompile(`(?i)\b(\d{1,2})\s+(am|pm)\b`).ReplaceAllString(raw, `$1$2`)

	// Regexes anchored at start to extract only the main time/period part
	time12re := regexp.MustCompile(`(?i)^(1[0-2]|0?[1-9]):?([0-5][0-9])?(am|pm)`)
	time24re := regexp.MustCompile(`^([01]?[0-9]|2[0-3]):([0-5][0-9])`)
	hourOnlyRe := regexp.MustCompile(`^\d{1,2}`)
	periodRe := regexp.MustCompile(`^(\d+)\s*(hours?|hrs?|hr|h)`)
	militaryRe := regexp.MustCompile(`^([01]?[0-9]|2[0-3])[0-5][0-9]`)

	// Try 12-hour format
	if m := time12re.FindString(raw); m != "" {
		t, err := time.Parse("3:04pm", m)
		if err != nil {
			t, err = time.Parse("3pm", m)
		}
		if err == nil {
			return NormalizedEntry{original, "time", t.Format("3:04 pm")}
		}
	}

	// Try 24-hour format
	if m := time24re.FindString(raw); m != "" {
		parts := time24re.FindStringSubmatch(raw)
		hour := parts[1]
		minute := parts[2]

		t, err := time.Parse("15:04", fmt.Sprintf("%s:%s", hour, minute))
		if err == nil {
			return NormalizedEntry{original, "time", t.Format("3:04 pm")}
		}
	}

	// Handle hour only or period if just a number
	if m := hourOnlyRe.FindString(raw); m != "" {
		hourInt := 0
		fmt.Sscanf(m, "%d", &hourInt)
		if hourInt >= 0 && hourInt <= 23 {
			t, err := time.Parse("15", fmt.Sprintf("%d", hourInt))
			if err == nil {
				return NormalizedEntry{original, "time", t.Format("3:04 pm")}
			}
			return NormalizedEntry{original, "time", fmt.Sprintf("%d:00", hourInt)}
		}
		if hourInt > 23 {
			return NormalizedEntry{original, "period", fmt.Sprintf("%d hours", hourInt)}
		}
	}

	// Try military time
	if m := militaryRe.FindString(raw); m != "" {
		hour := m[:len(m)-2]
		minute := m[len(m)-2:]
		t, err := time.Parse("15:04", fmt.Sprintf("%s:%s", hour, minute))
		if err == nil {
			return NormalizedEntry{original, "time", t.Format("3:04 pm")}
		}
	}

	// Try period (notice)
	if m := periodRe.FindString(raw); m != "" {
		matches := periodRe.FindStringSubmatch(raw)
		return NormalizedEntry{original, "period", matches[1] + " hours"}
	}

	// Fallback unknown
	return NormalizedEntry{original, "unknown", raw}
}
