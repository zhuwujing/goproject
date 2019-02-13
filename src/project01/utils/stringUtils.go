package utils
import (
   "fmt"
   "strconv"
)
/*
1)统计字符串的长度，按字节 len(str)
2)字符串遍历，同时处理有中文的问题 r := []rune(str)
3)字符串转整数:	n, err := strconv.Atoi("12")
4)整数转字符串	str = strconv.Itoa(12345)
5)字符串	转	[]byte: var bytes = []byte("hello go")
6)[]byte 转 字符串: str = string([]byte{97, 98, 99})
7) 10 进 制转	2, 8, 16 进制: str = strconv.FormatInt(123, 2) // 2-> 8 , 16
8)查找子串是否在指定的字符串中: strings.Contains("seafood", "foo") //true
9)统计一个字符串有几个指定的子串 ： strings.Count("ceheese", "e") //4
10)不区分大小写的字符串比较(== 是区分字母大小写的): fmt.Println(strings.EqualFold("abc", "Abc")) // true
11)返回子串在字符串第一次出现的 index 值，如果没有返回-1 : strings.Index("NLT_abc", "abc") // 4
12)返回子串在字符串最后一次出现的 index，如没有返回-1 : strings.LastIndex("go golang", "go")
13)将指定的子串替换成 另外一个子串: strings.Replace("go go hello", "go", "go 语言", n) n 可以指定你希望替换几个，如果 n=-1 表示全部替换
14)按 照 指 定 的 某 个 字 符 ， 为 分 割 标 识 ， 将 一 个 字 符 串 拆 分 成 字 符 串 数 组 ：
strings.Split("hello,wrold,ok", ",")
15)将字符串的字母进行大小写的转换: strings.ToLower("Go") // go strings.ToUpper("Go") // GO
16)将字符串左右两边的空格去掉：   strings.TrimSpace(" tn a lone gopher ntrn	")
17)将字符串左右两边指定的字符去掉	：	strings.Trim("! hello! ", " !") // ["hello"] //将左右两边 !
和 " "去掉
18)将字符串左边指定的字符去掉	：	strings.TrimLeft("! hello! ", " !") // ["hello"] // 将 左 边 ! 和 "
"去掉
19)将字符串右边指定的字符去掉	： strings.TrimRight("! hello! ", " !") // ["hello"] //将右边 ! 和 "
"去掉
20)判断字符串是否以指定的字符串开头: strings.HasPrefix("ftp://192.168.10.1", "ftp") // true
21)判断字符串是否以指定的字符串结束: strings.HasSuffix("NLT_abc.jpg", "abc") //false
*/

func Len(str string) int {
   return len(str)
}

/*
  string 的底层是byte数组，for循环是按字节遍历  []rune  是按字符遍历
*/
func traverse(str string) {
   for i := 0; i < len(str); i++ {
      fmt.Printf("utils str[%d] = %c \n",i ,str[i])
   }
  
}

func TraverseByRune(str string) {
   // rune := []rune(str)
   // var rune []rune = []rune(str)
   var rune = []rune(str)
   for i := 0; i < len(rune); i++ {
      fmt.Printf("utils str[%d] = %c \n",i ,rune[i])
   }
  
}

func Atoi(str string)(int ,error){
 return strconv.Atoi(str)
}

func Itoa(a int) (string){
   return strconv.Itoa(a)
}

func Byte2Str(bytes []byte)string{
   s := string(bytes)
   return s
}

func Str2Byte(str string) []byte{
   return []byte(str)
}

func FormatInt(a int64,base int)string{
   return strconv.FormatInt(a,base)
}