package parser

import (
	"strconv"
)
import (
	"log"
)
import ("bufio";"io";"strings")
type dfa struct {
  acc []bool
  f []func(rune) int
  id int
}
type family struct {
  a []dfa
  endcase int
}
var a0 [62]dfa
var a []family
func init() {
a = make([]family, 1)
{
var acc [3]bool
var fun [3]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 45: return 1
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    default: return -1
    }
  }
  panic("unreachable")
}
a0[0].acc = acc[:]
a0[0].f = fun[:]
a0[0].id = 0
}
{
var acc [5]bool
var fun [5]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 46: return -1
  case 45: return 1
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    default: return -1
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 45: return -1
  case 46: return -1
  default:
    switch {
    case 48 <= r && r <= 57: return 4
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 45: return -1
  case 46: return 3
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    default: return -1
    }
  }
  panic("unreachable")
}
acc[3] = true
fun[3] = func(r rune) int {
  switch(r) {
  case 45: return -1
  case 46: return -1
  default:
    switch {
    case 48 <= r && r <= 57: return 4
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 45: return -1
  case 46: return -1
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    default: return -1
    }
  }
  panic("unreachable")
}
a0[1].acc = acc[:]
a0[1].f = fun[:]
a0[1].id = 1
}
{
var acc [7]bool
var fun [7]func(rune) int
fun[4] = func(r rune) int {
  switch(r) {
  case 34: return 3
  case 92: return 2
  default:
    switch {
    default: return 4
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 92: return 5
  case 34: return 5
  default:
    switch {
    default: return 5
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 92: return 2
  case 34: return 3
  default:
    switch {
    default: return 6
    }
  }
  panic("unreachable")
}
acc[3] = true
fun[3] = func(r rune) int {
  switch(r) {
  case 92: return -1
  case 34: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 92: return 2
  case 34: return 3
  default:
    switch {
    default: return 6
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 92: return 2
  case 34: return 3
  default:
    switch {
    default: return 4
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 92: return -1
  case 34: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[2].acc = acc[:]
a0[2].f = fun[:]
a0[2].id = 2
}
{
var acc [7]bool
var fun [7]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 39: return 1
  case 92: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 39: return 2
  case 92: return 3
  default:
    switch {
    default: return 4
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 39: return 2
  case 92: return 3
  default:
    switch {
    default: return 6
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 92: return 3
  case 39: return 2
  default:
    switch {
    default: return 6
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 92: return -1
  case 39: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 39: return 2
  case 92: return 3
  default:
    switch {
    default: return 4
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 39: return 5
  case 92: return 5
  default:
    switch {
    default: return 5
    }
  }
  panic("unreachable")
}
a0[3].acc = acc[:]
a0[3].f = fun[:]
a0[3].id = 3
}
{
var acc [13]bool
var fun [13]func(rune) int
fun[4] = func(r rune) int {
  switch(r) {
  case 76: return -1
  case 116: return -1
  case 84: return -1
  case 101: return -1
  case 108: return -1
  case 99: return -1
  case 67: return -1
  case 115: return -1
  case 69: return 5
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 76: return -1
  case 116: return -1
  case 84: return -1
  case 101: return -1
  case 108: return 9
  case 99: return -1
  case 67: return -1
  case 115: return -1
  case 69: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 99: return -1
  case 67: return -1
  case 115: return -1
  case 69: return -1
  case 83: return -1
  case 76: return -1
  case 116: return -1
  case 84: return 7
  case 101: return -1
  case 108: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[11] = func(r rune) int {
  switch(r) {
  case 76: return -1
  case 116: return 12
  case 84: return -1
  case 101: return -1
  case 108: return -1
  case 99: return -1
  case 67: return -1
  case 115: return -1
  case 69: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 76: return -1
  case 116: return -1
  case 84: return -1
  case 101: return -1
  case 108: return -1
  case 99: return -1
  case 67: return 6
  case 115: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[12] = true
fun[12] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 76: return -1
  case 116: return -1
  case 84: return -1
  case 101: return -1
  case 108: return -1
  case 99: return -1
  case 67: return -1
  case 115: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[7] = true
fun[7] = func(r rune) int {
  switch(r) {
  case 76: return -1
  case 116: return -1
  case 84: return -1
  case 101: return -1
  case 108: return -1
  case 99: return -1
  case 67: return -1
  case 115: return -1
  case 69: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[10] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 69: return -1
  case 83: return -1
  case 76: return -1
  case 116: return -1
  case 84: return -1
  case 101: return -1
  case 108: return -1
  case 99: return 11
  case 67: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 76: return -1
  case 116: return -1
  case 84: return -1
  case 101: return 8
  case 108: return -1
  case 99: return -1
  case 67: return -1
  case 115: return -1
  case 69: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 116: return -1
  case 84: return -1
  case 101: return -1
  case 108: return -1
  case 99: return -1
  case 67: return -1
  case 115: return 1
  case 69: return -1
  case 83: return 2
  case 76: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 76: return -1
  case 116: return -1
  case 84: return -1
  case 101: return 10
  case 108: return -1
  case 99: return -1
  case 67: return -1
  case 115: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 67: return -1
  case 115: return -1
  case 69: return 3
  case 83: return -1
  case 76: return -1
  case 116: return -1
  case 84: return -1
  case 101: return -1
  case 108: return -1
  case 99: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 67: return -1
  case 115: return -1
  case 69: return -1
  case 83: return -1
  case 76: return 4
  case 116: return -1
  case 84: return -1
  case 101: return -1
  case 108: return -1
  case 99: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[4].acc = acc[:]
a0[4].f = fun[:]
a0[4].id = 4
}
{
var acc [17]bool
var fun [17]func(rune) int
fun[13] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 115: return -1
  case 100: return -1
  case 78: return 14
  case 99: return -1
  case 105: return -1
  case 110: return -1
  case 67: return -1
  case 83: return -1
  case 68: return -1
  case 84: return -1
  case 116: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 67: return -1
  case 83: return -1
  case 68: return -1
  case 84: return -1
  case 116: return -1
  case 73: return -1
  case 115: return -1
  case 100: return -1
  case 78: return -1
  case 99: return 8
  case 105: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[12] = func(r rune) int {
  switch(r) {
  case 99: return -1
  case 105: return -1
  case 110: return -1
  case 67: return -1
  case 83: return -1
  case 68: return -1
  case 84: return -1
  case 116: return -1
  case 73: return 13
  case 115: return -1
  case 100: return -1
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 116: return -1
  case 73: return -1
  case 115: return -1
  case 100: return -1
  case 78: return -1
  case 99: return -1
  case 105: return 3
  case 110: return -1
  case 67: return -1
  case 83: return -1
  case 68: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[9] = true
fun[9] = func(r rune) int {
  switch(r) {
  case 99: return -1
  case 105: return -1
  case 110: return -1
  case 67: return -1
  case 83: return -1
  case 68: return -1
  case 84: return -1
  case 116: return -1
  case 73: return -1
  case 115: return -1
  case 100: return -1
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[14] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 116: return -1
  case 73: return -1
  case 115: return -1
  case 100: return -1
  case 78: return -1
  case 99: return -1
  case 105: return -1
  case 110: return -1
  case 67: return 15
  case 83: return -1
  case 68: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 68: return -1
  case 84: return -1
  case 116: return 9
  case 73: return -1
  case 115: return -1
  case 100: return -1
  case 78: return -1
  case 99: return -1
  case 105: return -1
  case 110: return -1
  case 67: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[10] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 115: return -1
  case 100: return -1
  case 78: return -1
  case 99: return -1
  case 105: return -1
  case 110: return -1
  case 67: return -1
  case 83: return 11
  case 68: return -1
  case 84: return -1
  case 116: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 78: return -1
  case 99: return -1
  case 105: return -1
  case 110: return 7
  case 67: return -1
  case 83: return -1
  case 68: return -1
  case 84: return -1
  case 116: return -1
  case 73: return -1
  case 115: return -1
  case 100: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[15] = func(r rune) int {
  switch(r) {
  case 78: return -1
  case 99: return -1
  case 105: return -1
  case 110: return -1
  case 67: return -1
  case 83: return -1
  case 68: return -1
  case 84: return 16
  case 116: return -1
  case 73: return -1
  case 115: return -1
  case 100: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 68: return 1
  case 84: return -1
  case 116: return -1
  case 73: return -1
  case 115: return -1
  case 100: return 2
  case 78: return -1
  case 99: return -1
  case 105: return -1
  case 110: return -1
  case 67: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 99: return -1
  case 105: return -1
  case 110: return -1
  case 67: return -1
  case 83: return -1
  case 68: return -1
  case 84: return -1
  case 116: return -1
  case 73: return -1
  case 115: return 4
  case 100: return -1
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 100: return -1
  case 78: return -1
  case 99: return -1
  case 105: return 6
  case 110: return -1
  case 67: return -1
  case 83: return -1
  case 68: return -1
  case 84: return -1
  case 116: return -1
  case 73: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 99: return -1
  case 105: return -1
  case 110: return -1
  case 67: return -1
  case 83: return -1
  case 68: return -1
  case 84: return -1
  case 116: return -1
  case 73: return 10
  case 115: return -1
  case 100: return -1
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[11] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 100: return -1
  case 78: return -1
  case 99: return -1
  case 105: return -1
  case 110: return -1
  case 67: return -1
  case 83: return -1
  case 68: return -1
  case 84: return 12
  case 116: return -1
  case 73: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 115: return -1
  case 100: return -1
  case 78: return -1
  case 99: return -1
  case 105: return -1
  case 110: return -1
  case 67: return -1
  case 83: return -1
  case 68: return -1
  case 84: return -1
  case 116: return 5
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[16] = true
fun[16] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 67: return -1
  case 83: return -1
  case 68: return -1
  case 84: return -1
  case 116: return -1
  case 73: return -1
  case 115: return -1
  case 100: return -1
  case 78: return -1
  case 99: return -1
  case 105: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[5].acc = acc[:]
a0[5].f = fun[:]
a0[5].id = 5
}
{
var acc [7]bool
var fun [7]func(rune) int
fun[5] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 97: return -1
  case 76: return 6
  case 108: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 65: return 1
  case 97: return 2
  case 76: return -1
  case 108: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 97: return -1
  case 76: return 5
  case 108: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 108: return 3
  case 65: return -1
  case 97: return -1
  case 76: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 97: return -1
  case 76: return -1
  case 108: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 108: return -1
  case 65: return -1
  case 97: return -1
  case 76: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 76: return -1
  case 108: return 4
  case 65: return -1
  case 97: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[6].acc = acc[:]
a0[6].f = fun[:]
a0[6].id = 6
}
{
var acc [5]bool
var fun [5]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 65: return 1
  case 97: return 2
  case 115: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 65: return -1
  case 97: return -1
  case 115: return 3
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 97: return -1
  case 115: return -1
  case 83: return 4
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 97: return -1
  case 115: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[3] = true
fun[3] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 97: return -1
  case 115: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[7].acc = acc[:]
a0[7].f = fun[:]
a0[7].id = 7
}
{
var acc [9]bool
var fun [9]func(rune) int
acc[5] = true
fun[5] = func(r rune) int {
  switch(r) {
  case 82: return -1
  case 70: return -1
  case 79: return -1
  case 109: return -1
  case 114: return -1
  case 102: return -1
  case 77: return -1
  case 111: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 70: return -1
  case 79: return 4
  case 109: return -1
  case 114: return -1
  case 102: return -1
  case 77: return -1
  case 111: return -1
  case 82: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 82: return -1
  case 70: return -1
  case 79: return -1
  case 109: return 8
  case 114: return -1
  case 102: return -1
  case 77: return -1
  case 111: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 70: return -1
  case 79: return -1
  case 109: return -1
  case 114: return -1
  case 102: return -1
  case 77: return -1
  case 111: return -1
  case 82: return 3
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 111: return -1
  case 82: return -1
  case 70: return -1
  case 79: return -1
  case 109: return -1
  case 114: return 6
  case 102: return -1
  case 77: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 102: return 1
  case 77: return -1
  case 111: return -1
  case 82: return -1
  case 70: return 2
  case 79: return -1
  case 109: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[8] = true
fun[8] = func(r rune) int {
  switch(r) {
  case 102: return -1
  case 77: return -1
  case 111: return -1
  case 82: return -1
  case 70: return -1
  case 79: return -1
  case 109: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 82: return -1
  case 70: return -1
  case 79: return -1
  case 109: return -1
  case 114: return -1
  case 102: return -1
  case 77: return 5
  case 111: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 102: return -1
  case 77: return -1
  case 111: return 7
  case 82: return -1
  case 70: return -1
  case 79: return -1
  case 109: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[8].acc = acc[:]
a0[8].f = fun[:]
a0[8].id = 8
}
{
var acc [11]bool
var fun [11]func(rune) int
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 119: return -1
  case 114: return -1
  case 82: return -1
  case 87: return -1
  case 101: return -1
  case 104: return -1
  case 69: return -1
  case 72: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 114: return -1
  case 82: return -1
  case 87: return -1
  case 101: return -1
  case 104: return -1
  case 69: return 6
  case 72: return -1
  case 119: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 72: return -1
  case 119: return 1
  case 114: return -1
  case 82: return -1
  case 87: return 2
  case 101: return -1
  case 104: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 104: return 7
  case 69: return -1
  case 72: return -1
  case 119: return -1
  case 114: return -1
  case 82: return -1
  case 87: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 104: return -1
  case 69: return 4
  case 72: return -1
  case 119: return -1
  case 114: return -1
  case 82: return -1
  case 87: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 72: return -1
  case 119: return -1
  case 114: return -1
  case 82: return 5
  case 87: return -1
  case 101: return -1
  case 104: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[10] = true
fun[10] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 72: return -1
  case 119: return -1
  case 114: return -1
  case 82: return -1
  case 87: return -1
  case 101: return -1
  case 104: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 119: return -1
  case 114: return -1
  case 82: return -1
  case 87: return -1
  case 101: return -1
  case 104: return -1
  case 69: return -1
  case 72: return 3
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 72: return -1
  case 119: return -1
  case 114: return 9
  case 82: return -1
  case 87: return -1
  case 101: return -1
  case 104: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 72: return -1
  case 119: return -1
  case 114: return -1
  case 82: return -1
  case 87: return -1
  case 101: return 10
  case 104: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 72: return -1
  case 119: return -1
  case 114: return -1
  case 82: return -1
  case 87: return -1
  case 101: return 8
  case 104: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[9].acc = acc[:]
a0[9].f = fun[:]
a0[9].id = 9
}
{
var acc [11]bool
var fun [11]func(rune) int
fun[8] = func(r rune) int {
  switch(r) {
  case 112: return -1
  case 82: return -1
  case 80: return -1
  case 111: return -1
  case 117: return 9
  case 85: return -1
  case 79: return -1
  case 103: return -1
  case 71: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 112: return -1
  case 82: return -1
  case 80: return 6
  case 111: return -1
  case 117: return -1
  case 85: return -1
  case 79: return -1
  case 103: return -1
  case 71: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[10] = true
fun[10] = func(r rune) int {
  switch(r) {
  case 112: return -1
  case 82: return -1
  case 80: return -1
  case 111: return -1
  case 117: return -1
  case 85: return -1
  case 79: return -1
  case 103: return -1
  case 71: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 112: return -1
  case 82: return -1
  case 80: return -1
  case 111: return -1
  case 117: return -1
  case 85: return -1
  case 79: return -1
  case 103: return -1
  case 71: return -1
  case 114: return 7
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 112: return -1
  case 82: return -1
  case 80: return -1
  case 111: return -1
  case 117: return -1
  case 85: return 5
  case 79: return -1
  case 103: return -1
  case 71: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 112: return 10
  case 82: return -1
  case 80: return -1
  case 111: return -1
  case 117: return -1
  case 85: return -1
  case 79: return -1
  case 103: return -1
  case 71: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 82: return -1
  case 80: return -1
  case 111: return -1
  case 117: return -1
  case 85: return -1
  case 79: return 4
  case 103: return -1
  case 71: return -1
  case 114: return -1
  case 112: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 111: return 8
  case 117: return -1
  case 85: return -1
  case 79: return -1
  case 103: return -1
  case 71: return -1
  case 114: return -1
  case 112: return -1
  case 82: return -1
  case 80: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 111: return -1
  case 117: return -1
  case 85: return -1
  case 79: return -1
  case 103: return 1
  case 71: return 2
  case 114: return -1
  case 112: return -1
  case 82: return -1
  case 80: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 112: return -1
  case 82: return 3
  case 80: return -1
  case 111: return -1
  case 117: return -1
  case 85: return -1
  case 79: return -1
  case 103: return -1
  case 71: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 112: return -1
  case 82: return -1
  case 80: return -1
  case 111: return -1
  case 117: return -1
  case 85: return -1
  case 79: return -1
  case 103: return -1
  case 71: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[10].acc = acc[:]
a0[10].f = fun[:]
a0[10].id = 10
}
{
var acc [5]bool
var fun [5]func(rune) int
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 66: return -1
  case 121: return -1
  case 89: return -1
  case 98: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 66: return 1
  case 121: return -1
  case 89: return -1
  case 98: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 66: return -1
  case 121: return -1
  case 89: return 4
  case 98: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[3] = true
fun[3] = func(r rune) int {
  switch(r) {
  case 89: return -1
  case 98: return -1
  case 66: return -1
  case 121: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 89: return -1
  case 98: return -1
  case 66: return -1
  case 121: return 3
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[11].acc = acc[:]
a0[11].f = fun[:]
a0[11].id = 11
}
{
var acc [13]bool
var fun [13]func(rune) int
fun[2] = func(r rune) int {
  switch(r) {
  case 78: return -1
  case 103: return -1
  case 72: return -1
  case 71: return -1
  case 118: return -1
  case 97: return -1
  case 110: return -1
  case 73: return -1
  case 104: return -1
  case 86: return -1
  case 65: return 3
  case 105: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[11] = func(r rune) int {
  switch(r) {
  case 105: return -1
  case 78: return -1
  case 103: return 12
  case 72: return -1
  case 71: return -1
  case 118: return -1
  case 97: return -1
  case 110: return -1
  case 73: return -1
  case 104: return -1
  case 86: return -1
  case 65: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 104: return -1
  case 86: return -1
  case 65: return -1
  case 105: return -1
  case 78: return -1
  case 103: return -1
  case 72: return -1
  case 71: return -1
  case 118: return 9
  case 97: return -1
  case 110: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 72: return -1
  case 71: return -1
  case 118: return -1
  case 97: return -1
  case 110: return -1
  case 73: return -1
  case 104: return -1
  case 86: return -1
  case 65: return -1
  case 105: return 10
  case 78: return -1
  case 103: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[12] = true
fun[12] = func(r rune) int {
  switch(r) {
  case 71: return -1
  case 118: return -1
  case 97: return -1
  case 110: return -1
  case 73: return -1
  case 104: return -1
  case 86: return -1
  case 65: return -1
  case 105: return -1
  case 78: return -1
  case 103: return -1
  case 72: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[10] = func(r rune) int {
  switch(r) {
  case 78: return -1
  case 103: return -1
  case 72: return -1
  case 71: return -1
  case 118: return -1
  case 97: return -1
  case 110: return 11
  case 73: return -1
  case 104: return -1
  case 86: return -1
  case 65: return -1
  case 105: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 104: return -1
  case 86: return -1
  case 65: return -1
  case 105: return -1
  case 78: return -1
  case 103: return -1
  case 72: return -1
  case 71: return 7
  case 118: return -1
  case 97: return -1
  case 110: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 78: return -1
  case 103: return -1
  case 72: return -1
  case 71: return -1
  case 118: return -1
  case 97: return -1
  case 110: return -1
  case 73: return 5
  case 104: return -1
  case 86: return -1
  case 65: return -1
  case 105: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 105: return -1
  case 78: return 6
  case 103: return -1
  case 72: return -1
  case 71: return -1
  case 118: return -1
  case 97: return -1
  case 110: return -1
  case 73: return -1
  case 104: return -1
  case 86: return -1
  case 65: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[7] = true
fun[7] = func(r rune) int {
  switch(r) {
  case 118: return -1
  case 97: return -1
  case 110: return -1
  case 73: return -1
  case 104: return -1
  case 86: return -1
  case 65: return -1
  case 105: return -1
  case 78: return -1
  case 103: return -1
  case 72: return -1
  case 71: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 97: return -1
  case 110: return -1
  case 73: return -1
  case 104: return 1
  case 86: return -1
  case 65: return -1
  case 105: return -1
  case 78: return -1
  case 103: return -1
  case 72: return 2
  case 71: return -1
  case 118: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 86: return -1
  case 65: return -1
  case 105: return -1
  case 78: return -1
  case 103: return -1
  case 72: return -1
  case 71: return -1
  case 118: return -1
  case 97: return 8
  case 110: return -1
  case 73: return -1
  case 104: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 105: return -1
  case 78: return -1
  case 103: return -1
  case 72: return -1
  case 71: return -1
  case 118: return -1
  case 97: return -1
  case 110: return -1
  case 73: return -1
  case 104: return -1
  case 86: return 4
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[12].acc = acc[:]
a0[12].f = fun[:]
a0[12].id = 12
}
{
var acc [15]bool
var fun [15]func(rune) int
fun[9] = func(r rune) int {
  switch(r) {
  case 116: return -1
  case 101: return -1
  case 76: return -1
  case 84: return -1
  case 108: return -1
  case 110: return -1
  case 65: return 10
  case 69: return -1
  case 70: return -1
  case 97: return -1
  case 78: return -1
  case 102: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 76: return -1
  case 84: return -1
  case 108: return -1
  case 110: return -1
  case 65: return -1
  case 69: return -1
  case 70: return -1
  case 97: return -1
  case 78: return -1
  case 102: return -1
  case 116: return 6
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 76: return -1
  case 84: return -1
  case 108: return 3
  case 110: return -1
  case 65: return -1
  case 69: return -1
  case 70: return -1
  case 97: return -1
  case 78: return -1
  case 102: return -1
  case 116: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 69: return -1
  case 70: return -1
  case 97: return -1
  case 78: return -1
  case 102: return -1
  case 116: return 5
  case 101: return -1
  case 76: return -1
  case 84: return -1
  case 108: return -1
  case 110: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[13] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 76: return -1
  case 84: return -1
  case 108: return -1
  case 110: return -1
  case 65: return -1
  case 69: return -1
  case 70: return -1
  case 97: return -1
  case 78: return 14
  case 102: return -1
  case 116: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[14] = true
fun[14] = func(r rune) int {
  switch(r) {
  case 70: return -1
  case 97: return -1
  case 78: return -1
  case 102: return -1
  case 116: return -1
  case 101: return -1
  case 76: return -1
  case 84: return -1
  case 108: return -1
  case 110: return -1
  case 65: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[12] = func(r rune) int {
  switch(r) {
  case 78: return -1
  case 102: return -1
  case 116: return -1
  case 101: return -1
  case 76: return -1
  case 84: return -1
  case 108: return -1
  case 110: return -1
  case 65: return -1
  case 69: return 13
  case 70: return -1
  case 97: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[10] = func(r rune) int {
  switch(r) {
  case 108: return -1
  case 110: return -1
  case 65: return -1
  case 69: return -1
  case 70: return -1
  case 97: return -1
  case 78: return -1
  case 102: return -1
  case 116: return -1
  case 101: return -1
  case 76: return -1
  case 84: return 11
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 69: return -1
  case 70: return 1
  case 97: return -1
  case 78: return -1
  case 102: return 2
  case 116: return -1
  case 101: return -1
  case 76: return -1
  case 84: return -1
  case 108: return -1
  case 110: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[11] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 65: return -1
  case 69: return -1
  case 70: return -1
  case 97: return -1
  case 78: return -1
  case 102: return -1
  case 116: return -1
  case 101: return -1
  case 76: return -1
  case 84: return 12
  case 108: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 108: return -1
  case 110: return -1
  case 65: return -1
  case 69: return -1
  case 70: return -1
  case 97: return 4
  case 78: return -1
  case 102: return -1
  case 116: return -1
  case 101: return -1
  case 76: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[8] = true
fun[8] = func(r rune) int {
  switch(r) {
  case 116: return -1
  case 101: return -1
  case 76: return -1
  case 84: return -1
  case 108: return -1
  case 110: return -1
  case 65: return -1
  case 69: return -1
  case 70: return -1
  case 97: return -1
  case 78: return -1
  case 102: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 108: return -1
  case 110: return 8
  case 65: return -1
  case 69: return -1
  case 70: return -1
  case 97: return -1
  case 78: return -1
  case 102: return -1
  case 116: return -1
  case 101: return -1
  case 76: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 108: return -1
  case 110: return -1
  case 65: return -1
  case 69: return -1
  case 70: return -1
  case 97: return -1
  case 78: return -1
  case 102: return -1
  case 116: return -1
  case 101: return -1
  case 76: return 9
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 108: return -1
  case 110: return -1
  case 65: return -1
  case 69: return -1
  case 70: return -1
  case 97: return -1
  case 78: return -1
  case 102: return -1
  case 116: return -1
  case 101: return 7
  case 76: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[13].acc = acc[:]
a0[13].f = fun[:]
a0[13].id = 13
}
{
var acc [9]bool
var fun [9]func(rune) int
fun[7] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 99: return -1
  case 65: return -1
  case 97: return -1
  case 72: return 8
  case 67: return -1
  case 104: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[8] = true
fun[8] = func(r rune) int {
  switch(r) {
  case 99: return -1
  case 65: return -1
  case 97: return -1
  case 72: return -1
  case 67: return -1
  case 104: return -1
  case 69: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 97: return 3
  case 72: return -1
  case 67: return -1
  case 104: return -1
  case 69: return -1
  case 101: return -1
  case 99: return -1
  case 65: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[5] = true
fun[5] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 99: return -1
  case 65: return -1
  case 97: return -1
  case 72: return -1
  case 67: return -1
  case 104: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 99: return -1
  case 65: return -1
  case 97: return -1
  case 72: return -1
  case 67: return -1
  case 104: return -1
  case 69: return 1
  case 101: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 104: return 5
  case 69: return -1
  case 101: return -1
  case 99: return -1
  case 65: return -1
  case 97: return -1
  case 72: return -1
  case 67: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 99: return -1
  case 65: return 6
  case 97: return -1
  case 72: return -1
  case 67: return -1
  case 104: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 99: return 4
  case 65: return -1
  case 97: return -1
  case 72: return -1
  case 67: return -1
  case 104: return -1
  case 69: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 99: return -1
  case 65: return -1
  case 97: return -1
  case 72: return -1
  case 67: return 7
  case 104: return -1
  case 69: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[14].acc = acc[:]
a0[14].f = fun[:]
a0[14].id = 14
}
{
var acc [11]bool
var fun [11]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 85: return -1
  case 117: return -1
  case 79: return -1
  case 110: return -1
  case 105: return -1
  case 111: return -1
  case 73: return -1
  case 78: return 7
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 85: return -1
  case 117: return -1
  case 79: return -1
  case 110: return -1
  case 105: return 4
  case 111: return -1
  case 73: return -1
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[10] = true
fun[10] = func(r rune) int {
  switch(r) {
  case 85: return -1
  case 117: return -1
  case 79: return -1
  case 110: return -1
  case 105: return -1
  case 111: return -1
  case 73: return -1
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 85: return -1
  case 117: return -1
  case 79: return -1
  case 110: return 3
  case 105: return -1
  case 111: return -1
  case 73: return -1
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 85: return 1
  case 117: return 2
  case 79: return -1
  case 110: return -1
  case 105: return -1
  case 111: return -1
  case 73: return -1
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 105: return -1
  case 111: return -1
  case 73: return 8
  case 78: return -1
  case 85: return -1
  case 117: return -1
  case 79: return -1
  case 110: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 78: return -1
  case 85: return -1
  case 117: return -1
  case 79: return 9
  case 110: return -1
  case 105: return -1
  case 111: return -1
  case 73: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 78: return -1
  case 85: return -1
  case 117: return -1
  case 79: return -1
  case 110: return 6
  case 105: return -1
  case 111: return -1
  case 73: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 111: return -1
  case 73: return -1
  case 78: return -1
  case 85: return -1
  case 117: return -1
  case 79: return -1
  case 110: return -1
  case 105: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 111: return -1
  case 73: return -1
  case 78: return 10
  case 85: return -1
  case 117: return -1
  case 79: return -1
  case 110: return -1
  case 105: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 105: return -1
  case 111: return 5
  case 73: return -1
  case 78: return -1
  case 85: return -1
  case 117: return -1
  case 79: return -1
  case 110: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[15].acc = acc[:]
a0[15].f = fun[:]
a0[15].id = 15
}
{
var acc [19]bool
var fun [19]func(rune) int
fun[16] = func(r rune) int {
  switch(r) {
  case 78: return -1
  case 67: return 17
  case 82: return -1
  case 83: return -1
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  case 84: return -1
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[11] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  case 84: return 12
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 116: return -1
  case 115: return 7
  case 105: return -1
  case 114: return -1
  case 84: return -1
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[15] = func(r rune) int {
  switch(r) {
  case 69: return 16
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  case 84: return -1
  case 73: return -1
  case 99: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  case 110: return 3
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[14] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return 15
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[12] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  case 84: return -1
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return 13
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 73: return -1
  case 99: return -1
  case 101: return 5
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[10] = true
fun[10] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 73: return 1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return 2
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[18] = true
fun[18] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  case 84: return -1
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 116: return 10
  case 115: return -1
  case 105: return -1
  case 114: return -1
  case 84: return -1
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return 6
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  case 84: return -1
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  case 78: return 11
  case 67: return -1
  case 82: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[13] = func(r rune) int {
  switch(r) {
  case 78: return -1
  case 67: return -1
  case 82: return 14
  case 83: return -1
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  case 84: return -1
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  case 110: return -1
  case 116: return 4
  case 115: return -1
  case 105: return -1
  case 114: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  case 84: return -1
  case 73: return -1
  case 99: return -1
  case 101: return 8
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 99: return 9
  case 101: return -1
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[17] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 116: return -1
  case 115: return -1
  case 105: return -1
  case 114: return -1
  case 84: return 18
  case 73: return -1
  case 99: return -1
  case 101: return -1
  case 69: return -1
  case 78: return -1
  case 67: return -1
  case 82: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[16].acc = acc[:]
a0[16].f = fun[:]
a0[16].id = 16
}
{
var acc [13]bool
var fun [13]func(rune) int
fun[9] = func(r rune) int {
  switch(r) {
  case 112: return -1
  case 99: return -1
  case 120: return -1
  case 69: return 10
  case 88: return -1
  case 84: return -1
  case 101: return -1
  case 116: return -1
  case 67: return -1
  case 80: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 80: return -1
  case 112: return -1
  case 99: return -1
  case 120: return -1
  case 69: return 1
  case 88: return -1
  case 84: return -1
  case 101: return 2
  case 116: return -1
  case 67: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[12] = true
fun[12] = func(r rune) int {
  switch(r) {
  case 80: return -1
  case 112: return -1
  case 99: return -1
  case 120: return -1
  case 69: return -1
  case 88: return -1
  case 84: return -1
  case 101: return -1
  case 116: return -1
  case 67: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 101: return -1
  case 116: return -1
  case 67: return -1
  case 80: return -1
  case 112: return -1
  case 99: return 4
  case 120: return -1
  case 69: return -1
  case 88: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 88: return 8
  case 84: return -1
  case 101: return -1
  case 116: return -1
  case 67: return -1
  case 80: return -1
  case 112: return -1
  case 99: return -1
  case 120: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[7] = true
fun[7] = func(r rune) int {
  switch(r) {
  case 88: return -1
  case 84: return -1
  case 101: return -1
  case 116: return -1
  case 67: return -1
  case 80: return -1
  case 112: return -1
  case 99: return -1
  case 120: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 101: return 5
  case 116: return -1
  case 67: return -1
  case 80: return -1
  case 112: return -1
  case 99: return -1
  case 120: return -1
  case 69: return -1
  case 88: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[10] = func(r rune) int {
  switch(r) {
  case 80: return 11
  case 112: return -1
  case 99: return -1
  case 120: return -1
  case 69: return -1
  case 88: return -1
  case 84: return -1
  case 101: return -1
  case 116: return -1
  case 67: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 120: return -1
  case 69: return -1
  case 88: return -1
  case 84: return -1
  case 101: return -1
  case 116: return 7
  case 67: return -1
  case 80: return -1
  case 112: return -1
  case 99: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 112: return 6
  case 99: return -1
  case 120: return -1
  case 69: return -1
  case 88: return -1
  case 84: return -1
  case 101: return -1
  case 116: return -1
  case 67: return -1
  case 80: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[11] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 116: return -1
  case 67: return -1
  case 80: return -1
  case 112: return -1
  case 99: return -1
  case 120: return -1
  case 69: return -1
  case 88: return -1
  case 84: return 12
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 112: return -1
  case 99: return -1
  case 120: return -1
  case 69: return -1
  case 88: return -1
  case 84: return -1
  case 101: return -1
  case 116: return -1
  case 67: return 9
  case 80: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 101: return -1
  case 116: return -1
  case 67: return -1
  case 80: return -1
  case 112: return -1
  case 99: return -1
  case 120: return 3
  case 69: return -1
  case 88: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[17].acc = acc[:]
a0[17].f = fun[:]
a0[17].id = 17
}
{
var acc [11]bool
var fun [11]func(rune) int
acc[10] = true
fun[10] = func(r rune) int {
  switch(r) {
  case 68: return -1
  case 111: return -1
  case 69: return -1
  case 114: return -1
  case 79: return -1
  case 100: return -1
  case 82: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 68: return -1
  case 111: return -1
  case 69: return 5
  case 114: return -1
  case 79: return -1
  case 100: return -1
  case 82: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 68: return -1
  case 111: return -1
  case 69: return -1
  case 114: return -1
  case 79: return -1
  case 100: return -1
  case 82: return 3
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 68: return -1
  case 111: return -1
  case 69: return -1
  case 114: return -1
  case 79: return -1
  case 100: return 8
  case 82: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 114: return -1
  case 79: return -1
  case 100: return -1
  case 82: return 6
  case 101: return -1
  case 68: return -1
  case 111: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 68: return -1
  case 111: return 1
  case 69: return -1
  case 114: return -1
  case 79: return 2
  case 100: return -1
  case 82: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 79: return -1
  case 100: return -1
  case 82: return -1
  case 101: return 9
  case 68: return -1
  case 111: return -1
  case 69: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 114: return -1
  case 79: return -1
  case 100: return -1
  case 82: return -1
  case 101: return -1
  case 68: return 4
  case 111: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 68: return -1
  case 111: return -1
  case 69: return -1
  case 114: return -1
  case 79: return -1
  case 100: return -1
  case 82: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 68: return -1
  case 111: return -1
  case 69: return -1
  case 114: return 7
  case 79: return -1
  case 100: return -1
  case 82: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 68: return -1
  case 111: return -1
  case 69: return -1
  case 114: return 10
  case 79: return -1
  case 100: return -1
  case 82: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[18].acc = acc[:]
a0[18].f = fun[:]
a0[18].id = 18
}
{
var acc [11]bool
var fun [11]func(rune) int
fun[9] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 105: return -1
  case 73: return -1
  case 108: return -1
  case 76: return -1
  case 116: return 10
  case 109: return -1
  case 77: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 105: return -1
  case 73: return -1
  case 108: return -1
  case 76: return -1
  case 116: return -1
  case 109: return 8
  case 77: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 105: return -1
  case 73: return -1
  case 108: return -1
  case 76: return -1
  case 116: return -1
  case 109: return -1
  case 77: return 4
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 105: return -1
  case 73: return 3
  case 108: return -1
  case 76: return -1
  case 116: return -1
  case 109: return -1
  case 77: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 105: return -1
  case 73: return -1
  case 108: return -1
  case 76: return -1
  case 116: return -1
  case 109: return -1
  case 77: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 109: return -1
  case 77: return -1
  case 84: return -1
  case 105: return 7
  case 73: return -1
  case 108: return -1
  case 76: return -1
  case 116: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[10] = true
fun[10] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 105: return -1
  case 73: return -1
  case 108: return -1
  case 76: return -1
  case 116: return -1
  case 109: return -1
  case 77: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 105: return -1
  case 73: return -1
  case 108: return -1
  case 76: return -1
  case 116: return -1
  case 109: return -1
  case 77: return -1
  case 84: return 6
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 105: return 9
  case 73: return -1
  case 108: return -1
  case 76: return -1
  case 116: return -1
  case 109: return -1
  case 77: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 105: return -1
  case 73: return 5
  case 108: return -1
  case 76: return -1
  case 116: return -1
  case 109: return -1
  case 77: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 77: return -1
  case 84: return -1
  case 105: return -1
  case 73: return -1
  case 108: return 1
  case 76: return 2
  case 116: return -1
  case 109: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[19].acc = acc[:]
a0[19].f = fun[:]
a0[19].id = 19
}
{
var acc [13]bool
var fun [13]func(rune) int
fun[2] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 101: return -1
  case 79: return -1
  case 115: return -1
  case 111: return -1
  case 102: return 3
  case 70: return -1
  case 83: return -1
  case 116: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 101: return -1
  case 79: return -1
  case 115: return -1
  case 111: return -1
  case 102: return -1
  case 70: return -1
  case 83: return 10
  case 116: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 116: return -1
  case 84: return -1
  case 69: return -1
  case 101: return -1
  case 79: return -1
  case 115: return -1
  case 111: return -1
  case 102: return -1
  case 70: return 8
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 102: return -1
  case 70: return -1
  case 83: return -1
  case 116: return -1
  case 84: return -1
  case 69: return -1
  case 101: return -1
  case 79: return -1
  case 115: return 5
  case 111: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[12] = true
fun[12] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 101: return -1
  case 79: return -1
  case 115: return -1
  case 111: return -1
  case 102: return -1
  case 70: return -1
  case 83: return -1
  case 116: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[11] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 79: return -1
  case 115: return -1
  case 111: return -1
  case 102: return -1
  case 70: return -1
  case 83: return -1
  case 116: return -1
  case 84: return 12
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 79: return 1
  case 115: return -1
  case 111: return 2
  case 102: return -1
  case 70: return -1
  case 83: return -1
  case 116: return -1
  case 84: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 102: return -1
  case 70: return -1
  case 83: return -1
  case 116: return -1
  case 84: return -1
  case 69: return -1
  case 101: return 6
  case 79: return -1
  case 115: return -1
  case 111: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 116: return -1
  case 84: return -1
  case 69: return -1
  case 101: return -1
  case 79: return -1
  case 115: return -1
  case 111: return -1
  case 102: return 4
  case 70: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 116: return 7
  case 84: return -1
  case 69: return -1
  case 101: return -1
  case 79: return -1
  case 115: return -1
  case 111: return -1
  case 102: return -1
  case 70: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 116: return -1
  case 84: return -1
  case 69: return -1
  case 101: return -1
  case 79: return -1
  case 115: return -1
  case 111: return -1
  case 102: return -1
  case 70: return 9
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[10] = func(r rune) int {
  switch(r) {
  case 69: return 11
  case 101: return -1
  case 79: return -1
  case 115: return -1
  case 111: return -1
  case 102: return -1
  case 70: return -1
  case 83: return -1
  case 116: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[7] = true
fun[7] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 101: return -1
  case 79: return -1
  case 115: return -1
  case 111: return -1
  case 102: return -1
  case 70: return -1
  case 83: return -1
  case 116: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[20].acc = acc[:]
a0[20].f = fun[:]
a0[20].id = 20
}
{
var acc [7]bool
var fun [7]func(rune) int
fun[3] = func(r rune) int {
  switch(r) {
  case 97: return -1
  case 115: return -1
  case 99: return 4
  case 67: return -1
  case 83: return -1
  case 65: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 99: return -1
  case 67: return -1
  case 83: return -1
  case 65: return 1
  case 97: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 99: return -1
  case 67: return 6
  case 83: return -1
  case 65: return -1
  case 97: return -1
  case 115: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 67: return -1
  case 83: return 5
  case 65: return -1
  case 97: return -1
  case 115: return -1
  case 99: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 99: return -1
  case 67: return -1
  case 83: return -1
  case 65: return -1
  case 97: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 65: return -1
  case 97: return -1
  case 115: return 3
  case 99: return -1
  case 67: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 99: return -1
  case 67: return -1
  case 83: return -1
  case 65: return -1
  case 97: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[21].acc = acc[:]
a0[21].f = fun[:]
a0[21].id = 21
}
{
var acc [9]bool
var fun [9]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 68: return -1
  case 67: return -1
  case 100: return -1
  case 115: return -1
  case 69: return 6
  case 99: return -1
  case 101: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[8] = true
fun[8] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 69: return -1
  case 99: return -1
  case 101: return -1
  case 83: return -1
  case 68: return -1
  case 67: return -1
  case 100: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[5] = true
fun[5] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 99: return -1
  case 101: return -1
  case 83: return -1
  case 68: return -1
  case 67: return -1
  case 100: return -1
  case 115: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 68: return -1
  case 67: return -1
  case 100: return -1
  case 115: return -1
  case 69: return -1
  case 99: return -1
  case 101: return 3
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 83: return 7
  case 68: return -1
  case 67: return -1
  case 100: return -1
  case 115: return -1
  case 69: return -1
  case 99: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 99: return -1
  case 101: return -1
  case 83: return -1
  case 68: return -1
  case 67: return 8
  case 100: return -1
  case 115: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 115: return 4
  case 69: return -1
  case 99: return -1
  case 101: return -1
  case 83: return -1
  case 68: return -1
  case 67: return -1
  case 100: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 69: return -1
  case 99: return 5
  case 101: return -1
  case 83: return -1
  case 68: return -1
  case 67: return -1
  case 100: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 68: return 1
  case 67: return -1
  case 100: return 2
  case 115: return -1
  case 69: return -1
  case 99: return -1
  case 101: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[22].acc = acc[:]
a0[22].f = fun[:]
a0[22].id = 22
}
{
var acc [9]bool
var fun [9]func(rune) int
acc[8] = true
fun[8] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 85: return -1
  case 116: return -1
  case 117: return -1
  case 114: return -1
  case 82: return -1
  case 101: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 69: return 8
  case 84: return -1
  case 85: return -1
  case 116: return -1
  case 117: return -1
  case 114: return -1
  case 82: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[5] = true
fun[5] = func(r rune) int {
  switch(r) {
  case 82: return -1
  case 101: return -1
  case 69: return -1
  case 84: return -1
  case 85: return -1
  case 116: return -1
  case 117: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 84: return 1
  case 85: return -1
  case 116: return 2
  case 117: return -1
  case 114: return -1
  case 82: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 82: return -1
  case 101: return -1
  case 69: return -1
  case 84: return -1
  case 85: return -1
  case 116: return -1
  case 117: return -1
  case 114: return 3
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 85: return -1
  case 116: return -1
  case 117: return -1
  case 114: return -1
  case 82: return 6
  case 101: return -1
  case 69: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 114: return -1
  case 82: return -1
  case 101: return -1
  case 69: return -1
  case 84: return -1
  case 85: return -1
  case 116: return -1
  case 117: return 4
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 84: return -1
  case 85: return 7
  case 116: return -1
  case 117: return -1
  case 114: return -1
  case 82: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 85: return -1
  case 116: return -1
  case 117: return -1
  case 114: return -1
  case 82: return -1
  case 101: return 5
  case 69: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[23].acc = acc[:]
a0[23].f = fun[:]
a0[23].id = 23
}
{
var acc [11]bool
var fun [11]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 101: return -1
  case 102: return 1
  case 69: return -1
  case 115: return -1
  case 97: return -1
  case 83: return -1
  case 108: return -1
  case 76: return -1
  case 70: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 108: return 8
  case 76: return -1
  case 70: return -1
  case 65: return -1
  case 101: return -1
  case 102: return -1
  case 69: return -1
  case 115: return -1
  case 97: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[10] = true
fun[10] = func(r rune) int {
  switch(r) {
  case 108: return -1
  case 76: return -1
  case 70: return -1
  case 65: return -1
  case 101: return -1
  case 102: return -1
  case 69: return -1
  case 115: return -1
  case 97: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 102: return -1
  case 69: return -1
  case 115: return -1
  case 97: return -1
  case 83: return 5
  case 108: return -1
  case 76: return -1
  case 70: return -1
  case 65: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 97: return -1
  case 83: return -1
  case 108: return -1
  case 76: return -1
  case 70: return -1
  case 65: return -1
  case 101: return -1
  case 102: return -1
  case 69: return 6
  case 115: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 101: return -1
  case 102: return -1
  case 69: return -1
  case 115: return 9
  case 97: return -1
  case 83: return -1
  case 108: return -1
  case 76: return -1
  case 70: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 97: return -1
  case 83: return -1
  case 108: return -1
  case 76: return 4
  case 70: return -1
  case 65: return -1
  case 101: return -1
  case 102: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 101: return 10
  case 102: return -1
  case 69: return -1
  case 115: return -1
  case 97: return -1
  case 83: return -1
  case 108: return -1
  case 76: return -1
  case 70: return -1
  case 65: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 108: return -1
  case 76: return -1
  case 70: return -1
  case 65: return -1
  case 101: return -1
  case 102: return -1
  case 69: return -1
  case 115: return -1
  case 97: return 7
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 70: return -1
  case 65: return -1
  case 101: return -1
  case 102: return -1
  case 69: return -1
  case 115: return -1
  case 97: return -1
  case 83: return -1
  case 108: return -1
  case 76: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 65: return 3
  case 101: return -1
  case 102: return -1
  case 69: return -1
  case 115: return -1
  case 97: return -1
  case 83: return -1
  case 108: return -1
  case 76: return -1
  case 70: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[24].acc = acc[:]
a0[24].f = fun[:]
a0[24].id = 24
}
{
var acc [7]bool
var fun [7]func(rune) int
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 97: return -1
  case 77: return -1
  case 109: return -1
  case 88: return -1
  case 65: return -1
  case 120: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 97: return -1
  case 77: return -1
  case 109: return -1
  case 88: return -1
  case 65: return 5
  case 120: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 120: return -1
  case 97: return 3
  case 77: return -1
  case 109: return -1
  case 88: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 77: return -1
  case 109: return -1
  case 88: return 6
  case 65: return -1
  case 120: return -1
  case 97: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 120: return -1
  case 97: return -1
  case 77: return 1
  case 109: return 2
  case 88: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 120: return -1
  case 97: return -1
  case 77: return -1
  case 109: return -1
  case 88: return -1
  case 65: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 97: return -1
  case 77: return -1
  case 109: return -1
  case 88: return -1
  case 65: return -1
  case 120: return 4
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[25].acc = acc[:]
a0[25].f = fun[:]
a0[25].id = 25
}
{
var acc [7]bool
var fun [7]func(rune) int
fun[2] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 109: return -1
  case 77: return -1
  case 105: return -1
  case 73: return 3
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 78: return -1
  case 110: return -1
  case 109: return -1
  case 77: return -1
  case 105: return 5
  case 73: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 109: return -1
  case 77: return -1
  case 105: return -1
  case 73: return -1
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 109: return -1
  case 77: return -1
  case 105: return -1
  case 73: return -1
  case 78: return -1
  case 110: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 105: return -1
  case 73: return -1
  case 78: return -1
  case 110: return 6
  case 109: return -1
  case 77: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 109: return 1
  case 77: return 2
  case 105: return -1
  case 73: return -1
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 78: return 4
  case 110: return -1
  case 109: return -1
  case 77: return -1
  case 105: return -1
  case 73: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[26].acc = acc[:]
a0[26].f = fun[:]
a0[26].id = 26
}
{
var acc [7]bool
var fun [7]func(rune) int
fun[5] = func(r rune) int {
  switch(r) {
  case 103: return 6
  case 71: return -1
  case 97: return -1
  case 86: return -1
  case 118: return -1
  case 65: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 71: return -1
  case 97: return -1
  case 86: return -1
  case 118: return -1
  case 65: return -1
  case 103: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 118: return -1
  case 65: return -1
  case 103: return -1
  case 71: return -1
  case 97: return -1
  case 86: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 103: return -1
  case 71: return -1
  case 97: return -1
  case 86: return -1
  case 118: return 5
  case 65: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 103: return -1
  case 71: return -1
  case 97: return 1
  case 86: return -1
  case 118: return -1
  case 65: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 71: return -1
  case 97: return -1
  case 86: return 3
  case 118: return -1
  case 65: return -1
  case 103: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 71: return 4
  case 97: return -1
  case 86: return -1
  case 118: return -1
  case 65: return -1
  case 103: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[27].acc = acc[:]
a0[27].f = fun[:]
a0[27].id = 27
}
{
var acc [11]bool
var fun [11]func(rune) int
acc[10] = true
fun[10] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 78: return -1
  case 117: return -1
  case 99: return -1
  case 111: return -1
  case 85: return -1
  case 110: return -1
  case 67: return -1
  case 79: return -1
  case 116: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 110: return 9
  case 67: return -1
  case 79: return -1
  case 116: return -1
  case 84: return -1
  case 78: return -1
  case 117: return -1
  case 99: return -1
  case 111: return -1
  case 85: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 79: return -1
  case 116: return -1
  case 84: return -1
  case 78: return -1
  case 117: return -1
  case 99: return -1
  case 111: return 7
  case 85: return -1
  case 110: return -1
  case 67: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 78: return -1
  case 117: return -1
  case 99: return -1
  case 111: return -1
  case 85: return -1
  case 110: return -1
  case 67: return -1
  case 79: return -1
  case 116: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 78: return -1
  case 117: return -1
  case 99: return -1
  case 111: return -1
  case 85: return -1
  case 110: return -1
  case 67: return -1
  case 79: return 3
  case 116: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 99: return 1
  case 111: return -1
  case 85: return -1
  case 110: return -1
  case 67: return 2
  case 79: return -1
  case 116: return -1
  case 84: return -1
  case 78: return -1
  case 117: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 78: return -1
  case 117: return -1
  case 99: return -1
  case 111: return -1
  case 85: return -1
  case 110: return -1
  case 67: return -1
  case 79: return -1
  case 116: return 10
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 85: return -1
  case 110: return -1
  case 67: return -1
  case 79: return -1
  case 116: return -1
  case 84: return -1
  case 78: return 5
  case 117: return -1
  case 99: return -1
  case 111: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 78: return -1
  case 117: return 8
  case 99: return -1
  case 111: return -1
  case 85: return -1
  case 110: return -1
  case 67: return -1
  case 79: return -1
  case 116: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 79: return -1
  case 116: return -1
  case 84: return 6
  case 78: return -1
  case 117: return -1
  case 99: return -1
  case 111: return -1
  case 85: return -1
  case 110: return -1
  case 67: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 67: return -1
  case 79: return -1
  case 116: return -1
  case 84: return -1
  case 78: return -1
  case 117: return -1
  case 99: return -1
  case 111: return -1
  case 85: return 4
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[28].acc = acc[:]
a0[28].f = fun[:]
a0[28].id = 28
}
{
var acc [7]bool
var fun [7]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 117: return -1
  case 83: return 1
  case 77: return -1
  case 115: return 2
  case 109: return -1
  case 85: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 109: return -1
  case 85: return -1
  case 117: return -1
  case 83: return -1
  case 77: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 109: return -1
  case 85: return -1
  case 117: return -1
  case 83: return -1
  case 77: return -1
  case 115: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 109: return -1
  case 85: return -1
  case 117: return 3
  case 83: return -1
  case 77: return -1
  case 115: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 77: return -1
  case 115: return -1
  case 109: return 4
  case 85: return -1
  case 117: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 85: return -1
  case 117: return -1
  case 83: return -1
  case 77: return 6
  case 115: return -1
  case 109: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 109: return -1
  case 85: return 5
  case 117: return -1
  case 83: return -1
  case 77: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[29].acc = acc[:]
a0[29].f = fun[:]
a0[29].id = 29
}
{
var acc [13]bool
var fun [13]func(rune) int
fun[8] = func(r rune) int {
  switch(r) {
  case 112: return -1
  case 65: return 9
  case 77: return -1
  case 80: return -1
  case 114: return -1
  case 97: return -1
  case 109: return -1
  case 82: return -1
  case 71: return -1
  case 103: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 112: return -1
  case 65: return -1
  case 77: return -1
  case 80: return -1
  case 114: return -1
  case 97: return -1
  case 109: return 6
  case 82: return -1
  case 71: return -1
  case 103: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[7] = true
fun[7] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 77: return -1
  case 80: return -1
  case 114: return -1
  case 97: return -1
  case 109: return -1
  case 82: return -1
  case 71: return -1
  case 103: return -1
  case 112: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 103: return -1
  case 112: return -1
  case 65: return -1
  case 77: return -1
  case 80: return -1
  case 114: return -1
  case 97: return -1
  case 109: return -1
  case 82: return -1
  case 71: return 10
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 77: return -1
  case 80: return 1
  case 114: return -1
  case 97: return -1
  case 109: return -1
  case 82: return -1
  case 71: return -1
  case 103: return -1
  case 112: return 2
  case 65: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 103: return -1
  case 112: return -1
  case 65: return -1
  case 77: return -1
  case 80: return -1
  case 114: return 3
  case 97: return -1
  case 109: return -1
  case 82: return -1
  case 71: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[12] = true
fun[12] = func(r rune) int {
  switch(r) {
  case 82: return -1
  case 71: return -1
  case 103: return -1
  case 112: return -1
  case 65: return -1
  case 77: return -1
  case 80: return -1
  case 114: return -1
  case 97: return -1
  case 109: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 97: return -1
  case 109: return -1
  case 82: return -1
  case 71: return -1
  case 103: return 5
  case 112: return -1
  case 65: return -1
  case 77: return -1
  case 80: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 112: return -1
  case 65: return -1
  case 77: return -1
  case 80: return -1
  case 114: return -1
  case 97: return 7
  case 109: return -1
  case 82: return -1
  case 71: return -1
  case 103: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[10] = func(r rune) int {
  switch(r) {
  case 97: return -1
  case 109: return -1
  case 82: return -1
  case 71: return -1
  case 103: return -1
  case 112: return -1
  case 65: return -1
  case 77: return 11
  case 80: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 97: return -1
  case 109: return -1
  case 82: return 8
  case 71: return -1
  case 103: return -1
  case 112: return -1
  case 65: return -1
  case 77: return -1
  case 80: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[11] = func(r rune) int {
  switch(r) {
  case 65: return 12
  case 77: return -1
  case 80: return -1
  case 114: return -1
  case 97: return -1
  case 109: return -1
  case 82: return -1
  case 71: return -1
  case 103: return -1
  case 112: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 103: return -1
  case 112: return -1
  case 65: return -1
  case 77: return -1
  case 80: return -1
  case 114: return -1
  case 97: return 4
  case 109: return -1
  case 82: return -1
  case 71: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[30].acc = acc[:]
a0[30].f = fun[:]
a0[30].id = 30
}
{
var acc [15]bool
var fun [15]func(rune) int
acc[14] = true
fun[14] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 76: return -1
  case 101: return -1
  case 69: return -1
  case 108: return -1
  case 80: return -1
  case 110: return -1
  case 112: return -1
  case 88: return -1
  case 120: return -1
  case 97: return -1
  case 78: return -1
  case 65: return -1
  case 105: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[8] = true
fun[8] = func(r rune) int {
  switch(r) {
  case 108: return -1
  case 80: return -1
  case 110: return -1
  case 112: return -1
  case 88: return -1
  case 120: return -1
  case 97: return -1
  case 78: return -1
  case 65: return -1
  case 105: return -1
  case 73: return -1
  case 76: return -1
  case 101: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 69: return -1
  case 108: return -1
  case 80: return -1
  case 110: return -1
  case 112: return -1
  case 88: return 3
  case 120: return -1
  case 97: return -1
  case 78: return -1
  case 65: return -1
  case 105: return -1
  case 73: return -1
  case 76: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[11] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 69: return -1
  case 108: return -1
  case 80: return -1
  case 110: return -1
  case 112: return -1
  case 88: return -1
  case 120: return -1
  case 97: return 12
  case 78: return -1
  case 65: return -1
  case 105: return -1
  case 73: return -1
  case 76: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 69: return -1
  case 108: return -1
  case 80: return -1
  case 110: return -1
  case 112: return -1
  case 88: return -1
  case 120: return -1
  case 97: return -1
  case 78: return -1
  case 65: return -1
  case 105: return -1
  case 73: return 7
  case 76: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[13] = func(r rune) int {
  switch(r) {
  case 105: return -1
  case 73: return -1
  case 76: return -1
  case 101: return -1
  case 69: return -1
  case 108: return -1
  case 80: return -1
  case 110: return 14
  case 112: return -1
  case 88: return -1
  case 120: return -1
  case 97: return -1
  case 78: return -1
  case 65: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 80: return -1
  case 110: return -1
  case 112: return -1
  case 88: return -1
  case 120: return -1
  case 97: return -1
  case 78: return -1
  case 65: return -1
  case 105: return -1
  case 73: return -1
  case 76: return -1
  case 101: return 1
  case 69: return 2
  case 108: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 105: return -1
  case 73: return -1
  case 76: return -1
  case 101: return -1
  case 69: return -1
  case 108: return -1
  case 80: return 4
  case 110: return -1
  case 112: return -1
  case 88: return -1
  case 120: return -1
  case 97: return -1
  case 78: return -1
  case 65: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 76: return 5
  case 101: return -1
  case 69: return -1
  case 108: return -1
  case 80: return -1
  case 110: return -1
  case 112: return -1
  case 88: return -1
  case 120: return -1
  case 97: return -1
  case 78: return -1
  case 65: return -1
  case 105: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 80: return -1
  case 110: return -1
  case 112: return -1
  case 88: return -1
  case 120: return -1
  case 97: return -1
  case 78: return 8
  case 65: return -1
  case 105: return -1
  case 73: return -1
  case 76: return -1
  case 101: return -1
  case 69: return -1
  case 108: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 69: return -1
  case 108: return -1
  case 80: return -1
  case 110: return -1
  case 112: return -1
  case 88: return -1
  case 120: return -1
  case 97: return -1
  case 78: return -1
  case 65: return 6
  case 105: return -1
  case 73: return -1
  case 76: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 76: return -1
  case 101: return -1
  case 69: return -1
  case 108: return -1
  case 80: return -1
  case 110: return -1
  case 112: return -1
  case 88: return -1
  case 120: return 9
  case 97: return -1
  case 78: return -1
  case 65: return -1
  case 105: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 76: return -1
  case 101: return -1
  case 69: return -1
  case 108: return -1
  case 80: return -1
  case 110: return -1
  case 112: return 10
  case 88: return -1
  case 120: return -1
  case 97: return -1
  case 78: return -1
  case 65: return -1
  case 105: return -1
  case 73: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[12] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 69: return -1
  case 108: return -1
  case 80: return -1
  case 110: return -1
  case 112: return -1
  case 88: return -1
  case 120: return -1
  case 97: return -1
  case 78: return -1
  case 65: return -1
  case 105: return 13
  case 73: return -1
  case 76: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[10] = func(r rune) int {
  switch(r) {
  case 88: return -1
  case 120: return -1
  case 97: return -1
  case 78: return -1
  case 65: return -1
  case 105: return -1
  case 73: return -1
  case 76: return -1
  case 101: return -1
  case 69: return -1
  case 108: return 11
  case 80: return -1
  case 110: return -1
  case 112: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[31].acc = acc[:]
a0[31].f = fun[:]
a0[31].id = 31
}
{
var acc [5]bool
var fun [5]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 108: return -1
  case 117: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 108: return 3
  case 117: return -1
  case 110: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 117: return -1
  case 110: return -1
  case 108: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 108: return 4
  case 117: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 110: return 1
  case 108: return -1
  case 117: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[32].acc = acc[:]
a0[32].f = fun[:]
a0[32].id = 32
}
{
var acc [3]bool
var fun [3]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 95: return 1
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 57: return -1
    case 65 <= r && r <= 90: return 1
    case 97 <= r && r <= 122: return 1
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 45: return 2
  case 95: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 90: return 2
    case 97 <= r && r <= 122: return 2
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 45: return 2
  case 95: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 90: return 2
    case 97 <= r && r <= 122: return 2
    default: return -1
    }
  }
  panic("unreachable")
}
a0[33].acc = acc[:]
a0[33].f = fun[:]
a0[33].id = 33
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 61: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 61: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[34].acc = acc[:]
a0[34].f = fun[:]
a0[34].id = 34
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 43: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 43: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[35].acc = acc[:]
a0[35].f = fun[:]
a0[35].id = 35
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 45: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 45: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[36].acc = acc[:]
a0[36].f = fun[:]
a0[36].id = 36
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 42: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 42: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[37].acc = acc[:]
a0[37].f = fun[:]
a0[37].id = 37
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 47: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 47: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[38].acc = acc[:]
a0[38].f = fun[:]
a0[38].id = 38
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 37: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 37: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[39].acc = acc[:]
a0[39].f = fun[:]
a0[39].id = 39
}
{
var acc [3]bool
var fun [3]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 38: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 38: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 38: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[40].acc = acc[:]
a0[40].f = fun[:]
a0[40].id = 40
}
{
var acc [3]bool
var fun [3]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 124: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 124: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 124: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[41].acc = acc[:]
a0[41].f = fun[:]
a0[41].id = 41
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 33: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 33: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[42].acc = acc[:]
a0[42].f = fun[:]
a0[42].id = 42
}
{
var acc [3]bool
var fun [3]func(rune) int
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 61: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 61: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 61: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[43].acc = acc[:]
a0[43].f = fun[:]
a0[43].id = 43
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 60: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 60: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[44].acc = acc[:]
a0[44].f = fun[:]
a0[44].id = 44
}
{
var acc [3]bool
var fun [3]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 61: return 2
  case 60: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 60: return -1
  case 61: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 60: return 1
  case 61: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[45].acc = acc[:]
a0[45].f = fun[:]
a0[45].id = 45
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 62: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 62: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[46].acc = acc[:]
a0[46].f = fun[:]
a0[46].id = 46
}
{
var acc [3]bool
var fun [3]func(rune) int
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 61: return -1
  case 62: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 61: return 2
  case 62: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 62: return 1
  case 61: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[47].acc = acc[:]
a0[47].f = fun[:]
a0[47].id = 47
}
{
var acc [3]bool
var fun [3]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 61: return 2
  case 33: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 61: return -1
  case 33: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 33: return -1
  case 61: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[48].acc = acc[:]
a0[48].f = fun[:]
a0[48].id = 48
}
{
var acc [3]bool
var fun [3]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 62: return 2
  case 60: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 60: return -1
  case 62: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 60: return 1
  case 62: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[49].acc = acc[:]
a0[49].f = fun[:]
a0[49].id = 49
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 46: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 46: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[50].acc = acc[:]
a0[50].f = fun[:]
a0[50].id = 50
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 40: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 40: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[51].acc = acc[:]
a0[51].f = fun[:]
a0[51].id = 51
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 41: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 41: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[52].acc = acc[:]
a0[52].f = fun[:]
a0[52].id = 52
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 44: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 44: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[53].acc = acc[:]
a0[53].f = fun[:]
a0[53].id = 53
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 123: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 123: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[54].acc = acc[:]
a0[54].f = fun[:]
a0[54].id = 54
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 125: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 125: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[55].acc = acc[:]
a0[55].f = fun[:]
a0[55].id = 55
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 91: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 91: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[56].acc = acc[:]
a0[56].f = fun[:]
a0[56].id = 56
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 93: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 93: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[57].acc = acc[:]
a0[57].f = fun[:]
a0[57].id = 57
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 63: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 63: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[58].acc = acc[:]
a0[58].f = fun[:]
a0[58].id = 58
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 58: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 58: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[59].acc = acc[:]
a0[59].f = fun[:]
a0[59].id = 59
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 10: return 1
  case 9: return 1
  case 32: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 32: return 1
  case 10: return 1
  case 9: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[60].acc = acc[:]
a0[60].f = fun[:]
a0[60].id = 60
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  default:
    switch {
    default: return 1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[61].acc = acc[:]
a0[61].f = fun[:]
a0[61].id = 61
}
a[0].endcase = 62
a[0].a = a0[:]
}
func getAction(c *frame) int {
  if -1 == c.match { return -1 }
  c.action = c.fam.a[c.match].id
  c.match = -1
  return c.action
}
type frame struct {
  atEOF bool
  action, match, matchn, n int
  buf []rune
  text string
  in *bufio.Reader
  state []int
  fam family
}
func newFrame(in *bufio.Reader, index int) *frame {
  f := new(frame)
  f.buf = make([]rune, 0, 128)
  f.in = in
  f.match = -1
  f.fam = a[index]
  f.state = make([]int, len(f.fam.a))
  return f
}
type Lexer []*frame
func NewLexer(in io.Reader) Lexer {
  stack := make([]*frame, 0, 4)
  stack = append(stack, newFrame(bufio.NewReader(in), 0))
  return stack
}
func (stack Lexer) isDone() bool {
  return 1 == len(stack) && stack[0].atEOF
}
func (stack Lexer) nextAction() int {
  c := stack[len(stack) - 1]
  for {
    if c.atEOF { return c.fam.endcase }
    if c.n == len(c.buf) {
      r,_,er := c.in.ReadRune()
      switch er {
      case nil: c.buf = append(c.buf, r)
      case io.EOF:
	c.atEOF = true
	if c.n > 0 {
	  c.text = string(c.buf)
	  return getAction(c)
	}
	return c.fam.endcase
      default: panic(er.Error())
      }
    }
    jammed := true
    r := c.buf[c.n]
    for i, x := range c.fam.a {
      if -1 == c.state[i] { continue }
      c.state[i] = x.f[c.state[i]](r)
      if -1 == c.state[i] { continue }
      jammed = false
      if x.acc[c.state[i]] {
	if -1 == c.match || c.matchn < c.n+1 || c.match > i {
	  c.match = i
	  c.matchn = c.n+1
	}
      }
    }
    if jammed {
      a := getAction(c)
      if -1 == a { c.matchn = c.n + 1 }
      c.n = 0
      for i, _ := range c.state { c.state[i] = 0 }
      c.text = string(c.buf[:c.matchn])
      copy(c.buf, c.buf[c.matchn:])
      c.buf = c.buf[:len(c.buf) - c.matchn]
      return a
    }
    c.n++
  }
  panic("unreachable")
}
func (stack Lexer) push(index int) Lexer {
  c := stack[len(stack) - 1]
  return append(stack,
      newFrame(bufio.NewReader(strings.NewReader(c.text)), index))
}
func (stack Lexer) pop() Lexer {
  return stack[:len(stack) - 1]
}
func (stack Lexer) Text() string {
  c := stack[len(stack) - 1]
  return c.text
}
func (yylex Lexer) Error(e string) {
  panic(e)
}
func (yylex Lexer) Lex(lval *yySymType) int {
  for !yylex.isDone() {
    switch yylex.nextAction() {
    case -1:
    case 0:  //(-)?[0-9]+/
{                      
                    lval.n,_ = strconv.Atoi(yylex.Text());
                    logDebugTokens("INT: %d", lval.n); 
                    return INT 
                  }
    case 1:  //(-)?[0-9]+\.[0-9]*/
{ 
                    lval.f,_ = strconv.ParseFloat(yylex.Text(), 64);
                    logDebugTokens("REAL: %f", lval.f);  
                    return REAL 
                  }
    case 2:  //"[^"\\]*(\\.[^"\\]*)*"/
{ 
                    lval.s = yylex.Text()[1:len(yylex.Text())-1];
                    logDebugTokens("STRING: %s", lval.s);
                    return STRING }
    case 3:  //'[^'\\]*(\\.[^'\\]*)*'/
{ 
                    lval.s = yylex.Text()[1:len(yylex.Text())-1];
                    logDebugTokens("STRING: %s", lval.s);
                    return STRING }
    case 4:  //SELECT|select/
{ logDebugTokens("SELECT"); return SELECT }
    case 5:  //DISTINCT|distinct/
{ logDebugTokens("DISTINCT"); return DISTINCT }
    case 6:  //ALL|all/
{ logDebugTokens("ALL"); return ALL }
    case 7:  //AS|as/
{ logDebugTokens("AS"); return AS }
    case 8:  //FROM|from/
{ logDebugTokens("FROM"); return FROM }
    case 9:  //WHERE|where/
{ logDebugTokens("WHERE"); return WHERE }
    case 10:  //GROUP|group/
{ logDebugTokens("GROUP"); return GROUP }
    case 11:  //BY|by/
{ logDebugTokens("BY"); return BY }
    case 12:  //HAVING|having/
{ logDebugTokens("HAVING"); return HAVING }
    case 13:  //FLATTEN|flatten/
{ logDebugTokens("FLATTEN"); return FLATTEN }
    case 14:  //EACH|each/
{ logDebugTokens("EACH"); return EACH }
    case 15:  //UNION|union/
{ logDebugTokens("UNION"); return UNION }
    case 16:  //INTERSECT|intersect/
{ logDebugTokens("INTERSECT"); return INTERSECT }
    case 17:  //EXCEPT|except/
{ logDebugTokens("EXCEPT"); return EXCEPT }
    case 18:  //ORDER|order/
{ logDebugTokens("ORDER"); return ORDER }
    case 19:  //LIMIT|limit/
{ logDebugTokens("LIMIT"); return LIMIT }
    case 20:  //OFFSET|offset/
{ logDebugTokens("OFFSET"); return OFFSET }
    case 21:  //ASC|asc/
{ logDebugTokens("ASC"); return ASC }
    case 22:  //DESC|desc/
{ logDebugTokens("DESC"); return DESC }
    case 23:  //TRUE|true/
{ logDebugTokens("TRUE"); return TRUE }
    case 24:  //FALSE|false/
{ logDebugTokens("FALSE"); return FALSE }
    case 25:  //MAX|max/
{ logDebugTokens("MAX"); return MAX }
    case 26:  //MIN|min/
{ logDebugTokens("MIN"); return MIN }
    case 27:  //AVG|avg/
{ logDebugTokens("AVG"); return AVG }
    case 28:  //COUNT|count/
{ logDebugTokens("COUNT"); return COUNT }
    case 29:  //SUM|sum/
{ logDebugTokens("SUM"); return SUM }
    case 30:  //PRAGMA|pragma/
{ logDebugTokens("PRAGMA"); return PRAGMA }
    case 31:  //EXPLAIN|explain/
{ logDebugTokens("EXPLAIN"); return EXPLAIN }
    case 32:  //null/
{ logDebugTokens("NULL"); return NULL }
    case 33:  //[a-zA-Z_][a-zA-Z0-9\-_]*/
{ 
                        lval.s = yylex.Text();
                        logDebugTokens("IDENTIFIER: %s", lval.s);
                        return IDENTIFIER 
                    }
    case 34:  //\=/
{ logDebugTokens("ASSIGN"); return ASSIGN }
    case 35:  //\+/
{ logDebugTokens("PLUS"); return PLUS }
    case 36:  //-/
{ logDebugTokens("MINUS"); return MINUS }
    case 37:  //\*/
{ logDebugTokens("MULT"); return MULT }
    case 38:  //\//
{ logDebugTokens("DIV"); return DIV }
    case 39:  //\%/
{ logDebugTokens("MOD"); return MOD }
    case 40:  //\&\&/
{ logDebugTokens("AND"); return AND }
    case 41:  //\|\|/
{ logDebugTokens("OR"); return OR }
    case 42:  //\!/
{ logDebugTokens("NOT"); return NOT }
    case 43:  //\=\=/
{ logDebugTokens("EQ"); return EQ }
    case 44:  //\</
{ logDebugTokens("LT"); return LT }
    case 45:  //\<\=/
{ logDebugTokens("LTE"); return LTE }
    case 46:  //\>/
{ logDebugTokens("GT"); return GT }
    case 47:  //\>\=/
{ logDebugTokens("GTE"); return GTE }
    case 48:  //\!\=/
{ logDebugTokens("NE"); return NE }
    case 49:  //\<\>/
{ logDebugTokens("NE"); return NE }
    case 50:  //\./
{ logDebugTokens("DOT"); return DOT }
    case 51:  //\(/
{ logDebugTokens("LPAREN"); return LPAREN }
    case 52:  //\)/
{ logDebugTokens("RPAREN"); return RPAREN }
    case 53:  //\,/
{ logDebugTokens("COMMA"); return COMMA }
    case 54:  //\{/
{ logDebugTokens("LBRACE"); return LBRACE }
    case 55:  //\}/
{ logDebugTokens("RBRACE"); return RBRACE }
    case 56:  //\[/
{ logDebugTokens("LBRACKET"); return LBRACKET }
    case 57:  //\]/
{ logDebugTokens("RBRACKET"); return RBRACKET }
    case 58:  //\?/
{ logDebugTokens("QUESTION"); return QUESTION }
    case 59:  //\:/
{ logDebugTokens("COLON"); return COLON }
    case 60:  //[ \t\n]+/
{ logDebugTokens("WHITESPACE (count=%d)", len(yylex.Text())) /* eat up whitespace */ }
    case 61:  //./
{ log.Printf("see problem: %v", yylex.Text()); return int(yylex.Text()[0]) }
    case 62:  ///
// [END]
    }
  }
  return 0
}
func logDebugTokens(format string, v ...interface{}) {
    if debugTokens && len(v) > 0 {
        log.Printf("DEBUG TOKEN " + format, v)
    } else if debugTokens {
        log.Printf("DEBUG TOKEN " + format)
    }
}

func required_to_make_nex_happy_with_main_somewhere_else() {
}