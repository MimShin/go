package roman

func romanToArabic(roman []rune) int {
	v0 := DigitValue(roman[0])
	if len(roman) < 2 {
		return v0
	}

	v1 := DigitValue(roman[1])
	value := romanToArabic(roman[1:])

	if v0 < v1 {
		return value - v0
	}
	return value + v0
}

func RomanToArabic(roman string) int {
	r := []rune(roman)

	value := 0
	oldDigitValue := 0

	for i:=len(r)-1; i>=0; i-- {
		v := DigitValue(r[i])
		if v < oldDigitValue {
			value -= v
		} else {
			value += v
		}
		oldDigitValue = v
	}

	return value
}

func DigitValue(digit rune) int {
	switch digit {
	case 'I': return 1
	case 'V': return 5
	case 'X': return 10
	case 'L': return 50
	case 'C': return 100
	case 'D': return 500
	case 'M': return 1000
	default: return 0
	}
}

func ArabicToRoman(num int) string {
	roman := []rune{}

	romanDigits := []rune{'M', 'D', 'C', 'L', 'X', 'V', 'I'}
	i := 0
	l := len(romanDigits)

	for ; i<l; i++ {
		digit := romanDigits[i]
		vi := DigitValue(digit)
		//fmt.Printf("A2R %d/%d %c %d\n", i, l, digit, vi)
		for num >= vi {
			num -= vi
			roman = append(roman, digit)
		}

		// Support for IV, IX, XL, XC, CD, CM
		if i < l-1 {
			subDigit := romanDigits[i/2*2+2]
			delNumVal := DigitValue(subDigit)
			if num+delNumVal >= vi {
				roman = append(roman, subDigit)
				roman = append(roman, digit)
				num -= (vi - delNumVal)
			}
		}
	}
	return string(roman)
}
