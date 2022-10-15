# Section 3: Deeper Into Go

## Khai báo biến (variables)

<summary>
<details>

### Cách khai báo
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
### Một số lưu ý khi khai báo biến
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

</details>
</summary>


## Function và kiểu dữ liệu trả về

<summary>
<details>

### Cách khai báo function
```go
func NewCard() string {
    return "Five of Diamonds"
}
```
Trong đó:
- `func`: từ khóa.
- `NewCard`: tên function.
- `string`: kiểu dữ liệu trả về của function, có thể có hoặc không. Mọi hàm có kiểu trả về, bắt buộc phải xác định kiểu dữ liệu - đây có thể là một trong số những kiểu dữ liệu cơ bản hoặc do người dùng tự định nghĩa.
> Ví dụ sau là không hợp lệ, vì đây là hàm không trả về dữ liệu:
```go
func getSize() {
    return "30 meters"
}
```
Để sửa lỗi, ta có thể thêm kiểu dữ liệu trả về cho hàm như say:
```go
func getSize() string {
    return "30 meters"
}
```

### Cách gọi hàm
Trong chương trình, một function được gọi (`invoke`) thông qua tên của nó. 
```go
package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Ace of Spades"
}
```
Ở ví dụ trên, hàm `newCard` trả về dữ liệu kiểu `string`. Khi gọi hàm để gán giá trị cho biến `card`, thì biến này sẽ có kiểu dữ liệu `string` (là kiểu trả về của hàm `newCard`).

</details>
</summary>

## Array, Slice và vòng lặp `for`

<summaryz>
<detailsz>

### Phân biệt array với slice
#### `array`
Với ngôn ngữ `go`, `array` là một danh sách chứa dữ liệu có số lượng phần tử cố định (`fix length`). Chính vì vậy, khi khai báo một array, ta phải khai báo số lượng phần tử tối đa của nó. Ví dụ:
```go
cards := [2]string {"A", "B"}
```
Trong đó:
  - `[2]`: khai báo sức chứa tối đa (`capacity`) của một `array`.
  - `string`: kiểu dữ liệu của các phần tử trong `array`.
  - `{"A", "B"}`: liệt kê các phần tử có trong `array`.

Lưu ý: 
- `capacity` và độ dài (`len`) của một `array` không nhất thiết phải bằng nhau, tuy nhiên `len <= capacity`. Khi đó, `array` sẽ được làm đầy bởi các giá trị `zero` tương ứng của kiểu dữ liệu. 

Ví dụ 1:
```go
numbers := [5]int{1, 2}
```
Khi đó, các phần tử thực tế của `numbers` là `1`, `2`, `0`, `0`, `0`.

Ví dụ 2:
```go
cards := [5]string{"A", "B"}
```
Khi đó, các phần tử thực tế của `cards` là `"A"`, `"B"`, `""`, `""`, `""`.
- Ta có thể khai báo `[...]` để tạo một array có sức chứa tối đa bằng số lượng phần tử được liệt kê.
- Để truy xuất giá trị tại một vị trí `index` trong `array`, ta dùng `[]`. Trong đó, phần tử đầu tiên có index là `0` và `index >= 0`.

Ví dụ:
```go
cards := [...]string {"A", "B"}
first := cards[0]
```

Khi đó, `cards` có capacity đúng bằng 2 và giá trị của `first` là `"A"`.

#### `slice`
Đây cũng là một `array` nhưng có thể mở rộng (`grow`) hoặc thu gọn (`shrink`) số phần tử. Khi khai báo, ta không cần chỉ rõ `capacity` của slice.

Ví dụ:
```go
cards := []string {"A", "B"}
```

***Một số built-in function thông dụng cho `slice`:***
- `append()`: thêm (các) phần tử vào cuối.
```go
cards := []string {"A", "B"}
cards = append(cards, "C")
cards = append(cards, "D", "E")
```
Sau khi thực hiện 3 câu lệnh trên, `cards` sẽ gồm các phần tử là: `"A"`, `"B"`, `"C"`, `"D"`, `"E"`.

- `len()`: trả về số lượng phần tử hiện có trong slice.
```go
cards := []string {"A", "B"}
length := len(cards)
```
Giá trị của `length` khi này là `2`.

> Cả array và slide đều chứa các phần tử có cùng kiểu dữ liệu.

</detailsz>
</summaryz>