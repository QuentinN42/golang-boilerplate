package functool

import "context"

func sendAll[K any](ctx context.Context, ch chan K, data []K) {
	for _, elm := range data {
		if ctx.Done() != nil {
			break
		}
		ch <- elm
	}
	close(ch)
}

func Iter[K any](ctx context.Context, arr []K) chan K {
	chn := make(chan K)
	go sendAll(ctx, chn, arr)
	return chn
}

func Reduce[K any, V any](ch chan K, f func(K, V) V, i V) V {
	current := i
	for elm := range ch {
		current = f(elm, current)
	}
	return current
}

func Collect[K any](ch chan K) []K {
	return Reduce(
		ch,
		func(elm K, acc []K) []K { return append(acc, elm) },
		[]K{},
	)
}

func Apply[K any, V any](ch chan K, f func(K) V) chan V {
	chn := make(chan V)
	go func() {
		for elm := range ch {
			chn <- f(elm)
		}
		close(chn)
	}()
	return chn
}

func Filter[K any](ch chan K, f func(K) bool) chan K {
	chn := make(chan K)
	go func() {
		for elm := range ch {
			if f(elm) {
				chn <- elm
			}
		}
		close(chn)
	}()
	return chn
}
