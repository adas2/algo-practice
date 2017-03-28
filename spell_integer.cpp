#include <iostream>
//#include <unordered_map>
#include <climits>
#include <vector>

void print_digits(int num);

using namespace std;

//global map of one digit strings
vector<string> spell = {
  "zero", "one", "two", "three", "four", "five",
  "six", "seven", "eight", "nine", "ten",
  "eleven", "twleve", "thirteen", "fourteen",
  "fifeteen", "sizteen", "sevemteen", "eighteen",
  "nineteen"
}; 

vector<string> tens = {
  "", "", "twenty", "thirty", "forty", "fifty", 
  "sixty", "seventy", "eighty", "ninety"
};

//recrsive function to call
//starts with billion i.e. 10^9 as factor
void print_number(int num)
{
  cout << "Num: " << num << endl; 
  if (num < 0){
    cout << "Negative " << endl;
    num *= -1;
    //return;
  }
   
  //max possible value is 2 billion+offset; hence start with billion
  //if it has digit in billionth place
  int factor = 1000000000;
  int digit = num/factor; 

  //while (num >= 1000)
  //{
  if(digit>0)
    {
      print_digits(digit);
      cout << " billion ";
      num -= (digit*factor);
    }
  //move to the next 3 digits
  factor /= 1000;
  digit = num/factor;
  if(digit>0)
    {
      print_digits(digit);
      cout << " million ";
      num -= (digit*factor);
    }
  factor /= 1000;
  digit = num/factor;
  if(digit>0)
    {
      print_digits(digit);
      cout << " thousand ";
      num -= (digit*factor);
    }
  //}
  //if small handle by print_digits utlity
  if(num<1000){
    print_digits(num);
    cout << endl << endl;
  }
  
  return;
}  

void print_digits(int num)
{
  if (num < 0 || num > 999){
    cout << "Error out of range" << endl;
    return;
  }
  /**
  if (num < 20){
    cout << spell[num] << " ";
    cout << endl;
    return;
  }
  **/
  //else case when 2 or 3 digits
  int d = num/100;
  if(d>0)
    {
      cout << spell[d] << " hundred ";
      num -= d*100;
    }
  if(num/10 > 1){
    cout << tens[num/10]; // << spell[num%10] ; 
    if((num%10)!=0)
      cout << " " << spell[num%10];
  }
  else
    cout << spell[num];
}

int main()
{
  int n=1205670098;
  int n1=345;
  int n2=007;
  int n3=-234567;
  int n4=1000000001;

  print_number(n);
  print_number(n1);
  print_number(n2);
  print_number(n3);
  print_number(n4);
}
