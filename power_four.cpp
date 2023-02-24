#include <math.h>
#include <iostream>

using namespace std;

class Solution
{
public:
  bool isPowerOfFour(int num)
  {
    // check if (x^4) --> if y^2n i.e.
    // iff number is a perfect square and
    // number is a power of 2
    double sq_num = sqrt(num);
    if (num != 0 && sq_num == floor(sq_num) && (num & (num - 1)) == 0)
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
  if (s.isPowerOfFour(num))
    cout << "True" << endl;
  else
    cout << "False" << endl;
}
