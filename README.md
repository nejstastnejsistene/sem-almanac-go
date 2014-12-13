sem-almanac-go
==============

A small library for reading GPS [SEM Almanacs](http://www.navcen.uscg.gov/?pageName=gpsSem).

#### Download the current almanac

```shell
curl 'http://www.navcen.uscg.gov/?pageName=currentAlmanac&format=sem-al3' > current.al3
```

#### Read the almanac

```go
package main

import (
	"io/ioutil"
	"log"

	"github.com/nejstastnejsistene/sem-almanac-go"
)

func main() {
	buf, err := ioutil.ReadFile("current.al3")
	if err != nil {
		log.Fatal(err)
	}
	almanac, err := sem.Unmarshal(buf)
	if err != nil {
		log.Fatal(err)
	}
    // Now simulate the GPS constellation or something!
}
```
