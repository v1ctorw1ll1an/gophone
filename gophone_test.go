package gophone

import "testing"

func TestPhone(t *testing.T) {
	t.Run("test add prefix", func(t *testing.T) {
		want := "5531982273456"
		got := FormatPhoneBr("31982273456", "55")

		if want != got {
			t.Errorf("want: %s got: %s", want, got)
		}
	})
}
