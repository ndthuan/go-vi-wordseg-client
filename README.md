# go-vi-wordseg-client
Go client library for the Vietnamese word segmenter service https://github.com/ndthuan/vi-word-segmenter.

Supported versions: Go 1.12 and newer.

# Testing Prerequisites
* Docker on local machine and docker-compose
* GNU make

## How to run tests
```shell script
make test
```

## How to run examples
```shell script
make examples
```

# Usage
```go
package main

import (
	"fmt"
	"github.com/ndthuan/go-vi-wordseg-client/pkg/apiv1"
)

func main() {
	c := apiv1.NewClient("http://segmenterv1:8080")

	tagged, err := c.Tag("Việt Nam tổng tấn công COVID!")

	if err != nil {
		panic(err)
	}

	println("Word-segmented text with tagging:")

	for _, sentence := range tagged.Sentences {
		for _, word := range sentence {
			fmt.Printf("form=%s pos=%s ner=%s dep=%s\n", word.Form, word.Pos, word.Ner, word.Dep)
		}
	}
}
```

Should output:

```text
Word-segmented text with tagging:
form=Việt_Nam pos=Np ner=B-PER dep=sub
form=tổng_tấn_công pos=V ner=O dep=root
form=COVID pos=Ny ner=O dep=dob
form=! pos=CH ner=O dep=punct
```
