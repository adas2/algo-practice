#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int wiggleMaxLength(vector<int>& nums) {
      int count =0;
      int p=1, i=0;
      bool incr ;
      if(nums[1]-nums[0] > 0)
	incr = false;
      else
	incr = true;

      if(nums.size()==0)
	return count;

      //traverse the vector with greedy approach
      while (p < nums.size())
	{
	  //int diff = nums[p]-nums[i];

	  cout << "processing i: " << i << " j: " << p;
	  if((nums[p]-nums[i])>0 && incr==false)
	    {
	      //increasing after decreasing
	      count++; 
	      //i=p;
	      incr = true; //increasing subsequence
	    }
	  else if ((nums[p]-nums[i])<0 && incr==true)
	    {
	      //decreasing subsequence after increasing
	      count++;
	      //i=p;
	      incr = false; //decreasing subs
	    }
	  /***
	  else if ((nums[p]-nums[i]) >0 && incr=true)
	    {
	      //continually increasing
	      i=p
	    }
	  ***/
	  i=p;
	  p++;

	  cout << " Count: " << count << endl;;
	}

      return count+1;
    }
};





int main()
{
  Solution s;
  
  vector<int> V = {1,17,5,10,13,15,10,5,16,8};
  //vector<int> V = {3, 3, 2, 5};
  cout << "Lenght of wiggle sequence: " << endl;
  cout << s.wiggleMaxLength(V) << endl;
  
}
