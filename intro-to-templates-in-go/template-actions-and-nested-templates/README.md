# Template Actions and Nested Templates in Go

Part 2 of the Intro to Templates in Go series.

Up until now we have been outputting templates to the terminal, but as we start to dive into more HTML this starts to make less sense. Instead we want to visualize the HTML being generated in a web browser. To do this we first need to set up a web server that will render our HTML templates.

## Removing whitespace

```HTML
<h1>
    Hello,
    {{if .Name}}
        {{.Name}}
    {{- else}}
        there
    {{- end}}!
</h1>
```

Versus:

```HTML
<h1>Hello, {{if .Name}}{{.Name}}{{else}}there{{end}}!</h1>
```

Both snippets accomplish the same but the first option can be generated even when working with a space sensitive language like Python.

In the first snippet we are telling the template package that we don't want all spaces between the `Name` variable and whatever comes after it by putting the minus character at the front of the else keyword, and we are also doing the same with the end keyword on the second to last line.

## `range` action

The most common source of confusion with the `range` action is that we are accessing individual attributes of a widget without needing to use an index or any other accessor inside the `.Widgets` value. This is because the `range` action will set the value of each object in the collection to the (`.`) inside of the range block. For example, if you were to render `{{.}}` inside of the range block you would see the same output as if you used `fmt.Println()` on a `Widget` object.