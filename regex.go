package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// last modified 2024-01-10.0945

// compile regex
func compileRegex(mode string, customPattern string) (*regexp.Regexp, error) {
	var pattern string

	if customPattern != "" {
		pattern = customPattern
	} else {
		// use pattern returned by getPatternByMode
		pattern = getPatternByMode(mode)
		if pattern == "" {
			return nil, fmt.Errorf("invalid mode: %s", mode)
		}
	}

	// sanity checks
	if err := validateRegexPattern(pattern); err != nil {
		return nil, err
	}

	// compile regex
	compiledRegex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to compile regex: %v", err)
	}
	return compiledRegex, nil
}

// set regex pattern
func getPatternByMode(mode string) string {
	if strings.HasPrefix(mode, "hex") {
		length, _ := strconv.Atoi(strings.TrimPrefix(mode, "hex"))
		return fmt.Sprintf(`(^|\s)([A-Fa-f0-9]{%d})(\s|$|:)`, length)
	}

	// predefined modes
	switch mode {
	case "md5", "ntlm", "0", "1000": // hex32
		return `(^|\s)([A-Fa-f0-9]{32})(\s|$|:)`
	case "sha1", "mysql", "ripemd160", "100", "300", "6000": // hex40
		return `(^|\s)([A-Fa-f0-9]{40})(\s|$|:)`
	case "sha224", "1300": // hex56
		return `(^|\s)([A-Fa-f0-9]{56})(\s|$|:)`
	case "sha256", "1400": // hex64
		return `(^|\s)([A-Fa-f0-9]{64})(\s|$|:)`
	case "sha384", "10800": // hex96
		return `(^|\s)([A-Fa-f0-9]{96})(\s|$|:)`
	case "sha512", "1700": // hex128
		return `(^|\s)([A-Fa-f0-9]{128})(\s|$|:)`
	case "crc32", "11500": // hex8
		return `(^|\s)([A-Fa-f0-9]{8})(\s|$|:)`
	case "crc64", "28000": // hex16
		return `(^|\s)([A-Fa-f0-9]{16})(\s|$|:)`
	case "bcrypt", "3200":
		return `(^|\s)(\$2[a-zA-Z]\$\d{2}\$[\x20-\x7E]{53})(\s|$|:)`
	case "phpass", "phpbb3", "joomla", "wordpress", "400":
		return `(^|\s)(\$[PH]\$[./0-9A-Za-z]{31})(\s|$|:)`
	case "md5crypt", "500":
		return `(^|\s)(\$1\$.{0,8}\$[./0-9A-Za-z]{22})(\s|$|:)`
	case "pbkdf2sha256", "pbkdf2-sha256", "pbkdf2_sha256", "10000":
		return `(^|\s)(pbkdf2_sha256\$\d{1,6}\$[\x20-\x7E]{57})(\s|$|:)`
	case "sha512crypt", "sha512-crypt", "sha512_crypt", "1800":
		return `(^|\s)(\$6\$.{0,16}\$[./0-9A-Za-z]{86})(\s|$|:)`
	case "bitcoin", "11300":
		return `(^|\s)\$bitcoin\$[0-9]{1,3}\$[a-f0-9]{40,}\$[0-9]{2}\$[a-f0-9]{2,}\$\d{2,}\$\d{1,}\$\d{1,}\$\d{1,}\$\d{1,}(\s|$|:)`
	case "metamask", "26600":
		return `(^|\s)\$metamask\$[\x20-\x7E]{3,}\$[\x20-\x7E]{3,}\$[\x20-\x7E]{3,}(\s|$|:|=)`
	default:
		return ""
	}
}

// santiy check for custom regex
func validateRegexPattern(pattern string) error {
	if len(pattern) == 0 {
		return fmt.Errorf("empty pattern is not valid")
	}
	if !hasRegexElements(pattern) {
		return fmt.Errorf("pattern does not contain regex elements")
	}
	if !isBalanced(pattern, []string{"[", "]", "{", "}", "(", ")"}) {
		return fmt.Errorf("unbalanced brackets in pattern")
	}
	return nil
}

// santiy check for regex elements
func hasRegexElements(pattern string) bool {
	regexElements := []string{"[", "]", "(", ")", "{", "}", "*", "+", "?", "|", ".", "'", "\""}
	for _, elem := range regexElements {
		if strings.Contains(pattern, elem) {
			return true
		}
	}
	return false
}

// make sure all opening brackets have closing brackets
func isBalanced(pattern string, pairs []string) bool {
	counts := make(map[string]int)
	for _, char := range pattern {
		for _, pair := range pairs {
			if string(char) == pair {
				counts[pair]++
			}
		}
	}
	return counts["["] == counts["]"] && counts["{"] == counts["}"] && counts["("] == counts[")"]
}
