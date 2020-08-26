# Creating the V in MVC

Part 4 of 4

In this tutorial we are going to look at how to create a reusable view layer for the MVC architectural pattern using Golang and the `html/template` package. This includes setting up a layout with Twitter's Bootstrap, reusing this layout with a dynamic template, and adding flash messages across our views.

## What is a view

In the MVC pattern, a view is the portion of code responsible for rendering data. Typically this means rendering HTML, but your view layer could also render `JSON`, `XML`, or even text.

Views are generally only responsible for rendering data. This means that business logic is frowned upon inside the view layer, as this often leads to bugs. Instead it is common to prepare data that your view will need ahead of time in the controller, and then pass it to the view to use. This means that our view will need to provide some way for our controllers to pass data to it.

Reuse of views is also fairly common practice in the MVC apps. For example, if you have the endpoints `GET /signup` and `POST /signup`, the `POST` endpoint might end up rendering the same view as the `GET` endpoint if there is an error with the data the user provided. This means we should try to make these views reusable.

In my application I also like to have the ability to create a common layouts that will be shared across my app, and then use those inside my view layer. This ins't a requirement in the MVC pattern, but it is incredibly helpful so we will be adding this functionality as well.

## Creating a layout

---
**Variadic functions**

`template.ParseFiles()` is a variadic function, which basically means it takes a variable number of arguments. Since we want to pass in all of our strings in the slice as our arguments, we use the `...` syntax to expand the slice into individual arguments without having to type that out.

---

Looking at the `handler()` function:

```Go
func handler(w http.ResponseWriter, r *http.Request) {
    bootstrap.ExecuteTemplate(w, "bootstrap", nil)
}
```

This should look familiar with one exception - we are executing a specific template, and we pass in the name of that template as the second argument. In previous parts of this tutorial we would instead execute the default template, but we want to use the specific "bootstrap" template now.

## Yielding dynamic content in a layout