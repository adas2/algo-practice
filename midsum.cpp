#include <iostream>
#include <stdlib.h>
#include <vector>

using namespace std;

int find_mid_sum(vector<int> arr)
{
  int sum=0, lsum=0;
  //int i,j;

  //calculate sum
  for (int i=0; i<arr.size() ; ++i)
    {
      sum += arr[i];
    }
  cout << "Array Size: " << arr.size() << " Sum: " << sum << endl;

  //calculate partial sum for left side of array 
  for (int j=0; j<arr.size()-1; ++j )
    {
      lsum += arr[j];
      cout << "lsum: " << lsum << endl;
      if((2*lsum + arr[j+1])==sum)
	{
	  //match found, check if not last index
	  //not necessary to check last index since 2*lsum + a[last_index] cannot be sum
	  if((j+1)!=arr.size()-1)
	    return (j+1);
	}
    }
  //no match found
  return -1;

}

int main()
{
  vector<int> arr;
  arr.push_back(7);
  arr.push_back(-2);
  arr.push_back(5);
  arr.push_back(1);
  arr.push_back(2);
  arr.push_back(8);

  cout << "Mid sum Index: " << find_mid_sum(arr) << endl; 

}
