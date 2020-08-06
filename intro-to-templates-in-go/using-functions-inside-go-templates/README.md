# Using Functions Inside Go Templates

Part 3 on An Introduction to Templates in Go series.

In this tutorial we are going to cover how to use template functions like `and`, `eq` and `index` to add some basic logic to our templates.

## `and` function
- The `and` function takes in two arguments, lets call them `a` and `b`, and then runs logic roughly equivalent to `if a then b else a`.
- `and` is a function and NOT a logic operation.
- The template package also provides an `or` function that operates much like `and` except it will short circuit when true, i.e., `if a then a else b`; here `b` is never evaluated if `a` is not empty.

## The comparison functions
- The `html/template` package provides us with a few classes to help carry out comparisons.
- These are:

>- `eq` - Returns the boolean truth of `arg1 == arg2`
>- `ne` - Returns the boolean truth of `arg1 != arg2`
>- `lt` - Returns the boolean truth of `arg1 < arg2`
>- `le` - Returns the boolean truth of `arg1 <= arg2`
>- `gt` - Returns the boolean truth of `arg1 > arg2`
>- `ge` - Returns the boolean truth of `arg1 >= arg2`

- These are used similarly to how `and` and `or` are used, where you first type the function and type the arguments. For example:

```HTML
{{if (ge .Usage .Limit)}}
    <p class="danger">
        You have reached your API usage limit. Please upgrade or contact support for more help.
    </p>
{{else if (gt .Usage .Warning)}}
    <p class="warning">
        You have use {{.Usage}} of {{.Limit}} API calls and are nearing your limit. Have you considered blah blah blah ...
    </p>
{{else if (eq .Usage 0)}}
    <p>
        You have not used the API yet. What are waiting for?
    </p>
{{else}}
    <p>
        You have used {{.Usage}} of {{.Limit}} API calls.
    </p>
{{end}}
```