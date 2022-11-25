# smtpcli
smtp client tool to make sending emails easier.

# Install
``` 
go get github.com/ka1hung/smtpcli
```

# Sample
## Basic Sample
``` go
package main

import (
	"fmt"
	"github.com/ka1hung/smtpcli"
)

func main() {
	// set InsecureMode for testing server 
	smtpcli.InseureMode = true 

	server := "smtp.gmail.com"
	port := 587
	from := "user@gmail.com"
	user := "user@gmail.com"
	passwd := "password"

	sender := smtpcli.NewServer(server,port,from,user,passwd)
	m := smtpcli.NewMessage("Subject", "Body message.")
	m.To = []string{"user1@gmail.com", "user2@gmail.com"}
	fmt.Println(sender.Send(m))
}
```

## Sample with cc bcc
``` go
package main

import (
	"fmt"
	"github.com/ka1hung/smtpcli"
)

func main() {
	server := "smtp.gmail.com"
	port := 587
	from := "user@gmail.com"
	user := "user@gmail.com"
	passwd := "password"

	sender := smtpcli.NewServer(server,port,from,user,passwd)
	m := smtpcli.NewMessage("Subject", "Body message.")
	m.To = []string{"user1@gmail.com", "user2@gmail.com"}
	m.CC = []string{"user3@gmail.com", "user4@gmail.com"}
	m.BCC = []string{"bc@gmail.com"}
	fmt.Println(sender.Send(m))
}
```

## Sample to modify content type
``` go
package main

import (
	"fmt"
	"github.com/ka1hung/smtpcli"
)

func main() {
	server := "smtp.gmail.com"
	port := 587
	from := "user@gmail.com"
	user := "user@gmail.com"
	passwd := "password"

	sender := smtpcli.NewServer(server,port,from,user,passwd)
	m := smtpcli.NewMessage("Subject", "<h1>Body message</h1>")
	m.To = []string{"user1@gmail.com", "user2@gmail.com"}
	m.ContentType = "text/html; charset=utf-8"
	fmt.Println(sender.Send(m))
}
```

## Sample to attach files 
``` go
package main

import (
	"fmt"
	"github.com/ka1hung/smtpcli"
)

func main() {
	server := "smtp.gmail.com"
	port := 587
	from := "user@gmail.com"
	user := "user@gmail.com"
	passwd := "password"

	sender := smtpcli.NewServer(server,port,from,user,passwd)
	m := smtpcli.NewMessage("Subject", "<h1>Body message</h1>")
	m.To = []string{"user1@gmail.com", "user2@gmail.com"}
	m.ContentType = "text/html; charset=utf-8"
	m.AttachFiles([]string{"./123.txt", "./abc.png", "./aaa.svg"})
	fmt.Println(sender.Send(m))
}
```

## Sample to attach files by bytes
``` go
package main

import (
	"fmt"
	"github.com/ka1hung/smtpcli"
)

func main() {
	server := "smtp.gmail.com"
	port := 587
	from := "user@gmail.com"
	user := "user@gmail.com"
	passwd := "password"

	sender := smtpcli.NewServer(server,port,from,user,passwd)
	m := smtpcli.NewMessage("Subject", "<h1>Body message</h1>")
	m.To = []string{"user1@gmail.com", "user2@gmail.com"}
	m.ContentType = "text/html; charset=utf-8"
	m.Attachments = map[string][]byte{"123.txt": []byte("abc")}
	fmt.Println(sender.Send(m))
}
```

improve form this gist https://gist.github.com/douglasmakey/90753ecf37ac10c25873825097f46300

### LICENSE
[MIT](https://github.com/ka1hung/mbserver/blob/master/LICENSE)