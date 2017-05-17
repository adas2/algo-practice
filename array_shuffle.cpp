#include <iostream>
#include <vector>
#include <cstdlib>

using namespace std;

class Solution {
public:
    Solution(vector<int> nums) {
      arr=nums;
      original_arr=nums;
      len = nums.size();
    }
    
    /** Resets the array to its original configuration and return it. */
    vector<int> reset() {
      arr = original_arr;
      return arr;
    }
    
    /** Returns a random shuffling of the array. */
    vector<int> shuffle() {
      //generate a Knuth shuffle
      int n=len;
      while(n>1)
	{
	  //1. generate random index between 0 and n-1 i.e. size of array
	  int idx = rand()%n;
	  cout << n << " " << idx << endl;
	  //2. swap the last element with the current generated index
	  swap(arr[idx], arr[n-1]);
	  //3. reduce the size of the array
	  n--;
	}
      return this->arr;
    }

  void swap(int &x, int &y)
  {
    int temp;
    temp = x;
    x = y;
    y = temp;
  }

  void print()
  {
    for(auto& x: arr)
      cout << x << " ";
    cout << endl;
  }


private:
  vector<int> arr;
  vector<int> original_arr;
  int len; //size
};


int main()
{
  //vector <int> test = {1,2,3,4,5,6};
  vector <int> test = {0,-12893,128384};
  Solution s(test);
  s.print();
  s.shuffle();
  s.print();
  s.shuffle();
  s.print();
  s.reset();
  s.print();
  s.shuffle();
  s.print();
  
  return 1;					
}
