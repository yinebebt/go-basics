package blogposts_test

import (
	blogposts "gotest/blogposts" // if the path is  too long
	"reflect"
	"testing"
	"testing/fstest" // fstest implements support for testing implementations and users of file systems, as httptest in http
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: Post 2")},
	}
	//A MapFS is a simple in-memory file system for use in tests, represented as a map from path names to information about the files.

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}
	got := posts[0].Title
	want := blogposts.Post{Title: " Post 1", Description: ""}.Title
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
	t.Run("Test-two", func(t *testing.T) {
		const (
			firstBody = `Title:Post 1
Description:Description 1
Tags:tdd,Go
---
Hello World`
			secondBody = `Title: Post 2
Description: Description 2
Tags:rust,come
---
Body gos here below ---`
		)

		fs2 := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts2, err := blogposts.NewPostsFromFS(fs2)
		if err != nil {
			t.Fatal(err)
		}

		assertPost(t, posts2[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "Go"},
			Body:        "Hello World",
		})

	})
}

// helper function-assertPost
func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
