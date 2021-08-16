// https://www.hackerrank.com/challenges/count-triplets-1/problem
package Dictionaries_and_Hashmaps

func CountTriplets(arr []int64, r int64) (count int64) {
	dict := make(map[int64]int64)
	pairs := make(map[int64]int64)

	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		count += pairs[cur]
		pairs[cur*r] += dict[cur]
		dict[cur*r] += 1
	}
	return
}
