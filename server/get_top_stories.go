package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"context"

	"github.com/joscelyn/urban"
	"github.com/peterhellberg/hn"
	"google.golang.org/protobuf/types/known/emptypb"
)

type indexItem struct {
	Index int
	Item  *hn.Item
}

var (
	items    = map[int]*hn.Item{}
	messages = make(chan indexItem)
	maxPosts = 10
)

// GetTopStories fetch hacker news top 10 stories and return the title and url of each
func (s UrbanServer) GetTopStories(ctx context.Context, r *emptypb.Empty) (*urban.GetTopStoriesResponse, error) {
	hn := hn.NewClient(&http.Client{
		Timeout: time.Duration(5 * time.Second),
	})

	ids, err := hn.TopStories()
	if err != nil {
		return nil, errors.New(fmt.Sprint("unable to fetch hn top stories", err))
	}

	go func() {
		for i := range messages {
			items[i.Index] = i.Item
		}
	}()

	var wg sync.WaitGroup

	for i, id := range ids[:maxPosts] {
		wg.Add(1)
		go func(i, id int) {
			defer wg.Done()

			item, err := hn.Item(id)
			if err != nil {
				log.Print("unable to fetch hn item", err)
			}

			messages <- indexItem{i, item}
		}(i, id)
	}

	wg.Wait()

	res := urban.GetTopStoriesResponse{}
	for _, item := range items {
		res.Posts = append(res.Posts, &urban.HNPost{
			Title: item.Title,
			Url:   item.URL,
		})
	}

	return &res, nil
}
