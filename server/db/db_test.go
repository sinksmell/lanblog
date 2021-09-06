package db

import (
	"context"
	"lanblog/model"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoDB_Post(t *testing.T) {

	ctx := context.Background()

	m, err := mongo.Connect(ctx,&options.ClientOptions{
		Hosts: []string{"localhost"},
	})

	if err != nil {
		t.Fatal(err)
	}

	defer m.Disconnect(ctx)

	db := NewMongo(m)

	p := &model.Post{
		ID:         0,
		Subject:    "this is post one",
		Content:    "RT",
		Conver:     "https://pkg.go.dev/static/shared/logo/go-blue.svg",
		CategoryID: 1,
		TagID:      []uint32{1, 2, 3, 4},
		IsDeleted:  0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// create a post
	postId, err := db.CreatePost(ctx, p)
	if err != nil {
		t.Fatal(err)
	}

	require.NotZero(t, postId)

	// edit post
	p.ID = uint64(postId)
	p.Content = "RT [updated]"
	p.Subject = "this is post one [updated]"

	affected, err := db.EditPost(ctx, p)
	if err != nil {
		t.Fatal(err)
	}

	require.NotZero(t, affected)

	// list post
	list, err := db.ListPost(ctx, &model.Pager{Limit: 10, Offset: 0})
	if err != nil {
		t.Fatal(err)
	}

	require.NotZero(t, len(list))

	// delete then list post
	affected, err = db.DeletePost(ctx, postId)
	if err != nil {
		t.Fatal(err)
	}

	require.NotZero(t, affected)

	list, err = db.ListPost(ctx, &model.Pager{Limit: 10, Offset: 0})
	if err != nil {
		t.Fatal(err)
	}

	for _, p := range list {
		if p.ID == uint64(postId) {
			t.Fatal()
		}
	}
}
