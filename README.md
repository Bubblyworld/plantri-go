# Plantri

This package provides utilities for loading graphs generated by the `plantri`
tool, available for download [here](http://users.cecs.anu.edu.au/~bdm/plantri).


## Usage
```GO
package main

import "github.com/bubblyworld/plantri"

func main() {
  graphs, err := plantri.Load("10.tri")
  if err != nil {
    log.Fatal(err)
  }

  for _, graph := range graphs {
    // ... do stuff
  }
}
```

## License
MIT