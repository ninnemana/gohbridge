package light

import context "golang.org/x/net/context"

// Interactor describes the functionality of the Light service.
type Interactor interface {
	All(context.Context, *ListParams) ([]Light, error)
	New(context.Context, *NewParams) (*Scan, error)
	Search(context.Context, *SearchParams) (*SearchResult, error)
	Get(context.Context, *GetParams) (*Light, error)
}
