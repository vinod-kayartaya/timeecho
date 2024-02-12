# Custom go package

This is a custom `go` packaget to check the working of `go get` command. To test, follow the instructions below:

- In your current project folder, open a terminal
- Execute the command `go get github.com/vinod-kayartaya/timeecho`
  - This will download and install the packate from `https://github.com/vinod-kayartaya/timeecho` on to your local computer. Typically you can see it in `/Users/your-username/go/pkg/mod/github.com`
- Now you can import the `timeecho` package and call the exported `EchoTime()` function

```go
package main

import (
	"github.com/vinod-kayartaya/timeecho"
)

func main() {
	timeecho.EchoTime()
}
```
