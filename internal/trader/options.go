package trader

import "fmt"

func WithProvider(providers ...string) func(t *trader) {
	return func(t *trader) {
		// t.providers = providers
		fmt.Println("WithProvider", providers)
	}
}
