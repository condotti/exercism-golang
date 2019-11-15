package pascal

// binomial coefficient version
// 
// func Triangle(n int)[][]int {
// 	bc := func(n, k int)int {
// 		nk , fk := 1, 1
// 		for i := n; i > n - k; i-- { nk *= i }
// 		for i := k; i > 0; i-- { fk *= i }
// 		return nk / fk
// 	}
// 	row := []int{1}
// 	triangle := [][]int{row}
// 	for i := 1; i < n; i++ {
// 		row = []int{1}
// 		for j := 1; j <= i; j++ {
// 			row = append(row, bc(i, j))
// 		}
// 		triangle = append(triangle, row)
// 	}
// 	return triangle
// }
// BenchmarkPascalsTriangleFixed-4        	  131092	      9015 ns/op	    6144 B/op	      95 allocs/op
// BenchmarkPascalsTriangleIncreasing-4   	   16458	     68414 ns/op	   45512 B/op	     886 allocs/op

// recurrence version

func Triangle(n int)[][]int {
	row := []int{1}
	triangle := [][]int{row}
	for i := 1; i < n; i++ {
		seed := append(append([]int{0}, row...), 0)
		nextRow := []int{}
		for j := 0; j <= i; j++ {
			nextRow = append(nextRow, seed[j]+seed[j+1])
		}
		triangle = append(triangle, nextRow)
		row = nextRow
	}
	return triangle
}
// BenchmarkPascalsTriangleFixed-4        	  158313	      6804 ns/op	    9648 B/op	     124 allocs/op
// BenchmarkPascalsTriangleIncreasing-4   	   20497	     64133 ns/op	   70712 B/op	    1176 allocs/op
