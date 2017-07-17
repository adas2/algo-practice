#include <iostream>
#include <vector>


using namespace std;

int find_max_sum_subarray(vector<int> arr)
{
  //edge cases
  if (arr.size()==1)
    return arr[0];
  if (arr.empty())
    return -1;

  //define current_sum, max_sum, and max_index (optional)
  int curr_sum = arr[0];
  int max_sum = curr_sum;
  int max_index_start = 0; //index which starts the max sum
  int max_index_end = 0; //index which ends the max sum

  //iterate through the array 
  for(int i=0; i<arr.size()-1; ++i )
    {
      //case when running sum + next element is lesser than the element
      if((curr_sum+arr[i+1])<arr[i+1])
	{
	  //the left side may be discarded before the element
	  curr_sum = arr[i+1];
	  //mark the new start index
	  max_index_start = i+1;
	}
      else
	{ //running sum increases with the new element summed up
	  curr_sum += arr[i+1];
	}
      //update max_sum and end_index 
      if(max_sum < curr_sum){
	max_sum = curr_sum;
	max_index_end = i+1;
      }
    }

  cout << "Start index: " << max_index_start << endl;
  cout << "End index: " << max_index_end << endl;

  return max_sum;
  
}

int main()
{
  vector<int> my_arr = {1,-2,4,-1,7,3,-4,2};
  vector<int> my_arr2 = {-256, 893};
  
  cout << "Max sum subarray: \n";
  cout << find_max_sum_subarray(my_arr) << endl;
}
