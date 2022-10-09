# Section 3: Deeper Into Go

## Khai báo biến (variables)

```go
var card string = "Ace of Spades"
```
Trong đó:
- `var`: từ khóa dùng để khai báo biến.
- `card`: tên biến.
- `string`: kiểu dữ liệu. Lưu ý, `Go` là một ngôn ngữ `Static type`, nên chỉ những dữ liệu có đúng kiểu (ví dụ trên là `string`) thì mới có thể gán.
- `"Ace of Spades"`: giá trị được gán cho biến.

***Một số kiểu dữ liệu thông dụng trong `Go`:***
| Kiểu dữ liệu | Ví dụ |
| :---: | --- |
| `bool` | true, false |
| `string` | "hello, world!", "TienNHM" |
| `int` | `0`, `1000`, `-123` |
| `float64` | `10.000001`, `0.00009`, `-100.001` |

Bên cạnh đó, ta có thể vừa khai báo và gán giá trị theo một cách ngắn gọn hơn như sau:
```go
card := "Ace of Spades"
```
Lưu ý, chúng ta chỉ sử dụng `:=` khi khai báo một biến mới và gán giá trị cho nó. Đoạn code sau đây sẽ báo **lỗi**, do thay vì dùng `=` để gán lại giá trị thì lại dùng `:=`.
```go
card := "Ace of Spades"
card := "Five of Diamonds"
```
Ngoài ra, chúng ta hoàn toàn có thể khai báo trước một biến, sau đó gán giá trị cho biến đó. Ví dụ đoạn chương trình sau đây là hợp lệ:
```go
package main
 
import "fmt"
 
func main() {
  var deckSize int
  deckSize = 52
  fmt.Println(deckSize)
}
```
***Một số lưu ý khi khai báo biến:***
- Có thể khai báo biến bên ngoài một function:

Ví dụ:
```go
package main
 
import "fmt"
 
var deckSize int
 
func main() {
  deckSize = 50
  fmt.Println(deckSize)
}
```

- Không thể gán giá trị cho một biến, nếu biến đó nằm ngoài một function:
> Các ví dụ sau đây là không hợp lệ:

Ví dụ 1:
```go
package main

import "fmt"

var deckSize int
deckSize = 10

func main() {
	fmt.Println(deckSize)
}
```

Ví dụ 2:
```go
package main

import "fmt"

deckSize := 10

func main() {
	fmt.Println(deckSize)
}
```

Tuy nhiên, ta có thể vừa gán giá trị khi khai báo một biến nằm ngoài function.
> Ví dụ sau đây là hợp lệ:

Ví dụ 3:
```go
package main

import "fmt"

var deckSize int = 10

func main() {
	fmt.Println(deckSize)
}
```