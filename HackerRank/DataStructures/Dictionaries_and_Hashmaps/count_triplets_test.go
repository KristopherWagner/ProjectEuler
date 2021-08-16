package Dictionaries_and_Hashmaps

import "testing"

func TestCountTriplets(t *testing.T) {
	t.Run("Should handle test case 0", testCTSampleCaseZero)
	t.Run("Should handle test case 1", testCTSampleCaseOne)
	t.Run("Should handle test case 2", testCTSampleCaseTwo)
	t.Run("Should handle all 1s", testCTCaseTwo)
	t.Run("Should handle all 1237", testCTCaseThree)
}

func testCTSampleCaseZero(t *testing.T) {
	arr := []int64{1, 2, 2, 4}
	count := CountTriplets(arr, 2)
	if count != 2 {
		t.Errorf("Expected 2, got %d\n", count)
	}
}

func testCTSampleCaseOne(t *testing.T) {
	arr := []int64{1, 3, 9, 9, 27, 81}
	count := CountTriplets(arr, 3)
	if count != 6 {
		t.Errorf("Expected 6, got %d\n", count)
	}
}

func testCTSampleCaseTwo(t *testing.T) {
	arr := []int64{1, 5, 5, 25, 125}
	count := CountTriplets(arr, 5)
	if count != 4 {
		t.Errorf("Expected 4, got %d\n", count)
	}
}

func testCTCaseTwo(t *testing.T) {
	arr := make([]int64, 100)
	for i := 0; i < 100; i++ {
		arr[i] = 1
	}

	count := CountTriplets(arr, 1)
	if count != 161700 {
		t.Errorf("Expected 161700, got %d\n", count)
	}
}

func testCTCaseThree(t *testing.T) {
	arr := make([]int64, 100000)
	for i := 0; i < 100000; i++ {
		arr[i] = 1
	}

	count := CountTriplets(arr, 1)
	if count != 166661666700000 {
		t.Errorf("Expected 166661666700000, got %d\n", count)
	}
}
