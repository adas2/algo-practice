#include <iostream>
#include <string>
#include <vector>
#include <unordered_set>

#define MAX 10

using namespace std;

//using DP
//pass the number and cache containing the set list 
void print_valid_paranthesis(int n, vector<unordered_set<string> > p_cache )
{
  if (n<=0)
    return;
  //pupulate the cache 
  for(int i=2; i<=n; ++i)
    {
      for (const string& it: p_cache[i-1])
	{
	  string s1 = "(" + (it) + ")" ;
	  string s2 = "()" + (it);
	  string s3 = (it) + "()";
	  p_cache[i].insert(s1);
	  p_cache[i].insert(s2);
	  p_cache[i].insert(s3);
	}
    }
  //print the result
  //for(const string& st: p_cache[n])
  for(auto& st: p_cache[n])
    {
      cout << st << endl;
    }
  cout << endl;
  
  return;
  
}



int main()
{
  //define set list
  vector<unordered_set<string> > pCache(MAX+1);
  //set the base cases:
  pCache[0].clear();
  pCache[1].clear();
  pCache[1].insert("()");
  int num;
  cin >> num;
  cout << "Number of paranthesis: " << num << endl; 
  //call the function to print all valid permutations
  print_valid_paranthesis(num, pCache);
  return 0;
}
