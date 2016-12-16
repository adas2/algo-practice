#include <iostream>


using namespace std;

class Solution{
public:
  int find_gcd(int a, int b)
  {
    //error case
    if (a==b && b==0)
      return -1;

    if (b==0)
      return a;
    else
      return find_gcd(b, (a%b));
  }
};


int main()
{
  int x, y, gcd=0;
  Solution s;
  cout << "Enter two non-zero integers" << endl;
  cin >> x >> y;
  cout << "GCD for the two numbers: " << s.find_gcd(x,y) << endl;
  if(s.find_gcd(x,y) == 1)
    cout << "The numbers are mutually prime too"  << endl;
}
