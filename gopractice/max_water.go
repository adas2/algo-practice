package practice

// Max water volume problem (leetcode) [lc # 11]

func maxArea(height []int) int {

	maxVol, currVol := 0, 0
	i, j := 0, len(height)-1

	for i < j {
		// cal current volume by (i,j)
		currVol = (j - i) * min(height[i], height[j])
		if currVol > maxVol {
			maxVol = currVol
		}
		// move the smaller height index to the next closest bar
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}

	return maxVol
}

// logic:
// e.g. height = [1,8,6,2,5,4,8,3,7]
// start with two pointers at i=0 and j= len(height)-1
// curr_vol = (j-i)*min(height[i], height[j])
// if height[i] <= height[j] --> move i++
// esle move j--
// this is because there is a possibility of getting a larger bar
// by moving away from smaller bar, as smaller bar determines the vol
// space between the bar reduces as we move i or j;
// so idea is to move towards are larger i or j (tricky)
