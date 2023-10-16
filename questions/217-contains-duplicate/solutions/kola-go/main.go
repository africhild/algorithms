package main

func main() {

}

//	func containsDuplicate(nums []int) bool {
//		m := map[int]bool{}
//		for _, v := range nums {
//			if m[v] {
//				return true
//			}
//			m[v] = true
//		}
//		return false
//	}

// func containsDuplicate(nums []int) bool {
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] == nums[j] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// solution 2: Sort and compare
// func containsDuplicate(nums []int) bool {
// 	if len(nums) <= 1 {
// 		return false
// 	}
// 	sort.Ints(nums)
// 	for i := 0; i < len(nums)-1; i++ {
// 		if nums[i] == nums[i+1] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func containsDuplicate(nums []int) bool {
// 	set := newMultiSet()
// 	set.AddNums(nums)
// 	for i := 0; i < len(nums); i++ {
// 		if set.HasDuplicate(nums[i]) {
// 			return true
// 		}
// 	}
// 	return false
// }

// type multiSet struct {
// 	items map[int]int
// }

// func newMultiSet() *multiSet {
// 	return &multiSet{
// 		items: make(map[int]int),
// 	}
// }

// func (s *multiSet) AddNums(nums []int) {
// 	for _, num := range nums {
// 		s.AddNum(num)
// 	}
// }

// func (s *multiSet) AddNum(num int) {
// 	s.items[num]++
// }

// func (s *multiSet) HasDuplicate(num int) bool {
// 	ocr := s.OccurenceOf(num)
// 	return ocr > 1
// }

// func (s *multiSet) OccurenceOf(num int) int {
// 	ocr, ok := s.items[num]
// 	if !ok {
// 		return 0
// 	}
// 	return ocr

// }

func containsDuplicate(nums []int) bool {
	// set := map[int]struct{}{}
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, hasNum := set[num]; hasNum {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}

