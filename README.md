# Go Rubbrband

The official Rubbrband Go client library.

## Installation

Make sure your project is using Go Modules (it will have a go.mod file in its root if it already is):

`go mod init`

Then, reference rubbrband_go in a Go program with `import`:

```
import (
    "github.com/rubbrband/rubbrband_go"
)
```

Run any of the normal go commands (build/install/test). The Go toolchain will resolve and fetch the rubbrband_go module automatically.

### Dashboard

To view your cache entries or modify setting, you can visit [here](https://app.rubbrband.com/)

## Documentation

For a detailed documentation of API, check out [documentations here](https://app.rubbrband.com/docs).

Here are a few snippets to help you get started:

## Set API Key and store a string

```
import (
	"github.com/rubbrband/rubbrband_go"
)

func main() {
	rubbrband_go.SetApiKey("YOUR_KEY_HERE")

    user := rubbrband_go.User{
		Name:       "Jeremy",
		ProfileUrl: "http://google.com",
	}

	result := rubbrband_go.Store("foo", "bar", user)

	fmt.Println("Result,", result)
}
```

## Replay an entry

```
    result := rubbrband_go.Replay("foo")
    fmt.Println("Result,", result)
```

## Delete an entry

```
    deleted := rubbrband_go.Delete("foo:)
```
