# Introduction to Go
---

Go is a simple but powerful alternative to C++ and other web server languages that was developed by a group of devs who worked at Google. One of the focuses of Go (or Golang) was to reimagine what C++ would be like if developed today with current language needs. Go has become a popular enterprise tool, Kubernetes, Netflix, and Github are some examples of who is using Go.

## Breakdown
These are the core concepts that we will go over to get started in Go.

- Setup: How to get going so we can start developing.
- Modules: Go modules and packages.
- Structure: Structs and Functions.
- Assignments: Variables and loops.
- Http: Build a Simple web service.

## Setup
Installing Go is just as easy to setup as Node or Python. Go to https://golang.org/dl/, download and install. Once that is done you might need to restart so you have Go on your PATH. I'm going to use VSCode sense it is very popular and there is no official Go IDE. We need the Go extension for VSCode and the tools that go with it. Add Go from the extension panel then hit `F1` and run `Go: Install/Update Tools` and select everything that shows up. Running this command now will make things easier so we aren't constantly prompted to install a new tool. The last thing I recommend doing is to enable the new Go language server, It will just make things faster and allow us to use Go Modules (More on those latter). Add these options to your user settings json.

``` json
"go.useLanguageServer": true,
"go.coverageDecorator": {
    "type": "gutter",
    "coveredGutterStyle": "verticalgreen",
    "uncoveredGutterStyle": "blockred",
},
```

Now we are ready to Go!

## Modules
Modules are a fairly new concept to Go, they allow you to easily manage dependencies. I like to think of it as Modules are similar to Javascript packages where the old syntax was closer to python packages. To create a new project make a new folder, I'll call mine `go-app`. Go modules need a name, it is common to use the same name as your git repo, I'll call mine `Devync/go-app`. To create a new module run `go mod init Devync/go-app`. That created a `go.mod` file, this is similar to a package.json in Javascript projects. Now that we have a Module lets talk about Packages.

Everything in Go belongs to a package. Either your code is part of a library that will be imported by another package or it is the special `main` package that will turn into an executable. Lets make a main package, create a folder called `src` and add a file called `main.go`. Every file in go needs to declare what package it belongs too, sense we want our main.go to be an app we can run add `package main` to the top of our file. Lets also add a `func main()` which is the entrypoint for the main package. We should now have something that looks like this.

``` go
package main

func main() {

}
```

This is the bare bones of any Go app. Lets start building a simple web service.

## Structure

Unlike object oriented languages Go does not have classes. Instead it mainly uses structs and functions to organize code. You can attach functions to structs but structs don't have constructors you would create a factory function to create. Lets create a `chat` struct with message and id.

``` go
type chat struct {
    ID int `json:"id"`
    Message string `json:"message"`
    IsRead bool
}
```

That weird string at the end is called a `struct tag`. It's a way to add metadata to properties of a struct. In this case we are adding two json tags so then when we parse this struct to json those two properties will be included (but not IsRead sense it doesn't have one).

Lets start writing an http request function that will return a chat by id. A common function signature in Go for requests is `(w httpResponseWriter, r *http.Request)`. The first part is a helper to write a response body and the second is a pointer to the incoming request. Our function should look something like this.

``` go
func getChatRequest(w http.ResponseWriter, r *http.Request) {


}
```

So our api is going to look like `/chat/2` where 2 will be the id of the chat we want to lookup. To make life easier we are going to use Mux (https://www.gorillatoolkit.org/pkg/mux) which is a http helper. Go is very much about building your own tools rather than using tons of libraries because the standard library in Go is quite good. You might only see a hand full of external packages and they would be simple helper frameworks for the http or sql base libraries.

Before we finish this function lets make some test data and setup our main func to call this one.

## Assignments

We can declare package level vars as well as in functions. The syntax might be different but the idea is the same as other languages.

``` go
var data = []chat{
	chat{
		ID:      0,
		Message: "Chat-1!",
		IsRead:  false,
	},
	chat{
		ID:      1,
		Message: "Chat-2",
		IsRead:  true,
	},
}
```
Lets talk about declaring variables in functions. Both vars in this snippet are equivalent, `:=` is just shorthand for declaration and assignment.

``` go
var1 := 52

var var2 int
var2 = 52
```

In our main function lets setup our mux router which is what we will register our `getChatRequest` function. We will also setup our http.server and register our mux router.

``` go
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/chat/{id}", getChatRequest).Methods("GET")

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
```

One interesting thing is the creation of the `http.Server` struct, we added an ampersand at the beginning. To understand lets breakdown what we are doing. There are two things happening. First we are creating a `http.Server` struct and second we want to be able to do stuff to that struct in a variable called `server`. When we make a variable it can either be a pointer to something in memory or a new copy of the value in memory. Of course we don't want to store this struct in memory twice so we need a pointer! By adding an ampersand at the beginning we get an address to the value (a pointer).

In short most of the time when you create a struct you want a pointer to that struct rather than a copy of it's value.

