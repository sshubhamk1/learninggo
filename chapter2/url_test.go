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
