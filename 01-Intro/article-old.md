# Introduction to Go
---

Go is a simple but powerful alternative to c++ and other web server languages. Go (or Golang) is a reimagination of c++ but with a focus on current language needs. Go is used a lot in enterprise, Kubernetes, Netflix and Github all use Go. It was written by a group of devs at Google and is Open Source.

## Setup
Installing go is just as easy as node or python. Go to https://golang.org/dl/, download and install. Next we need to setup an IDE. The most popular are VSCode, vim, and GoLand (jetbrains). I'm going to use VSCode sense it is the most popular.

Install the Go extension for VSCode. Next we need to install all the go tools. Hit `F1` and run the command `Go: Install/Update Tools`, Make sure you select everything in the list. The Go extension will normally prompt you to install tools only when you try to run a command that needs it, instead we are just installing everything upfront.

The last configuration we need to make is to enable the Go Language Server. The language server is faster than how the Go extension use to call commands. It also enables us to use Go Modules which is the new way to manage Go dependencies. I also like to change the default test coverage colors. Add this to your user settings in VSCode.

``` json
"go.useLanguageServer": true,
"go.coverageDecorator": {
    "type": "gutter",
    "coveredGutterStyle": "verticalgreen",
    "uncoveredGutterStyle": "blockred",
},
```

Once that is done we are ready to Go :). As I mentioned we are going to be using the new Go Module syntax. The old way was similar to how Python works and wanted your code to be in a certain GOPATH. The module system is similar to Javascript where you have a file that keeps track of dependencies.

Lets start a new project, make a folder called `go-app`. Go modules need a name, it is common to use the repository name. I am going to call mine `Devync/go-app`. To create the module run `go mod init Devync/go-app`. This should create a `go.mod` file. Lets make a `src` folder and add a new file called `main.go`. Lets talk about packages in Go for a second.

Everything in Go is apart of a package. Each Go file will start with declaring what package it belongs too. If you are making something you want to import in another package like a library this would be its name (`package myAwesomeUtils`). The alternative would be if you want a .exe to actually run. There is a special package called main just for that. Every Go project ultimately ends up in a main package. Sense we want an executable lets use `package main`.

``` go
package main

func main() {

}
```

We also declared a `func main`. This function will get called when our app is run similar to `static void Main()` in C#. Lets make a simple web service that can return a chat object by id.

Go doesn't have classes like a normal Object Oriented language would. Most stuff is a struct, func, or a primitive type.

``` go
type chat struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}
```

This should look fairly straight forward except we have this weird string at the end. That string is called a Struct Tag. Basically it's a way of adding metadata to a struct. We want to return a json version of this struct so we need to add a struct tag to the properties we want to be parsed.

Now that we have our struct lets make some test data.

``` go
var data = []chat{
	chat{
		ID:      0,
		Message: "Chat-1!",
	},
	chat{
		ID:      1,
		Message: "Chat-2",
	},
}
```

We are declaring a variable in our package called data that is an array of chat structs. The Go extension is pretty good at guessing what you are doing.

Before we start creating our request function lets edit our main func to start a http server.

``` go
func main() {
    server := &http.Server{
        Addr: ":3000"
    }

    http.HandleFunc("/chat/", getChatRequest)
    log.Fatal(server.ListenAndServe())
}
```

We are using the `"net/http"` package that is included with go. We are making a new instance of the http.Server struct. We added an ampersand at the beginning. To understand lets breakdown what we are doing. There are two things happening. First we are creating a `http.Server` struct and second we want to be able to do stuff to that struct in a variable called `server`. When we make a variable it can either be a pointer to something in memory or a new copy of the value in memory. Of course we don't want to store this struct in memory twice so we need a pointer! By adding an ampersand at the beginning we get an address to the value (a pointer).

In short most of the time when you create a struct you want a pointer to that struct rather than a copy of it's value.

After that we register the path `"/chat/` to the func `getChatRequest`. Finally we start our server by running `server.ListenAndServe()`. That function will return an error if it ever gets one so we pass that error to log.Fatal to exit and log the error message.

That was a lot of words for 7 lines of code...If you are okay with just saying "This is just the magical way to write this" then ignore all that text. I feel like it is important to point out these specifics for those who are curious.

Lets get down to writing that request function (finally).

``` go
func getChatRequest(w http.ResponseWriter, r *http.Request) {

	reqParsed, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/chat/"), 10, 32)

	if err != nil {
		http.Error(w, "Unable to parse request id.", http.StatusInternalServerError)
		return
	}

	reqID := int(reqParsed)
	var res *chat

	for i := range data {
		if data[i].ID == reqID {
            res = &data[i]
            break
		}
	}

	if res == nil {
		http.Error(w, "No item found.", http.StatusNotFound)
		return
	}

	json, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
```

The handler func pattern of `(w http.ResponseWriter, r *http.Request)` is quite common. We have a helper that lets us write data for our new response and a pointer to the request object.

Next we are using some helpers to parse the request url and get the ID. Two weird things is that Go allows multiple return arguments from function. Here we are getting the parsed result and maybe an error if that failed. We also have `:=` which is basically shorthand for declaring and also initializing a variable.

``` go
if err != nil {
    http.Error(w, "Error message.", http.StatusInternalServerError)
    return
}
```

This is a common way of handling errors. Statements don't need parentheses in Go. There is also nil rather than undefined or null.

Afterwards we save the `reqParsed` to reqID which converts from `int64` to just `int`. We create a variable called `res` which is a pointer to a chat struct. We don't use `:=` here because there is no initializer.

We then do a for loop. The only loop in go is a for loop. Here we loop over our test data array and try to find a matching chat and break out.

We do some more error handling, parse our res to json, and finally write our json to the response.