package solution

import (
	"container/heap"
	"sort"
)

type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
func (h *MaxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }

/*
Runtime: 155 ms, faster than 82.35% of Go online submissions for Course Schedule III.
Memory Usage: 8.3 MB, less than 5.88% of Go online submissions for Course Schedule III.
*/
func scheduleCourse(courses [][]int) int {
	day, maxHeap := 0, &MaxHeap{}
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	for _, course := range courses {
		heap.Push(maxHeap, course[0])
		day += course[0]

		if day > course[1] {
			day -= heap.Pop(maxHeap).(int)
		}

		if day > course[1] {
			return maxHeap.Len()
		}
	}

	return maxHeap.Len()
}

func scheduleCourseV2(courses [][]int) int {

	// sort the classes so that classes that end later are in the list
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	// keep a max heap of times to complete classes we've already chosen
	// keep track of the total time
	chosenClasses := MaxHeap{}
	var time int

	// INVARIANT: we have enough time to take all chosen classes
	// BASE CASE: we start with no classes, which takes no time
	for _, this := range courses {

		// assume we take this class
		heap.Push(&chosenClasses, this[0])
		time += this[0]

		// if we're taking too much time: INVARIANT violated
		if time > this[1] {
			// remove the longest class from our list: INVARIANT restored
			time -= heap.Pop(&chosenClasses).(int)
		}

	}

	// return the number of classes on the list
	return len(chosenClasses)

}
