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