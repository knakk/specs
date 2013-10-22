package specs

import (
	"regexp"
	"testing"
)

// Specs is a wrapper for *testing.T
type Specs struct {
	t *testing.T
}

// Spec is a struct used for creating table-driven tests.
//
// ExpectAll takes []Spec as input.
type Spec struct {
	Actual   interface{}
	Expected interface{}
}

// New initializes the specs by storing a reference to the *testing.T.
func New(t *testing.T) *Specs {
	return &Specs{t}
}

// Expect checks for equality between expected and actual.
func (s *Specs) Expect(acutal, expected interface{}, msg ...string) {
	if actual != expected {
		if len(msg) > 0 {
			s.t.Error(msg[0])
		} else {
			s.t.Errorf("expected \"%+v\" to be \"%+v\"", actual, expected)
		}
	}
}

// ExpectMatches checks if the string actual has a match in expected.
func (s *Specs) ExpectMatches(actual, expected string, msg ...string) {
	r := regexp.MustCompile(expected)
	if !r.MatchString(actual) {
		s.t.Errorf("expected \"%+v\" to be a substring of \"%+v\"", expected, actual)
	}
}

// ExpectNotMatches checks if the string actual not has a match in expected.
func (s *Specs) ExpectNotMatches(actual, expected string, msg ...string) {
	r := regexp.MustCompile(expected)
	if r.MatchString(actual) {
		s.t.Errorf("expected \"%+v\" not to be a substring of \"%+v\"", expected, actual)
	}
}

// ExpectNot checks for inequality between expected and actual.
func (s *Specs) ExpectNot(actual, expected interface{}, msg ...string) {
	if actual == expected {
		if len(msg) > 0 {
			s.t.Error(msg[0])
		} else {
			s.t.Errorf("expected \"%+v\" not to be \"%+v\"", actual, expected)
		}
	}
}

// ExpectNil checks that actual is nil.
func (s *Specs) ExpectNil(actual interface{}, msg ...string) {
	if actual == nil {
		return
	}
	if len(msg) > 0 {
		s.t.Error(msg[0])
	} else {
		s.t.Errorf("expected \"%+v\" to be <nil>", actual)
	}
}

// ExpectNilFatal checks that actual is nil, and calls t.Fatal if not.
func (s *Specs) ExpectNilFatal(actual interface{}, msg ...string) {
	if actual == nil {
		return
	}
	if len(msg) > 0 {
		s.t.Fatal(msg[0])
	} else {
		s.t.Fatalf("expected \"%+v\" to be <nil>", actual)
	}
}

// ExpectNotNil checks that actual is not nil.
func (s *Specs) ExpectNotNil(actual interface{}, msg ...string) {
	if actual != nil {
		return
	}
	if len(msg) > 0 {
		s.t.Error(msg[0])
	} else {
		s.t.Errorf("expected \"%+v\" to be not <nil>", actual)
	}
}

// ErrExpect checks if err is nil before comparing expected and acutal.
// It calls t.Fatal(err) and shortcircut the test if err is not nil.
func (s *Specs) ErrExpect(err, actual, expected interface{}, msg ...string) {
	if err != nil {
		if len(msg) > 0 {
			s.t.Fatal(msg[0])
		} else {
			s.t.Fatal(err)
		}
	}
	if actual != expected {
		s.t.Errorf("expected \"%+v\" to be \"%+v\"", actual, expected)
	}
}

// ExpectAll takes []Spec and runs through all the tests, checking if
// Spec.actual is equal to Spec.expected.
func (s *Specs) ExpectAll(tests []Spec) {
	for _, t := range tests {
		if t.Actual != t.Expected {
			s.t.Errorf("expected \"%+v\" to be \"%+v\"", t.Actual, t.Expected)
		}
	}
}
