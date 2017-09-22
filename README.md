# Question 1

package main

```
import "fmt"

func main() {
	fmt.Println("Hello World!, こんにちは世界!")
}
```

# Question 2

package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Current Time: ", time.Now())
}

package main

# Question 3

import (
	"strconv"
	"fmt"
)

func main() {
	
	for i := 1; i < 100; i++ {

		output := "";

		if(i % 3 == 0){output = "Fizz";}
		if(i % 5 == 0){output = "Buzz";}

		if(output == ""){output = strconv.Itoa(i)}

		fmt.Println(output)
		
	}
}
