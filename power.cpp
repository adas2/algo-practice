#include <iostream>

using namespace std;

class Solution {
public:
  double myPow(double x, int n) {
    if(n==0)
      return 1;
    unsigned int abs_n = abs(n);
 
    int count = 0;
    long double result = 1, temp;
    unsigned int mask =1; // masking the bits in the exponent
    while (mask !=0){
      //first ietration is 2^0; x*(2^0)=x;
      if(count==0)
	temp = x;
      else
	temp *= temp;
      if(mask & abs_n) //when the bit is non zero int he mask location
	{
	  result *= temp;
	}
      //left shift mask by 1 bit position
      mask = mask << 1;
      count++;
    }
    if(n<0)
      return (1/result);
    return result;
  }
};


int main()
{
  double base;
  int exp;
  cout << "Enter the base and exponent" << endl;
  cin >> base >> exp;
  //cout << sizeof(double) << endl;
  Solution s;
  cout << "The result of base raised to power exponent: " << s.myPow(base, exp) << endl;
  return 0;
}
