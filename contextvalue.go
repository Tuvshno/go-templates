package main

import (
	"context"
	"errors"
	"fmt"
)

func main() {
	f := func(ctx context.Context, a int, b int) (int, error) {
		switch ctx.Value("action") {
		case "+":
			return a + b, nil
		case "-":
			return a - b, nil
		default:
			return 0, errors.New("unkown operator")
		}
	}

	ctx := context.WithValue(context.Background(), "action", "+")
	val, err := f(ctx, 2, 4)
	fmt.Println(val, err)
	ctx2 := context.WithValue(context.Background(), "action", "-")
	val, err = f(ctx2, 10, 6)
	fmt.Println(val, err)
	ctx3 := context.WithValue(context.Background(), "action", "/")
	val, err = f(ctx3, 10, 6)
	fmt.Println(val, err)
}
