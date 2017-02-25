/***divide two integer without using multiply, divide operation***/
#include <iostream>
//approx integer max is 2B
#define INT_MAX 2000000000

using namespace std;

int divide_int(int dividend, int divisor)
{
  if(divisor == 0)
    return INT_MAX;

  int result=0;
  int temp; //temporary varibale
  int factor=1; //keeps track of how many times divisior is substracted 
  
  //determine the sign of the quotient
  bool sign = ((dividend<0)&&(divisor>0))||((dividend>0)&&(divisor<0));
  if(dividend<0)
    dividend = -1*dividend;
  if(divisor<0)
    divisor= -1*divisor;
  //Note: include equal comparison since remainder could be zero 
  while (dividend>=divisor)
    {
      // assign temp to divisor, fcator 1 
      temp = divisor;
      factor = 1;
      
      //subtract temp from dividend and shift every step
      while(temp <= dividend)
	{
	  //cout << "factor: " << factor << " temp: "<< temp <<endl;
	  factor = (factor<<1); //kep track of many times
	  temp = (temp<<1); //multiply divisor by two
	}
      
      //update result
      result += (factor>>1);
      temp = temp>>1;
      dividend -= temp;
      
    }
  //optional
  cout << "remainder = " << dividend << endl;
  return (sign?(-1*result):result);
  
}

int main()
{
  int a,b;
  cout << "Enter two integers" << endl;
  cin >> a >> b;

  cout << "Result = " << divide_int(a, b) << endl;
}





