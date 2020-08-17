package directive

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Length validation
func MaxLength(ctx context.Context, obj interface{}, next graphql.Resolver, max int) (res interface{}, err error) {
	fmt.Println("MaxLength")
	v, _ := next(ctx)
	fmt.Println(v)
	if len(v.(string)) > max {
		return nil, gqlerror.Errorf("format")
	}
	switch obj.(type) {
	case string:
		src := obj.(string)
		/*if len(src) < min {
			return nil, fmt.Errorf("the length of the string `%s` is less than allowed (%d)", src, min)
		}*/
		if len(src) > max {
			return nil, fmt.Errorf("the length of the string `%s` has exceeded the maximum allowed (%d)", src, max)
		}

	}
	return next(ctx)
}

func Demo(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	fmt.Println("Demo")
	/*v, _ := next(ctx)
	if len(v.(string)) > 20 {
		return nil, gqlerror.Errorf("format")
	}
	fmt.Println(v)*/
	return next(ctx)
}
