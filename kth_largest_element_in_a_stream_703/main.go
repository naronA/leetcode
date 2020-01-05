package main

type KthLargest struct {
	MinHeap []int
	Size    int
	MaxSize int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{MaxSize: k, Size: 0}

	for _, e := range nums {
		kth.Add(e)
	}

	return kth
}

func (k *KthLargest) Add(val int) int {
	if k.Size < k.MaxSize {
		k.MinHeap = append(k.MinHeap, val)
		k.Size++
		k.HeapfyUp(k.Size - 1)
	} else if val > k.MinHeap[0] {
		k.MinHeap[0] = val
		k.HeapfyDown(0)
	}

	return k.MinHeap[0]
}

// 2分木のノードを交換する
func (k *KthLargest) Swap(a, b int) {
	k.MinHeap[a], k.MinHeap[b] = k.MinHeap[b], k.MinHeap[a]
}

func (k *KthLargest) HeapfyUp(index int) {
	for index > 0 && k.MinHeap[index] < k.MinHeap[(index-1)/2] {
		k.Swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func (k *KthLargest) HeapfyDown(index int) {
	for 2*index+1 < k.MaxSize {
		older := 2*index + 1

		if 2*index+2 < k.MaxSize && k.MinHeap[2*index+2] < k.MinHeap[older] {
			older = 2*index + 2
		}

		if k.MinHeap[index] < k.MinHeap[older] {
			break
		}

		k.Swap(index, older)
		index = older
	}
}
