// https://www.hackerrank.com/challenges/frequency-queries/problem
package Dictionaries_and_Hashmaps

func FreqQuery(queries [][]int32) (output []int32) {
	output = make([]int32, 0)
	entries := make(map[int32]int32)
	frequencies := make(map[int32]int32)

	for _, query := range queries {
		command, value := query[0], query[1]
		if command == 1 {
			frequencies[entries[value]]--
			entries[value]++
			frequencies[entries[value]]++
		} else if command == 2 {
			if entries[value] > 0 {
				frequencies[entries[value]]--
				entries[value]--
				frequencies[entries[value]]++
			}
		} else if command == 3 {
			if frequencies[value] > 0 {
				output = append(output, 1)
			} else {
				output = append(output, 0)
			}
		}
	}
	return
}
