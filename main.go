package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
          fmt.Println("CalculatorGo v.0.1")
       intType, first, second, sign, err := readLine()
       if err != nil {
       panic(fmt.Sprintf("Ошибка при вводе данных:\n%s", err))
 }
        if intType == "arab" {
       firstNum, err1 := strconv.Atoi(first)
       if err1 != nil {
       panic(fmt.Sprintf("Ошибка при переводе строки в число:\n%s", err1))
  }
       secondNum, err2 := strconv.Atoi(second)
       if err2 != nil {
       panic(fmt.Sprintf("Ошибка при переводе строки в число:\n%s", err2))
  }
       if firstNum > 10 || secondNum > 10 {
       panic("Число должно быть не больше 10")
  }
       res, err3 := calculator(firstNum, secondNum, sign)
       if err3 != nil {
        panic(fmt.Sprintf("Ошибка при работе калькулятора:\n%s", err3))
  } else {
        fmt.Println("Ответ: ", res)
  }
 } else {
      firstNum := fromRomanToInt(first)
      secondNum := fromRomanToInt(second)
      if firstNum > 10 || secondNum > 10 {
       panic("Число должно быть не больше 10")
      }
      res, err1 := calculator(firstNum, secondNum, sign)
      if err1 != nil {
       panic(fmt.Sprintf("Ошибка при работе калькулятора:\n%s", err1))
  } else {
        final, err2 := fromIntToRoman(res)
       if err2 != nil {
       panic(fmt.Sprintf("Ошибка при работе калькулятора:\n%s", err2))
   }
        fmt.Println("Ответ: ", final)
  }
 }
}

func calculator(first int, second int, sign string) (int, error) {
       switch {
       case sign == "+":
        return first + second, nil
       case sign == "-":
        return first - second, nil
       case sign == "*":
         return first * second, nil
       case sign == "/" && second != 0:
       return first / second, nil
        case sign == "/" && second == 0:
        panic("Деление на ноль")
       default:
       panic("Неверный оператор")
 }
}

func readLine() (string, string, string, string, error) {
       stdin := bufio.NewReader(os.Stdin)
       usInput, _ := stdin.ReadString('\n')
       usInput = strings.TrimSpace(usInput)
       intType, first, second, sign, err := checkInput(usInput)
 if err != nil {
   return "", "", "", "", err
 }
   return intType, first, second, sign, err
}

func checkInput(input string) (string, string, string, string, error) {
       r := regexp.MustCompile("\\s+")
       replace := r.ReplaceAllString(input, "")
       arr := strings.Split(replace, "")
 var intType, first, second, sign string
 for index, value := range arr {
        isN := isNumber(value)
       isS := isSign(value)
       isR := isRomanNumber(value)
  if !isN && !isS && !isR {
   panic("Недопустимые символы в выражении")
  }
  if isS {
   if sign != "" {
    panic("Нет допустимых операторов")
   } else {
    sign = arr[index]
   }
  }
  if (isN && intType != "roman") || (isR && intType != "arab") {
   if intType == "" {
    if isN {
     intType = "arab"
    } else {
     intType = "roman"
    }
   }
   if first == "" && !(index+1 == len(arr)) && isSign(arr[index+1]) {
    slice := arr[0 : index+1]
    first = strings.Join(slice, "")
   } else if index+1 == len(arr) && first != "" {
    slice := arr[(len(first) + 1):]
    second = strings.Join(slice, "")
   }
  } else if (intType == "arab" && isR) || (intType == "roman" && isN) {
   panic("Неверный тип числа (арабское и римское)")
  }
 }
 if second == "" || first == "" || sign == "" {
  panic("Недостаточно данных в выражении")
 }
 return intType, first, second, sign, nil
}

func isNumber(c string) bool {
 if c >= "0" && c <= "9" {
  return true
 }
 return false
}

func isSign(c string) bool {
 if c == "+" || c == "-" || c == "/" || c == "*" {
  return true
 }
 return false
}

func isRomanNumber(c string) bool {
 _, ok := dict[c]
 if ok {
       return true
 }
       return false
}

func errorHandler(code int) error {
	return errors.New(errorDict[code])
}

var errorDict = map[int]string{
	1: "Нераспознанные символы. Пожалуйста, используйте только арабские/римские цифры и математические операторы '+', '-', '/', '*' ",
	2: "Некорректный ввод. Пожалуйста, используйте только арабские или только римские цифры ",
	3: "Неверное количество аргументов. Для работы калькулятора необходимо два числа и математический оператор",
	4: "Не умею делить на ноль, но когда-нибудь обязательно научусь!",
	5: "Что-то пошло не так при вычислениях, нужно время чтобы разобраться",
	6: "Я пока умею выполнять только по одной операции за раз. Пожалуйста, введите только два числа и один математический оператор",
	7: "Не могу отобразить ответ, так как в римской системе нет отрицательных чисел",
	8: "Пожалуйста, введите числа от 0 до 10 включительно",
}

var dict = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func fromRomanToInt(roman string) int {
	var res int
	arr := strings.Split(roman, "")
	for index, value := range arr {
		if index+1 != len(arr) && dict[value] < dict[arr[index+1]] {
			res -= dict[value]
		} else {
			res += dict[value]
		}
	}
	return res
}

func fromIntToRoman(number int) (string, error) {
	if number <= 0 {
		return "", errorHandler(7)
	}
	arr1 := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	arr2 := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var str string
	for number > 0 {
		for i := 0; i < 13; i++ {
			if arr1[i] <= number {
				str += arr2[i]
				number -= arr1[i]
				break
			}
		}
	}
	return str, nil
}
