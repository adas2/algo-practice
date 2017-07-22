/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

#include <iostream>
#include <string>
//#include <algorithm>

#define max(a,b) (a>b)?a:b
#define min(a,b) (a<b)?a:b

// String add(String num1, String num2)
// add("10", "11") -> "101"

using namespace std;

void reverse(string &s)
{
    if(s.size()>1)
    {
        int i=0, j=s.size()-1;
        while(i<j)
        {
            char temp=s[i];
            s[i]=s[j];
            s[j]=temp;
            i++; j--;
        }
    }
    return;
}

char add_bit(char a, char b, bool &carry)
{
    char result;
    if((a=='1') && (b=='1'))
    {
      result = (carry)?('1'):('0');
      carry=true;
    }
    
    //case2: 1+0 || 0+1 = 1
    if (((a=='0') && (b=='1') )||((b=='0') && (a=='1') ))
    {
      result = (carry)?('0'):('1');
      cout << result << endl;
      
    }
    //case3: 0+0 = 0
    if((a=='0') && (b=='0'))
    {
     result =  (carry)?('1'):('0');
     carry=false;
    }
    
    return result;

}

//Bitwise addition of two strings, assumes no sign char
string add(string num1, string num2)
{
  
  int n1= num1.size();
  int n2= num2.size();
  
  if(n1==0 || n2==0)
    return "Error";

  
  //append the smaller number with zeros
  if(n1<n2)
  {
      string s(n2-n1, '0'); //repeated 0's for MSB
      num1 = s+num1;
  }
  else
  {
      string s(n1-n2, '0'); //repeated 0's for MSB
      num2 = s+num2;
  }
  cout << num1 << endl;
  cout << num2 << endl;
  
  bool carry=false;
  
  string result=""; 
  //start from the end of the strings to 0
  for(int i=max(n1,n2); i>=0; --i )
  {
    //cout << add_bit(num1[i], num2[i], carry) << endl;
    result += add_bit(num1[i], num2[i], carry) ;
  }
    
  if(carry)
    result +='1';
  
  cout << result << endl;
  
  //reverse to bring MSB first
  reverse(result);
 
 
  return result;
   
}

//Runtime: O(max(m,n))

int main()
{
    string a="1011";
    string b="111";
    
    cout << "Sum = " << add(a,b) << endl;
                   
}
