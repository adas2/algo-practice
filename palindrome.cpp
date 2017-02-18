#include <iostream>
#include <vector>
#include <string>
#include <ctype.h>

using namespace std;

class Solution{
public:
  //simple palindrome
  bool findPalindrome(const string s)
  {
    int i = 0;
    int j = s.length()-1;
    
    while (j>i)
      {
	if(s[i]!=s[j])
	  {
	    return false;
	  }
	i++;
	j--;
      }
    return true;
  }

  //alphabetic palindrome with upper and lower case support 
  bool findPalindromeAlphabetic(const string str)
  {
    if(!str.length())
      return false;

    int i=0;
    int j=str.size()-1;
    
    while(i<j)
      {
	while(!isalpha(str[i]) && i<j)
	  i++;
	while(!isalpha(str[j]) && i<j)
	  j--;
	//cout << i << str[i] << " : " << str[j] << endl;
	if(tolower(str[i])!=tolower(str[j]))
	  return false;
	
	i++; j--;
      }
    return true;
  }
  
};
int main()
{
  string s;
  cout << "Enter the line/string: " << endl;
  //cin >> s ;
  getline(cin, s);
  Solution r;
  if(r.findPalindromeAlphabetic(s))
    cout << "Palindrome!" << endl;
  else
    cout << "Not an Palindrome" << endl;  
}
