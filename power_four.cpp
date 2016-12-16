#include <math.h>
#include <iostream>

using namespace std;

class Solution {
public:
  bool isPowerOfFour(int num) {
    //if(num==0)
    //return false;
    //if(num==1)
    //return true;
    //power of four iff sqrt is a power of 2 i.e. even
    double sq_num = sqrt(num);
    if(num!=0 && sq_num == floor(sq_num) && (num & (num-1))==0)
      return true;
     
    return false;
  }
};

int main()
{
  int num;
  cout << "Enter num" << endl;
  cin >> num;
  Solution s;
  if(s.isPowerOfFour(num))
    cout << "True" << endl;
  else
    cout << "False" << endl;
}
