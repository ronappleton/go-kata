package kata

// Pipeline — TODO: implement as specified in README.md
//
// Signature:
//   func Pipeline[T any](ctx context.Context, in <-chan T, stages ...func(context.Context, <-chan T) <-chan T) <-chan T
func Pipeline[T any](ctx context.Context, in <-chan T, stages ...func(context.Context, <-chan T) <-chan T) <-chan T {
    // TODO: your implementation here
    return nil
}
