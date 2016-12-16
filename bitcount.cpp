#include <iostream>
#include <vector>


using namespace std;

class Solution {
public:
    vector<int> countBits(int num) {
      vector<int> bitCntArr;
      //check if the number is negative
      if (num < 0)
	{
	  cerr << "Error Case: Num negative";
	  return bitCntArr;
	}
      //define the return vector
      
      vector<int>::iterator it = bitCntArr.begin();
      //fill the vector array
      bitCntArr.insert(it, 0);
      int count=0;
      for (int n=1; n<=num; n++)
	{
	  count = count_bits(n);
	  bitCntArr.push_back(count);
	}
      return bitCntArr;
    }

  //internal function
  int count_bits(int n)
  {
    //cout << "Inside count_bits " << n << endl;
    int count = 0;
    while ((n&(n-1))!=0)
      {
	count++;
	n&=n-1;
      }
    count+=1;
    return count;
  }
};


int main()
{
  int num;
  Solution s;
  cout << "Enter the value " << endl;
  cin >> num;
  vector<int> Output = s.countBits(num);
  for (vector<int>::const_iterator i=Output.begin(); i!=Output.end(); i++)
    {
      cout << *i << " ";
    }
  cout << endl;

}
