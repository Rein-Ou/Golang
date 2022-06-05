package fibo

/**
func Fibonacci(n int)(res int){
	if n<=1{
		res=1
	}else{
		res=Fibonacci(n-1)+Fibonacci(n-2)
	}
	return
*/

func Fibonacci(op string, num int) (res int) {
	if num <= 1 {
		switch op {
		case "+":
			res = 1
		case "*":
			res = 2
		default:
			res = 0
		}
	} else {
		switch op {
		case "+":
			res = Fibonacci(op, num-1) + Fibonacci(op, num-2)
		case "*":
			res = Fibonacci(op, num-1) * Fibonacci(op, num-2)
		default:
			res = 0
		}
	}
	return
}
