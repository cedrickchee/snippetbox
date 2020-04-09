package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

// EmailRX is regular expression (regex) for sanity checking the format of an
// email address. Use the regexp.MustCompile() function to parse a pattern and
// compile a regex.
// This returns a *regexp.Regexp object, or panics in the event of an error.
// Doing this once at runtime, and storing the compiled regex
// object in a variable, is more performant than re-compiling the pattern with
// every request.
// The pattern is currently recommended by the W3C and WHATWG.
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Form struct anonymously embeds a url.Values object (to hold the form data)
// and an Errors field to hold any validation errors for the form data.
type Form struct {
	url.Values
	Errors errors
}

// New function initialize a custom Form struct. Notice that this takes the form
// data as the parameter.
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required method checks that specific fields in the form data are present and
// not blank. If any fields fail this check, add the appropriate message to the
// form errors.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MaxLength method checks that a specific field in the form contains a maximum
// number of characters. If the check fails then add the appropriate message to
// the form errors.
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters)", d))
	}
}

// PermittedValues method checks that a specific field in the form matches one
// of a set of specific permitted values. If the check fails then add the
// appropriate message to the form errors.
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

// MinLength method checks that a specific field in the form contains a minimum
// number of characters. If the check fails then add the appropriate message to
// the form errors.
func (f *Form) MinLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < d {
		f.Errors.Add(field, fmt.Sprintf("This field is too short (minimum is %d characters)", d))
	}
}

// MatchesPattern method checks that a specific field in the form matches a
// regular expression. If the check fails then add the appropriate message to
// the form errors.
func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
	}
}

// Valid method returns true if there are no errors.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
