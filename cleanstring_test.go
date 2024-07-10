package cleanstring

import "testing"
import "strings"

func TestSentence(t *testing.T) {
	got := Format("The quick brown fox jumps over the lazy dog.")
	want := "the_quick_brown_fox_jumps_over_the_lazy_dog"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestUnbrokenNonAlphanum(t *testing.T) {
	got := Format("Acme & Co.")
	want := "acme_co"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestUnbrokenNonAlphanumHead(t *testing.T) {
	got := Format("... hello? Is anybody there?")
	want := "hello_is_anybody_there"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestUnbrokenNonAlphanumLast(t *testing.T) {
	got := Format("What did you just say?!?!")
	want := "what_did_you_just_say"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestOnlyNonAlphanum(t *testing.T) {
	got := Format("???")
	want := ""
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestApostrophe(t *testing.T) {
	got := Format("Don't say that!")
	want := "dont_say_that"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestQuote(t *testing.T) {
	got := Format("\"To be, or not to be\"")
	want := "to_be_or_not_to_be"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestUnicode(t *testing.T) {
	got := Format("Hello 👋")
	want := "hello"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestOnlyUnicode(t *testing.T) {
	got := Format("가다")
	want := ""
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestSlice(t *testing.T) {
	got := Join([]string{"One", "Two", "Three"})
	want := "one+two+three"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestMap(t *testing.T) {
	got := Associate(map[string]string{"One": "Uno", "Two": "Dos", "Three": "Tres"})
	want := "one-uno+two-dos+three-tres"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestBoth(t *testing.T) {
	sl := Join(strings.Split("Alpha & Beta", "&"))
	st := Format("Gamma")
	m := map[string]string{sl: st}
	got := Associate(m)
	want := "alpha+beta-gamma"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
