package parsinglogfiles

import "regexp"

var (
	validLineRE       = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	splitLogLineRE    = regexp.MustCompile(`<[~*=-]*>`)
	quotedPasswordRE  = regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	endOfLineTextRE   = regexp.MustCompile(`end-of-line\d+`)
	userNameCaptureRE = regexp.MustCompile(`User\s+([^\s]+)\s`)
)

func IsValidLine(text string) bool {
	return validLineRE.MatchString(text)
}

func SplitLogLine(text string) []string {
	return splitLogLineRE.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if quotedPasswordRE.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return endOfLineTextRE.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	out := make([]string, len(lines))

	for i, line := range lines {
		m := userNameCaptureRE.FindStringSubmatch(line)
		if len(m) == 2 {
			out[i] = "[USR] " + m[1] + " " + line
		} else {
			out[i] = line
		}
	}

	return out
}
