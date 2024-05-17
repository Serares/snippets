package structural

import "testing"

func TestAdapter(t *testing.T) {

	var myClient = func(arbitraryClient Target) string {
		return arbitraryClient.Request()
	}

	mySpecificClient := Adapter{
		Adaptee: &Adaptee{},
	}

	t.Run("myClient", func(t *testing.T) {
		got := myClient(&mySpecificClient)
		want := "Specific Request"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

func TestFacade(t *testing.T) {

	ht := NewHomeTheaterFacade(
		&DVDPlayer{},
		&SoundSystem{},
		&Projector{},
	)

	ht.WatchMovie("Titanic")
	ht.EndMovie()

}
