package practice

import (
    "fmt"
    "sort"
)

// sort version O(nlogn)
func topKFrequent(nums []int, k int) []int {
    //collect frequency
    cMap := map[int]int{}
    for _, v := range nums {
        cMap[v] += 1
    }

    // create array of tuplets with <val, freq>
    freq := make([][]int, len(cMap))
    for i := range freq {
        freq[i] = make([]int, 2) //val, frequency
    }
    // freq := []int{}
    i := 0
    for k, v := range cMap {
        freq[i][0] = k
        freq[i][1] = v
        i++
    }

    // do reverse sort by freq
    sort.SliceStable(freq, func(i, j int) bool {
        return freq[i][1] > freq[j][1]
    })

    fmt.Println(freq)

    // choose the top k frequencies
    out := []int{}
    for i := 0; i < k; i++ {
        // key := freq[len(freq)-i]
        out = append(out, freq[i][0])
    }

    return out
}

// quickselect version ~ O(n). DEBUG
func topKFrequent2(nums []int, k int) []int {
    //collect frequency
    cMap := map[int]int{}
    for _, v := range nums {
        cMap[v] += 1
    }

    // create array of tuplets with <val, freq>
    freq := make([][]int, len(cMap))
    for i := range freq {
        freq[i] = make([]int, 2) //val, frequency
    }
    // freq := []int{}
    i := 0
    for k, v := range cMap {
        freq[i][0] = k
        freq[i][1] = v
        i++
    }

    out := []int{}
    // do randomized quickselect to do linear time
    pivot := findKthFreqIndex(freq, 0, len(freq)-1, k)
    if pivot == -1 {
        fmt.Printf("error pivot %d\n", pivot)
        return out
    }

    // choose the top k frequencies
    for i := 0; i < k; i++ {
        // key := freq[len(freq)-i]
        out = append(out, freq[pivot+i][0])
    }

    return out
}

func partition(arr *[][]int, s, e int) int {
    // choose lat element as pivot
    pivot := (*arr)[e][1]
    i := s // first large element
    // swap
    for j := s; j <= e; j++ {
        if (*arr)[j][1] < pivot {
            // swap(arr, i, j)
            (*arr)[i][1], (*arr)[j][1] = (*arr)[j][1], (*arr)[i][1]
            i++
        }
    }
    // place pivot in right position
    // swap(arr, i, r)
    (*arr)[i][1], (*arr)[e][1] = (*arr)[e][1], (*arr)[i][1]
    return i
}

// alterative version
func findKthFreqIndex(arr *[][]int, s, e, k int) int {
    // kth least frequent --> find (len-k) leats ferquent
    //  add everything to the right of the index
    target := len(arr) - k

    // iterative version
    for s <= e {
        p := partition(arr, s, e)

        if p == target {
            // this is the k th most frequent
            return p
        } else if p > target {
            // to the left side
            e = p - 1
        } else { // p < target-1
            // to the right side
            s = p + 1
            // target = target - p - 1
        }

    }
    // kth not found
    return -1
}
