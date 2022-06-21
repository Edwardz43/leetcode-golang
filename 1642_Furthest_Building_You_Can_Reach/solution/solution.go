package solution

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
Runtime: 135 ms, faster than 47.37% of Go online submissions for Furthest Building You Can Reach.
Memory Usage: 10.3 MB, less than 7.89% of Go online submissions for Furthest Building You Can Reach.

*/
func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < len(heights)-1; i++ {
		climb := heights[i+1] - heights[i]
		if climb > 0 {
			heap.Push(h, climb)
		}
		if h.Len() > ladders {
			bricks -= heap.Pop(h).(int)
		}
		if bricks < 0 {
			return i
		}
	}
	return len(heights) - 1
}

/*
Runtime: 212 ms, faster than 13.16% of Go online submissions for Furthest Building You Can Reach.
Memory Usage: 9.8 MB, less than 28.95% of Go online submissions for Furthest Building You Can Reach.
*/
func furthestBuilding_v2(heights []int, bricks int, ladders int) int {

	// reachedLast := false
	var i = 0
	h := &IntHeap{}
	heap.Init(h)
	// allocate the ladders first and keep on checking for the bricks allocation if the min
	// of ladder is less than or equal to current bricks allocation to be done, reduce the bricks.
	for ; i < len(heights)-1; i++ {
		if heights[i] >= heights[i+1] {
			continue
		}
		if heights[i] < heights[i+1] && ladders >= 1 {
			heap.Push(h, heights[i+1]-heights[i])
			ladders--
			continue
		}
		if heights[i] < heights[i+1] {
			popped := false
			var x = 0
			if h.Len() > 0 {
				popped = true
				x = heap.Pop(h).(int)
			}

			if popped && x < (heights[i+1]-heights[i]) && bricks >= x {
				bricks = bricks - x
				heap.Push(h, heights[i+1]-heights[i])
			} else if heights[i+1]-heights[i] <= bricks {
				bricks = bricks - (heights[i+1] - heights[i])
				if popped {
					heap.Push(h, x)
				}

			} else {
				break
			}
			continue
		}

		// reachedLast = true
		break
	}

	return i
}
