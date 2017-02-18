/***** Find the first index i such that sum of elemnts from ith idnex to windows size id equal to target  ***/

#include <iostream>
#include <vector>

using namespace std;

int find_first_window_sum(vector<int> arr, int window, int target)
{
  //initialize sum for window
  int sum =0;
  //int curr=0;
  for (int i=0; i<window; ++i)
    sum += arr[i];
  //iterate through the array from 0 to last possible position 
  int j;
  for (j=0; j<(arr.size()-window); ++j)
    {
	//first index will be covered here
      if(sum==target)
	return j;
      //update sum with the next element
      sum = sum + arr[j+window] - arr[j];
    }
  //IMPORTANT: if the (last index-window) is a match 
  if(sum==target)
	  return j;

  //else no index is found 
  return -1;
     
}

int main()
{
  vector<int> arr = {7,-3, 3,1,2,3,1,2,-5,9};
  cout << "First index with window sum equal to target:" << endl;
  cout << find_first_window_sum(arr, 3, 6) << endl;

}




