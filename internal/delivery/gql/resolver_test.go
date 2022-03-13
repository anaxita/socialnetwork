package gql

import (
	"context"
	"encoding/json"
	"synergycommunity/internal/delivery/api/apimodel"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/domain/interactor"
	"synergycommunity/internal/domain/service"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"

	"github.com/99designs/gqlgen/client"
)

func TestGraphQLTagsHandler(t *testing.T) { //nolint:funlen
	t.Parallel()

	type args struct {
		repo tagStorage
		id   int64
		e    entity.Tag
	}

	type testCase struct {
		name      string
		args      args
		want      interface{}
		wantError bool
	}

	tableTests := []testCase{
		{
			name: "ok",
			args: args{
				repo: tagStorage{
					FTags: func(ctx context.Context, o entity.Options) ([]entity.Tag, int64,
						error) {
						return []entity.Tag{
							{
								ID:        1,
								Name:      "tag one",
								CreatedAt: time.Now().UTC().Round(time.Minute),
								UpdatedAt: time.Now().UTC().Round(time.Minute),
							},
							{
								ID:        2,
								Name:      "tag two",
								CreatedAt: time.Now().UTC().Round(time.Minute),
								UpdatedAt: time.Now().UTC().Round(time.Minute),
							},
						}, 2, nil
					},
				},
				id: 1,
				e:  entity.Tag{},
			},
			want: []entity.Tag{
				{
					ID:        1,
					Name:      "tag one",
					CreatedAt: time.Now().UTC().Round(time.Minute),
					UpdatedAt: time.Now().UTC().Round(time.Minute),
				},
				{
					ID:        2,
					Name:      "tag two",
					CreatedAt: time.Now().UTC().Round(time.Minute),
					UpdatedAt: time.Now().UTC().Round(time.Minute),
				},
			},
			wantError: false,
		},
	}

	for _, tc := range tableTests {
		tc := tc

		s := service.NewTagService(&tc.args.repo)
		i := interactor.NewTagInteractor(s)
		mockInteractors := interactor.Interactors{Tags: i}
		handler := NewGQLHandler(&mockInteractors)
		c := client.New(handler)

		t.Run(
			tc.name, func(t *testing.T) {
				t.Parallel()

				var rsp map[string]interface{}

				err := c.Post(`{ showTags {
            items {
            created_at
            id
            title
            updated_at
        }
} }`, &rsp)

				if tc.wantError {
					assert.NotEqual(t, err, nil)
				} else {
					assert.Equal(t, err, nil)
				}

				// Decode rsp to struct
				var r struct {
					ShowTags apimodel.TagsWithPagination
				}

				jsonString, err := json.Marshal(rsp)
				assert.Equal(t, err, nil)

				err = json.Unmarshal(jsonString, &r)
				assert.Equal(t, err, nil)

				// assert.Equal(t, tc.want, r.ShowTags.Items)
				assert.IsEqual(tc.want, r.ShowTags.Items)
			},
		)
	}
}

type tagStorage struct {
	FTagByID             func(ctx context.Context, id int64) (entity.Tag, error)
	FUpdateTag           func(ctx context.Context, e entity.Tag) (entity.Tag, error)
	FDeleteTag           func(ctx context.Context, id int64) (int64, error)
	FTagsByModelObjectID func(ctx context.Context, modelID int64, objectID int64) ([]int64, error)
	FInsertTag           func(ctx context.Context, e entity.Tag) (int64, error)
	FTags                func(ctx context.Context, o entity.Options) ([]entity.Tag, int64, error)
}

func (t *tagStorage) InsertTag(ctx context.Context, e entity.Tag) (int64, error) {
	return t.FInsertTag(ctx, e)
}

func (t *tagStorage) Tags(ctx context.Context, o entity.Options) ([]entity.Tag, int64, error) {
	return t.FTags(ctx, o)
}

func (t *tagStorage) TagByID(ctx context.Context, id int64) (entity.Tag, error) {
	return t.FTagByID(ctx, id)
}

func (t *tagStorage) UpdateTag(ctx context.Context, e entity.Tag) (entity.Tag, error) {
	return t.FUpdateTag(ctx, e)
}

func (t *tagStorage) DeleteTag(ctx context.Context, id int64) (int64, error) {
	return t.FDeleteTag(ctx, id)
}

func (t *tagStorage) TagsByModelObjectID(ctx context.Context, modelID int64,
	objectID int64) ([]int64, error) {
	return t.FTagsByModelObjectID(ctx, modelID, objectID)
}
