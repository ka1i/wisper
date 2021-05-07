package usage

import (
	"fmt"
)

// Usage ...
func Usage() {
	fmt.Println(`Usage: wisper -[hv] 

If you are using it for the first time,
first execute the command "wisper -h" to check usage.

	 ------- < Commands Arguments > -------
optional:
  -h, help		Show this help message. 
  -v, version		Show the app version.
  `)
}
