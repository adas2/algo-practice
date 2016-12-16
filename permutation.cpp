#include <iostream>
#include <sstream>
#include <vector>
#include <string>
//#include <stdlib.h>

using namespace std;

class Solution {
public:
  string getPermutation(int n, int k) {
    string s;
    ostringstream t;
    //create the string
    for (int i=1 ; i <=n; i++)
      {
	t.str("");
	t.clear();
	t << i ;
	s.append(t.str());
      }
    //recursively permutate the string
    int cnt;
    //string s_out;
    findPermutation(s, 0, n-1, cnt, k);
    return s; 
  }
  //permutaion of string str starting at index i and ending at index n 
  void findPermutation(string &str, int i, int n, int &cnt, int k)
  {
    //int cnt=0;
    if (i==n)
      {
	cnt++;
	if(k==cnt)
	  {
	    cout << str << " " << cnt << endl;
	  }
	
      }

    for(int p=i; p<=n; p++)
      {
	swap(str[i], str[p]);
	findPermutation(str, i+1, n, cnt, k);
	swap(str[i], str[p]);
      }
    
  }

  void swap(char &x, char &y)
  {
    char temp;
    temp = x;
    x = y;
    y = temp;
  }

};

int main()
{
  int n, k;
  cout << "Enter the value of n and k:" << endl;
  cin >> n ;
  cin >> k;
  Solution s;
  cout << "The original string is: " << s.getPermutation(n,k) << endl;
}
