package main

func combine(n int, k int) [][]int {

	trace := make([]int, 0)
	result := make([][]int, 0)
	var backtrace func(int)
	backtrace = func(start int) {
		if len(trace) == k {
			result = append(result, append([]int(nil), trace...))
			return
		}
		for i := start; i < n; i++ {
			trace = append(trace, i+1)
			backtrace(i + 1)
			trace = trace[:len(trace)-1]
		}
	}
	backtrace(0)
	return result
}
