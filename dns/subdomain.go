package main

import (
	"fmt"
	"regexp"
	"strings"
)

// MaxLenError returns a string explanation of a "string too long" validation
// failure.
func MaxLenError(length int) string {
	return fmt.Sprintf("must be no more than %d characters", length)
}

// RegexError returns a string explanation of a regex validation failure.
func RegexError(fmt string, examples ...string) string {
	s := "must match the regex " + fmt
	if len(examples) == 0 {
		return s
	}
	s += " (e.g. "
	for i := range examples {
		if i > 0 {
			s += " or "
		}
		s += "'" + examples[i] + "'"
	}
	return s + ")"
}

// maskTrailingDash replaces the final character of a string with a subdomain safe
// value if is a dash.
func maskTrailingDash(name string) string {
	if strings.HasSuffix(name, "-") {
		return name[:len(name)-2] + "a"
	}
	return name
}

const dns1123LabelFmt string = "[a-z0-9]([-a-z0-9]*[a-z0-9])?"
const DNS1123LabelMaxLength int = 63

var dns1123LabelRegexp = regexp.MustCompile("^" + dns1123LabelFmt + "$")

// IsDNS1123Label tests for a string that conforms to the definition of a label in
// DNS (RFC 1123).
func IsDNS1123Label(value string) []string {
	var errs []string
	if len(value) > DNS1123LabelMaxLength {
		errs = append(errs, MaxLenError(DNS1123LabelMaxLength))
	}
	if !dns1123LabelRegexp.MatchString(value) {
		errs = append(errs, RegexError(dns1123LabelFmt, "my-name", "123-abc"))
	}
	return errs
}

const dns1123SubdomainFmt string = dns1123LabelFmt + "(\\." + dns1123LabelFmt + ")*"
const DNS1123SubdomainMaxLength int = 253

var dns1123SubdomainRegexp = regexp.MustCompile("^" + dns1123SubdomainFmt + "$")

// IsDNS1123Subdomain tests for a string that conforms to the definition of a
// subdomain in DNS (RFC 1123).
func IsDNS1123Subdomain(value string) []string {
	var errs []string
	if len(value) > DNS1123SubdomainMaxLength {
		errs = append(errs, MaxLenError(DNS1123SubdomainMaxLength))
	}
	if !dns1123SubdomainRegexp.MatchString(value) {
		errs = append(errs, RegexError(dns1123SubdomainFmt, "example.com"))
	}
	return errs
}

func main() {
	domain := "www.yeepay.com"
	//domain := "test-nginx.1.8"
	//domain := "aaa"
	//fmt.Printf("%s\n%s\n", domain, maskTrailingDash(domain))

	fmt.Printf("label: %s\n", dns1123LabelFmt)
	fmt.Printf("subdomain: %s\n", dns1123SubdomainFmt)
	fmt.Printf("domain: %s\n", domain)
	fmt.Printf("%v\n", IsDNS1123Subdomain(domain))

	fmt.Printf("subdomainRegexp: %s\n", dns1123SubdomainRegexp.String())

}
