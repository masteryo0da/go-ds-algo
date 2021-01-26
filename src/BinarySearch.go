package src

type BinarySearch struct {
	nums [5]int
}

func NewBinarySearch(nums [5]int) BinarySearch {
	return BinarySearch{nums: nums}
}

func (b *BinarySearch) Find(num int) int {
	start := 0
	end := len(b.nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if b.nums[mid] == num {
			return mid
		} else if b.nums[mid] > num {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
