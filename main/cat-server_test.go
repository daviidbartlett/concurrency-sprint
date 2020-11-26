package catserver

import (
	"reflect"
	"strings"
	"testing"
)

func TestCheckServerStatus(t *testing.T) {
	ch := make(chan string)

	go checkServerStatus(ch)

	got := <-ch
	want := "200 - the server is good"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestFetchBannerContent(t *testing.T) {
	ch := make(chan map[string]string)
	tests := []struct {
		key   string
		value string
	}{{
		"title",
		"Kitty Litter",
	},
		{
			"bannerImg",
			"https://riotfest.org/wp-content/uploads/2017/10/AcT9YIL.jpg",
		}}
	t.Run("banner has default properties", func(t *testing.T) {

		go fetchBannerContent(ch)

		got := <-ch

		for _, test := range tests {
			if got[test.key] != test.value {
				t.Errorf("got %s, want %s", got[test.key], test.value)
			}
		}
	})
	t.Run("copyrightYear has been updated", func(t *testing.T) {

		go fetchBannerContent(ch)

		got := <-ch
		want := "2020"
		if got["copyrightYear"] != want {
			t.Errorf("got copyright year %s, want %s", got["copyrightYear"], want)
		}
	})

}

func TestFetchAllOwners(t *testing.T) {
	ch := make(chan []string)
	t.Run("returns all owners", func(t *testing.T) {

		go fetchAllOwners(ch)

		got := <-ch
		want := ownerContent

		for i, owner := range got {
			if !strings.EqualFold(owner, want[i]) {
				t.Errorf("got %v, want %v", owner, want[i])
			}
		}
	})
	t.Run("all owners are lowercased", func(t *testing.T) {
		go fetchAllOwners(ch)

		got := <-ch
		want := ownerContent

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
