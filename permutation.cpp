#include <iostream>
#include <sstream>
#include <vector>
#include <string>
//#include <stdlib.h>

//#define val(x) int(x)-'1'

using namespace std;

class Solution {
public:
  string createStringFromNumber(int n)
  {
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
    return s;
  }
  void  getPermutation(int n, int k) {

    string s =createStringFromNumber(n);
    //recursively permutate the string
    int cnt=0;
    //ieteration for backtracking
    int i=0;   
    //create empty result string
    string result;
    result.resize(n,' ');

    //string s_out;
    //findPermutation(s, 0, n-1, cnt, k);
    findLexicographicPermutaion(s, k, cnt, result,i );
    //return result; 
  }

  //print perm in lexical order
  void findLexicographicPermutaion(string &s, int k, int &cnt, string &temp_str, int iter)
  {
    if (iter==s.size() )
      {
	//end of string reached print perm
	cnt++;
	if(cnt==k)
	{
		cout << temp_str << endl;
		return;
	}
      }
    //else start next iterations
    
    //contrsuct candidates by traversing the partial temp_str to see which elemnts are not added
    vector<bool> present;
    present.resize(s.size(),false);
 
    //cout << "iter: "  << iter << endl;
    //T=O(k) operation si.e. 1+2+3+...+n = O(n^2) total
    for(int i=0; i<iter; ++i){
      //mark the ones which are already present in partial soln temp_str
      //cout << temp_str[i] << " " << i << endl;
      present[int(temp_str[i])-int('1')]=true;
    }
    //try recursive backtracking for each candidate
    for(int j=0; j<s.size(); ++j)
      {
	if(!present[j])
	  {
	    temp_str[iter]=s[j];
	    //backtrack by increasing the itartion to one more count
	    findLexicographicPermutaion(s,k,cnt, temp_str,iter+1);
	  }
      }

  }


  //NOTE: this is not in lexicographic order
  //permutaion of string str starting at index i and ending at index n 
  void findPermutation(string &str, int i, int n, int &cnt, int k)
  {
    //int cnt=0;
    if (i==n)
      {
	cnt++;
	if(k==cnt)
	  {
	    cout << cnt << "-th permutation of original string " << str << endl;
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
  string original = s.createStringFromNumber(n);
  cout << "original string: " << original << endl;
 
  /***
  //ieteration for backtracking
  int i=0;   
  //create empty result string
  string result;
  result.resize(n,' ');
  s.findLexicographicPermutaion(original,k,result,i);
***/

  cout << "The  kth permutaion  string is: " << endl;
  s.getPermutation(n,k) ;
}
