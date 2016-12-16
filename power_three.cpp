#include <math.h>
#include <iostream>

using namespace std;

class Solution {
public:
  bool isPowerOfThree(int num) {
    if(num==0)
      return false;
    
    while(num%3==0 && num>1)
      {
	num=num/3;
      }
    //if you remian with only 1
    if(num==1)
      return true;
    else
      return false;

   
  }
};

int main()
{
  int num;
  cout << "Enter num" << endl;
  cin >> num;
  Solution s;
  if(s.isPowerOfThree(num))
    cout << "True" << endl;
  else
    cout << "False" << endl;
}
