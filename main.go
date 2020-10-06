package main

// Try to connect each neighbor numbers based orders in org;
// Meanwhile, if any follow situation occurs, return False
//   if current visiting number num is not in org
//   if there is a cycle or mis-order (they are basically same thing for this question), e.g.
//     [[1,2], [2,1]], there is a mis-order and cycle in this sequences
// If all nodes in org has been visited and connected in order, return True

type Set map[int]bool

func sequenceReconstruction(org []int, seqs [][]int) bool {
	visited := make(Set)
	init := make(Set)

	d := make(map[int]int)
	for i, v := range org {
		d[v] = i
		init[v] = true
	}

	var prev_idx = -1
	var prev_num = -1
	for _, nums := range seqs {
		prev_idx = prev_num
		for _, num := range nums {
			if init[num] == false {
				return false
			}
			cur_idx := d[num]

			if prev_idx+1 == cur_idx && visited[num] == false {
				visited[num] = true
			}

			if prev_idx >= cur_idx {
				return false
			}

			prev_num, prev_idx = num, cur_idx
		} // for nums
		prev_idx = -1
		prev_num = -1
	} // for seqs
	return len(visited) == len(org)
} // func()
