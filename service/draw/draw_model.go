package draw

import "context"

type DrawModel interface {
	// Draw from platform
	Draw(ctx context.Context, r *Request) (response *Response, err error)
}
