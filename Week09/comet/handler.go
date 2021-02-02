package comet

import (
	"context"
)

type Handler interface {
	Handle(context.Context, string, *string)
}
