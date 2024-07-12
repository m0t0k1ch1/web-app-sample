package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"app/domain/service"
	"app/gen/gqlgen"
	"context"
)

// Noop is the resolver for the noop field.
func (r *mutationResolver) Noop(ctx context.Context, input gqlgen.NoopInput) (*gqlgen.NoopPayload, error) {
	return &gqlgen.NoopPayload{
		ClientMutationId: input.ClientMutationId,
	}, nil
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (gqlgen.Node, error) {
	var node gqlgen.Node
	{
		out, err := r.nodeService.Get(ctx, service.NodeServiceGetInput{
			ID: id,
		})
		if err != nil {
			return nil, err
		}

		node = out.Node
	}

	return node, nil
}

// Mutation returns gqlgen.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgen.MutationResolver { return &mutationResolver{r} }

// Query returns gqlgen.QueryResolver implementation.
func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
