# mee

🍜 A random name generator package for Go.

## Installation

To install `mee`, run the following command:

```bash
go get github.com/somnek/mee
```

## Usage

This example shows how to generate a random name with "-" delimiter:

```go
package main

import (
  "fmt"

  "github.com/somnek/mee"
)

func main() {
  name := mee.Generate("-") // e.g. "noodle-exotic"
}
```

This will generate random in `noun-adjective` format.
