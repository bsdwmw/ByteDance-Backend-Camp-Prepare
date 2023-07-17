//36进制由0-9，a-z，共36个字符表示。

//要求按照加法规则计算出任意两个36进制正整数的和，如1b + 2x = 48 （解释：1b = 1*36 + 11*0, 2x = 2*36 + 33, 1b + 2x = 47 + 105 = 152）

//要求：不允许使用先将36进制数字整体转为10进制，相加后再转回为36进制的做法

package main

import "fmt"

func add(a, b string) string {
  refList := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
  i := len(a)-1
  j := len(b)-1
  carry := 0
  ans := ""
  for :i>=0||j>=0||carry>0;i,j = i-1, j-1 {
    valI := GetInt(a[i])
    valJ := GetInt(b[j])

    sum := valI + valJ + carry
    carry = sum/36
    ans = refList[sum%36]+ans
  }
  return ans
}

func ConvertInt(c uint8) int {
  if c == "0" {
    return 0
  } else if c > '0' && c < '9' {
    return int(c-'0')
  } else {
    return int(c-'a') + 10
  }
}

func main() {
  s := add("1a", "2b")
  fmt.Println(s)
}
  
