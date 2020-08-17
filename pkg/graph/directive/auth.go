package directive

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	fmt.Print("auth")
	/*switch obj.(type) {
	case string:
		return strings.ToLower(obj.(string)), nil
	}*/
	return next(ctx)
}
