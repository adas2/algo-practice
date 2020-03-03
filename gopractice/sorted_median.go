package practice

import "fmt"

func findSortedMedian(arr1, arr2 []int) float32 {
    // case 1: odd length
    if (len(arr1)+len(arr2))%2 != 0 {
        fmt.Printf("Median is the %d index element\n", (len(arr1)+len(arr2))/2)
    } else {
        fmt.Printf("Median is the %d and %d index element\n", (len(arr1)+len(arr2))/2-1, (len(arr1)+len(arr2))/2)
    }

    var minIdx, maxIdx, i, j int
    var median float32
    // var cond1, cond2 bool
    // choose the smaller array
    if len(arr1) > len(arr2) {
        temp := arr1
        arr1 = arr2
        arr2 = temp
    }

    fmt.Println(arr1)
    fmt.Println(arr2)

    // start from the start
    minIdx = 0
    maxIdx = len(arr1)
    //divide arr1 and arr2 into two halves s.t. all elements in left half > elements in right half
    // arr1[0:i], arr2[0:j] and arr1[i:], arr2[j:]
    // terminate when condition is met
    for minIdx <= maxIdx {
        i = (minIdx + maxIdx) / 2
        j = (len(arr1)+len(arr2)+1)/2 - i
        fmt.Println(minIdx, maxIdx, i, j)
        if j > 0 && i < len(arr1) && arr2[j-1] > arr1[i] {
            // move right in arr1
            minIdx = i + 1
        } else if i > 0 && j < len(arr2) && arr1[i-1] > arr2[j] {
            // move left in arr1
            maxIdx = i
        } else {
            // desired cond

            // two cases even and odd
            if (len(arr1)+len(arr2))%2 == 0 {
                if i == 0 {
                    median = float32((arr2[j-1] + arr2[j]) / 2)
                } else if j == 0 {
                    median = float32((arr1[i-1] + arr1[i])) / 2
                } else {
                    median = float32((max(arr1[i-1], arr2[j-1]) + min(arr1[i], arr2[j])) / 2)
                }
            } else {
                if i == 0 {
                    median = float32(arr2[j-1])
                } else if j == 0 {
                    median = float32(arr1[i-1])
                } else {
                    median = float32(max(arr1[i-1], arr2[j-1]))
                }
            }
            break
        }
    }

    return median
}

/*
func max(a,b int) int{
    if a<b{
        return b
    }
    return a
}
*/

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
