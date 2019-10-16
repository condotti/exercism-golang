package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

func (xs IntList) Foldl(f binFunc, initial int) int {
	for _, x := range xs {
		initial = f(initial, x)
	}
	return initial
}

func (xs IntList) Foldr(f binFunc, initial int) int {
	for i, _ := range xs {
		initial = f(xs[xs.Length()-i-1], initial)
	}
	return initial
}

func (xs IntList) Filter(f predFunc) IntList {
	filtered := IntList{}
	for _, x := range xs {
		if f(x) {
			filtered = filtered.Append(IntList{x})
		}
	}
	return filtered
}

func (xs IntList) Length() int {
	// i'm not sure if it is the intention
	length := 0
	for range xs { length++ }
	return length
}

func (xs IntList) Map(f unaryFunc) IntList {
	mapped := IntList{}
	for _, x := range xs {
		mapped = mapped.Append(IntList{f(x)})
	}
	return mapped
}

func (xs IntList) Reverse() IntList {
	reversed := IntList{}
	for i, _ := range xs {
		reversed = reversed.Append(IntList{xs[xs.Length()-i-1]})
	}
	return reversed
}

func (xs IntList) Append(ys IntList) IntList {
	// return xs.append(xs, ys...)
	appended := make(IntList, xs.Length()+ys.Length())
	for i, x := range xs {
		appended[i] = x
	}
	for i, x := range ys {
		appended[xs.Length()+i] = x
	}
	return appended
}

func (xs IntList) Concat(xss []IntList) IntList {
	for _, x := range xss {
		xs = xs.Append(x)
	}
	return xs
}
