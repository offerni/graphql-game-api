package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning/game"
	"github.com/offerni/graphqllearning/graph/generated"
	"github.com/offerni/graphqllearning/graph/model"
	"github.com/offerni/graphqllearning/store"
)

// CreateGame is the resolver for the CreateGame field.
func (r *mutationResolver) CreateGame(ctx context.Context, opts model.NewGame) (*model.Game, error) {
	res, err := game.Create(ctx, game.CreateOpts{
		Name:    opts.Name,
		StoreID: opts.StoreID,
		Price:   opts.Price,
	})
	if err != nil {
		return nil, fmt.Errorf("ERROR: STORE NOT CREATED")
	}

	return &model.Game{
		ID:      res.ID,
		Name:    res.Name,
		StoreID: res.StoreID,
		Price:   res.Price,
	}, nil
}

// CreateStore is the resolver for the CreateStore field.
func (r *mutationResolver) CreateStore(ctx context.Context, opts model.NewStore) (*model.Store, error) {
	res, err := store.Create(ctx, store.CreateOpts{
		Name: opts.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("ERROR: STORE NOT CREATED")
	}

	return &model.Store{
		ID:   res.ID,
		Name: res.Name,
	}, nil
}

// Game is the resolver for the game field.
func (r *queryResolver) Game(ctx context.Context, id string) (*model.Game, error) {
	res, err := game.Fetch(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("ERROR: GAME NOT FOUND")
	}

	return &model.Game{
		ID:   res.ID,
		Name: res.Name,
	}, nil
}

// Store is the resolver for the store field.
func (r *queryResolver) Store(ctx context.Context, id string) (*model.Store, error) {
	res, err := store.Fetch(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("ERROR: STORE NOT FOUND")
	}

	return &model.Store{
		ID:   res.ID,
		Name: res.Name,
	}, nil
}

// Stores is the resolver for the stores field.
func (r *queryResolver) Stores(ctx context.Context) ([]*model.Store, error) {
	res, err := store.FetchAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("ERROR: STORES NOT FOUND")
	}

	stores := []*model.Store{}
	for _, store := range res {
		stores = append(stores, &model.Store{
			ID:   store.ID,
			Name: store.Name,
		})

	}

	return stores, nil
}

// Games is the resolver for the games field.
func (r *storeResolver) Games(ctx context.Context, obj *model.Store) ([]*model.Game, error) {
	res, _ := game.FetchGamesByStoreID(ctx, obj.ID)

	games := []*model.Game{}

	for _, game := range res {
		games = append(games, &model.Game{
			ID:      game.ID,
			StoreID: game.StoreID,
			Name:    game.Name,
			Price:   game.Price,
		})

	}

	return games, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Store returns generated.StoreResolver implementation.
func (r *Resolver) Store() generated.StoreResolver { return &storeResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type storeResolver struct{ *Resolver }
