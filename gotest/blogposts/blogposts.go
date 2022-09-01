package blogposts

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/fs" //Package fs defines basic interfaces to a file system.
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) { //change ftest.MapFS to fs.FS to have a general interface.
	dir, err := fs.ReadDir(fileSystem, ".")
	var posts []Post
	if err != nil {
		return nil, err
	}
	for _, f := range dir {
		// posts = append(posts, Post{})
		post, err := getPost(fileSystem, f)
		if err != nil {
			return nil, err //todo: needs clarification, should we totally fail if one file fails? or just ignore?
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) { //you can pass string instead of fs.DirEntry
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile fs.File) (Post, error) { //you can pass io.Reader instead of fs.File
	// postData, err := io.ReadAll(postFile)
	// if err != nil {
	// 	return Post{}, err
	// }
	// post := Post{Title: string(postData)[7:]} //[7:] of []byte{"Title: Post 1"} = []byte{"Post 1"}
	// return post, nil
	//v-2
	scanner := bufio.NewScanner(postFile)
	readMetaLine := func(tagname string) string {
		scanner.Scan() //the scan stops, either by reaching the end of the input(postFile) or an error.
		return strings.TrimPrefix(scanner.Text(), tagname)
	}
	title := readMetaLine("Title:")
	description := readMetaLine("Description:")
	tags := strings.Split(readMetaLine("Tags:"), ",")

	return Post{Title: title,
		Description: description,
		Tags:        tags,
		Body:        readbody(scanner),
	}, nil

}
func readbody(bufscan *bufio.Scanner) string { //what if I used NewScanner here
	bufscan.Scan() // skip one line (---)
	bufr := bytes.Buffer{}
	for bufscan.Scan() {
		fmt.Fprintln(&bufr, bufscan.Text())
	}
	return strings.TrimSuffix(bufr.String(), "\n")
}

// StubFailingFS can implement fs.ReaderDir because it have a mehtod Open which is expected by the interface.
type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}
