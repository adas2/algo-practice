#include <iostream>
#include <string>
#include <vector>

using namespace std;

//str 1 is the text and str is subsequence
int find_num_subsequence(string str1, string str2)
{
  int n = str1.size();
  int m = str2.size();
 
  //initialize
  vector<vector<int> > cnt(n+1);
  for (int i=0; i<=n; ++i)
    cnt[i].resize(m+1,0);

  //itareate through match lengths
  for (int i=0; i<=n; ++i)
    for(int j=0; j<=m; ++j)
      {
	//if str2 has been all matched, count as 1 
	if(j==0)
	  cnt[i][j]=1;
	else if(i==0)
	  cnt[i][j]=0;
	//case 1: match of last character e.g. (abca, a)
	else if(str1[i-1]==str2[j-1])
	  cnt[i][j]=cnt[i-1][j-1]+cnt[i-1][j]; //add (abca, NULL) + (abc,a)
	//case 2: mismatch of character e.g. (abca, ab) 
	else 
	  cnt[i][j]=cnt[i-1][j]; //add (abc, ab)
      }
  
  return cnt[n][m];
}

int main()
{
  string s1 = "abcpbcqc";
  string s2 = "abc";
  cout << find_num_subsequence(s1, s2) << endl;
  
  return 0;
}
