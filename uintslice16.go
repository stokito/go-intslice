// SPDX-License-Identifier: 0BSD
package intslice

type UInt16Slice []uint16

func (a UInt16Slice) BinSearch(key uint16) int {
	low := 0
	high := len(a) - 1
	for low <= high {
		mid := low + ((high - low) / 2)
		midVal := a[mid]
		if midVal < key {
			low = mid + 1
		} else if midVal > key {
			high = mid - 1
		} else {
			return mid // found
		}
	}
	return -(low + 1) // not found, return negative pos
}

func (a UInt16Slice) Contains(key uint16) bool {
	return a.BinSearch(key) >= 0
}

func (a *UInt16Slice) Insert(idx int, val uint16) {
	if len(*a) == idx {
		// just append to end
		*a = append(*a, val)
		return
	}
	// if idx < len(a) we need to insert
	*a = append((*a)[:idx+1], (*a)[idx:]...)
	(*a)[idx] = val
}

func (a *UInt16Slice) Remove(idx int) {
	*a = append((*a)[:idx], (*a)[idx+1:]...)
}

func (a *UInt16Slice) AddOrRemove(remove bool, val uint16) {
	idx := a.BinSearch(val)
	if remove {
		// if a contains the val
		if idx >= 0 {
			a.Remove(idx)
		}
	} else {
		// if a not contains yet the val
		if idx < 0 {
			a.Insert(-idx-1, val)
		}
	}
}
