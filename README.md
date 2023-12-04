# jvm-in-go

[Original Resource](https://github.com/zxh0/jvmgo-book)

## How to Build and Test

I record all code in different chapter branch. The previous chapter can be used as the starter for the following one.

This is a method to map and pull all the remote branch.

```shell
$ git branch -r | grep -v 'main' | grep -v '\->' | \ # deselect branch mapped to local
    while read remote; do git branch --track "${remote#origin/}" "$remote"; done
branch 'chapter01' set up to track 'origin/chapter01'.
branch 'chapter02' set up to track 'origin/chapter02'.
branch 'chapter03' set up to track 'origin/chapter03'.
branch 'chapter04' set up to track 'origin/chapter04'.
branch 'chapter05' set up to track 'origin/chapter05'.
branch 'chapter06' set up to track 'origin/chapter06'.
branch 'chapter07' set up to track 'origin/chapter07'.
branch 'chapter08' set up to track 'origin/chapter08'.
branch 'chapter09' set up to track 'origin/chapter09'.
branch 'chapter10' set up to track 'origin/chapter10'.
branch 'chapter11' set up to track 'origin/chapter11'.

$ git fetch --all

$ git pull --all
```

## Noticeable Resources

[https://go.dev/doc/effective_go#defer]

## TODO

- [ ] Using `go.work` to track across multiple modules into one workspace.

- [ ] Using `go.mod` to use local relative modules.
