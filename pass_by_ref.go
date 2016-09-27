//usr/bin/env go run "$0" "$@" ; exit "$?"
// Shebang for go is not #!/usr/bin/env go
package main

/*
go run ./sample1_run.go

My playlist contain : (2) Pinky Brooker, Harley Hola
I add Big Deal in may playlist : (3) Pinky Brooker, Harley Hola, Big Deal
After this, my playlist contain : (3) Pinky Brooker, Harley Hola, Big Deal

I add Get Down to my playlist : (4) Pinky Brooker, Harley Hola, Big Deal, Get Down
...but without pointer => no change !
After this, my playlist contain : (3) Pinky Brooker, Harley Hola, Big Deal

*/

import (
	_ "./samples" // go run is script oriented: Accept relative path import
	"fmt"         // Show result
	"os"
	"strings" // Tooling around strings (here, make a Join)
)

// Playlist is a slide : A specific object describing the use of
// underlying fixed-size array
type Playlist []string

type Album struct {
	title   string
	authors []string
}

// fmt functions will call method String() before attemps to print internal
// data (by DuckTyping analysis). => "(size) song_1, song_2, ..."
func (s Playlist) String() string {
	return fmt.Sprintf(
		"(%d) %s",
		len(s),
		strings.Join(s, ", "),
	)
}

// We create a Pointer on object creation if the current usage implie always a
// change on original. We says : Always works "by reference" for this object.
func main() {
	// Create a Playlist, but return a pointer that is assigned to myPlaylist
	// Any use of "&" on variable or object creation (like here) return a Pointer.
	myPlaylist := &Playlist{"Pinky Brooker", "Harley Hola"}
	fmt.Printf(
		"My playlist contain : %s\n",
		myPlaylist,
	)

	// Use of *
	// * inside function signature / call show the use of Pointer
	// * outside say "Get access to the object pointed by this variable"
	//   Also called Dereferencing: we dereference the pointer to access the object

	fmt.Println("\n WITH POINTERS \n")
	fmt.Printf(
		"I add Big Deal in may playlist : %s\n",
		addTo(myPlaylist, "Big Deal"),
	)
	fmt.Printf(
		"After this, my playlist contain : %s\n\n",
		myPlaylist,
	)

	fmt.Println("\n WITHOUT POINTERS \n")
	fmt.Printf(
		"I can change the content of underlyng slide: %s\n"+
			"Because the slide is also a pointer, I gain access to the underlying array\n",
		changeLastPlaylist(*myPlaylist, "Beegies"),
	)
	fmt.Printf(
		"After this, my playlist contain : %s\n\n",
		myPlaylist,
	)

	fmt.Printf(
		"I add Get Down to my playlist : %s\n...but without pointer => no change !\n",
		failedAddTo(*myPlaylist, "Get Down"),
	)
	fmt.Printf(
		"After this, my playlist contain : %s\n\n",
		myPlaylist,
	)

	// Without pointer on string into struct
	var album Album
	fmt.Printf(
		"I set the album's title to : %s\n...but without pointer => no change !\n",
		setAlbumTitle(album, "My super album"),
	)
	fmt.Printf(
		"After this, my album title is : %s\n\n",
		album.title,
	)

	// Without pointer on slide into struct
	fmt.Printf(
		"I add author to album list : %s\n...but without pointer => no change !\n",
		addAlbumAuthor(album, "Edwidg Brand"),
	)
	fmt.Printf(
		"After this, my album authors is : %s\n\n",
		album.authors,
	)

	// Conclude :
	// If the object is :
	// - Subject to changes on normal use
	// - Or contain large data (other that pointers like slides or map)
	// ==> Alway use pointer.

	os.Exit(0)
}

// Get a copy of the pointer (point to the original Playlist).
// Dereference the pointer for accessing to the original Playlist
// => The recever (as original) get now the modified Playlist
func addTo(s *Playlist, title string) Playlist {
	*s = append(*s, title) // Create a bigger array & affect it to the original Playlist
	return *s
}

// Update content (without change on len/cap) is ok without pointer:
func changeLastPlaylist(s Playlist, title string) Playlist {
	idx := len(s) - 1
	s[idx] = title
	return s
}

// Get a copy of the Playlist.
// We can change values of underlying array but not the Playlist itself
// => The recever (as a copy) get now the modified Playlist, original remain unmodified
func failedAddTo(s Playlist, title string) Playlist {
	s = append(s, title) // Create a bigger array & affect it to the copy of the Playlist
	return s
}

func setAlbumTitle(album Album, title string) string {
	album.title = title
	return album.title
}

func addAlbumAuthor(album Album, author string) string {
	album.authors = append(album.authors, author)
	return strings.Join(album.authors, ", ")
}
