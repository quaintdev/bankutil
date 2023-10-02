package bankutil

import (
	"fmt"
	"regexp"
	"strings"
)

// extractUPIAddress extracts UPI address from transaction descrpiption
func extractUPIAddress(input string) (string, error) {
	if !strings.Contains(input, "UPI-") {
		return "", fmt.Errorf("not a UPI transaction")
	}

	re := regexp.MustCompile(`-([^@-]+)@([^@-]+)-`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 3 {
		return "", fmt.Errorf("error extracting UPI address")
	}

	result := matches[1] + "@" + matches[2]
	return result, nil
}
