package fizzbuzz

// var ops = []func(int) (string, bool){
// 	func(num int) (string, bool) {
// 		if num%3 == 0 {
// 			return "fizz", true
// 		}
// 		return "", false
// 	},
// 	func(num int) (string, bool) {
// 		if num%5 == 0 {
// 			return "fizz", true
// 		}
// 		return "", false
// 	},
// }

// Fizzbuzz takes a number, and returns a string and boolean true, or an empty string and boolean false
// Callers are supposed to print original number if false is returned
func Fizzbuzz(num int) (string, bool) {
	if num%15 == 0 {
		return "fizzbuzz", true
	}
	if num%5 == 0 {
		return "buzz", true
	}
	if num%3 == 0 {
		return "fizz", true
	}
	return "", false
}
