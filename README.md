# Go User Experience project

This is the source code repository for the `goux` project.

The *Go User Experience* project aims to improve the interaction the `Go`
*compiler* and the *programmer*, by improving error messages.

The project is inspired by these blog articles by the `Elm` *language* author:

* http://elm-lang.org/blog/compiler-errors-for-humans
* http://elm-lang.org/blog/compilers-as-assistants

The project repository contains minimal `Go` source code that cause a specified
error to be reported by the `Go` *compiler*. The source code are organized
based on where, in the `gc` source code, the error is reported.
