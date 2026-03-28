package url

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	const uri = "https://github.com/sshubhamk1"

	got, err := Parse(uri)
	if err != nil {
		t.Fatalf("Parse (%q) err = %q, want <nil>", uri, err)
	}
	want := &URL{Scheme: "https", Host: "github.com", Path: "sshubhamk1"}
	if *got != *want {
		t.Errorf("Parse (%q) \ngot %#v\nwant %#v", uri, got, want)
	}
}

func TestString(t *testing.T) {
	u := &URL{Scheme: "https", Host: "github.com", Path: "sshubhamk1"}
	got := u.String()
	want := "https://github.com/sshubhamk1"
	if got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestParseWithoutPath(t *testing.T) {
	const uri = "https://github.com"
	got, err := Parse(uri)
	if err != nil {
		t.Fatalf("parse(%q), err = %q, want <nil>", uri, err)
	}
	want := &URL{Scheme: "https", Host: "github.com"}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Parse (%q) mismatch (-want +got): \n%s", uri, diff)
	}
}

var parseTests = []struct {
	name string
	uri  string
	want *URL
}{
	{
		name: "full",
		uri:  "https://github.com/sshubhamk1",
		want: &URL{Scheme: "https", Host: "github.com", Path: "sshubhamk1"},
	},
	{
		name: "without_path",
		uri:  "https://github.com",
		want: &URL{Scheme: "https", Host: "github.com", Path: ""},
	},
}

func TestParseTable(t *testing.T) {
	for _, tt := range parseTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.uri)
			if err != nil {
				t.Fatalf("Parse (%q) err = %v, want <nil>", tt.uri, err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Parse (%q) mismatch (-want +got): \n%s", tt.uri, diff)
			}
		})

	}
}
