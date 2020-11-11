package main

// download
// github
//
import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	fmt.Println("does this work")

	fileURL := "https://cdn.stocksnap.io/img-thumbs/960w/autumn-tree_M7PLVOJYNZ.jpg"
	err := DownloadFile("rando.jpg", fileURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + fileURL)

	Clone()

}

func Clone() {
	// Clone the given repository to the given directory
	fmt.Println("git clone https://github.com/deriggi/ffs")

	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/deriggi/ffs",
		Progress: os.Stdout,
	})
	fmt.Println(err)
	// CheckIfError(err)
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
