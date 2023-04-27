# check-deprecated

check-deprecated is a Go static analysis tool.

It finds malformed deprecated comments, and finds the use of deprecated function/variable/constant/field

## Purpose

We need an analyzer that checks deprecated usage. But
as [this issue](https://github.com/golangci/golangci-lint/issues/439#issuecomment-854368728) says,
> Actually it looks like staticcheck does have an appropriate check (SA1019) but for whatever reason it doesn't seem to
> work. Might open a separate issue to discuss that.

And I think SA1019 is too strict about the deprecated comments,
so that some situations cannot be detected.

We use a very loose detection patterns, you can even customize it.

```go
patterns := []string{"MyDeprecated: "}

// MyDeprecated: don't use it.
func AFunc()  {}
```

Of course, standard comments are always popular,
so this repo also provides an analysis to detect malformed deprecated comments.

