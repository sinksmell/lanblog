package db

import (
	"context"
	"lanblog/model"
	"time"

	"github.com/douyu/jupiter/pkg/xlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type DB interface {
	GetTechnologies() ([]*model.Technology, error)

	// Post API
	CreatePost(ctx context.Context, p *model.Post) (int64, error)

	EditPost(ctx context.Context, p *model.Post) (int64, error)

	ListPost(ctx context.Context, pager *model.Pager) ([]*model.Post, error)

	DeletePost(ctx context.Context, postId int64) (int64, error)

	// Category API
	AddCategory(ctx context.Context, t *model.Category) (int64, error)

	DeleteCategory(ctx context.Context, topicId uint32) (affected int64, err error)

	GetAllCategory(ctx context.Context) ([]*model.Category, error)

	// Tag API
	AddTag(ctx context.Context, t *model.Tag) (id uint32, err error)

	GetAllTags(ctx context.Context) ([]*model.Tag, error)

	DeleteTag(ctx context.Context, tagId uint32) (affected int64, err error)
}

type MongoDB struct {
	client *mongo.Client
}

func NewMongo(client *mongo.Client) DB {
	return &MongoDB{client: client}
}

func (m *MongoDB) GetTechnologies() ([]*model.Technology, error) {
	var tech []*model.Technology
	return tech, nil
}

var (
	_post model.Post
	_cate model.Category
	_tag  model.Tag
)

// Post API
func (db *MongoDB) CreatePost(ctx context.Context, p *model.Post) (int64, error) {
	c := db.client.Database(_post.Database()).Collection(_post.Collection())

	// use snowflake uuid insted of increment id
	p.ID = uint64(time.Now().Unix())
	result, err := c.InsertOne(ctx, p)
	if err != nil {
		xlog.Error("create post err", zap.Any("post", p))
		return 0, err
	}

	xlog.Info("create post success", zap.Any("_id", result.InsertedID))

	return int64(p.ID), nil
}

func (db *MongoDB) EditPost(ctx context.Context, p *model.Post) (int64, error) {
	c := db.client.Database(_post.Database()).Collection(_post.Collection())

	filter := bson.M{
		"id": p.ID,
	}

	updated := bson.M{
		"$set": bson.D{
			{"subject", p.Subject},
			{"content", p.Content},
			{"cover", p.Conver},
			{"category_id", p.CategoryID},
			{"tag_id", p.TagID},
			{"updated_at", time.Now()},
		},
	}

	result, err := c.UpdateOne(ctx, filter, updated)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (db *MongoDB) ListPost(ctx context.Context, pager *model.Pager) ([]*model.Post, error) {
	c := db.client.Database(_post.Database()).Collection(_post.Collection())

	var (
		posts  []*model.Post
		cursor *mongo.Cursor
		err    error
	)

	filter := bson.M{
		"is_deleted": model.NotDeleted,
	}

	opt := &options.FindOptions{
		Limit: &pager.Limit,
		Skip:  &pager.Offset,
	}

	if cursor, err = c.Find(ctx, filter, opt); err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, err
}

func (db *MongoDB) DeletePost(ctx context.Context, postId int64) (int64, error) {
	c := db.client.Database(_post.Database()).Collection(_post.Collection())

	filter := bson.M{
		"id": postId,
	}

	result, err := c.UpdateOne(ctx, filter, bson.M{
		"$set": bson.M{
			"is_deleted": model.IsDeleted,
		},
	})
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

// Category API
func (db *MongoDB) AddCategory(ctx context.Context, t *model.Category) (int64, error) {
	return 0, nil
}

func (db *MongoDB) DeleteCategory(ctx context.Context, topicId uint32) (affected int64, err error) {
	return 0, nil
}

func (db *MongoDB) GetAllCategory(ctx context.Context) ([]*model.Category, error) {
	return nil, nil
}

// Tag API
func (db *MongoDB) AddTag(ctx context.Context, t *model.Tag) (id uint32, err error) {
	return 0, nil
}

func (db *MongoDB) GetAllTags(ctx context.Context) ([]*model.Tag, error) {
	return nil, nil
}

func (db *MongoDB) DeleteTag(ctx context.Context, tagId uint32) (affected int64, err error) {
	return 0, nil
}
