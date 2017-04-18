/***Given a pattern and a string input - find if the string follows the same pattern and return true or false.
 ***/

#include <iostream>
#include <unordered_map>
#include <string>

using namespace std;

bool isPatternRule(string str, string p, int s, int k, unordered_map<char, string> hMap)
{
  if(k==p.size()&&s==str.size())
    {
      //print solution
      return true;
    }
  //if pattern or string is exhausted
  //if (k==p.size() || s==str.size())
  //return false;

  //recursive backtracking
  //k++;
  //create candidate strings and check for each
  for(int i=1; i<=str.size()-s; ++i )
    {
      string cnd = str.substr(s, i);
      cout << "Candidate string: " << cnd << endl;
      if(hMap.find(p[k])!=hMap.end())
	{
	  //candidate string matches with hashmap rule
	  if((hMap[p[k]]==cnd)&& isPatternRule(str,p,s+i,k+1,hMap)) 
	    return true;
	  //return false;
	}
      else // insert entry
	{
	  cout << "Entry inserted for " << p[k-1] <<": " << cnd << endl;
	  hMap.insert(make_pair(p[k],cnd));
	  if(isPatternRule(str, p, s+i, k+1, hMap))
	    return true;
	  else //backtrack
	    hMap.erase(p[k]);
      	}
      //recurse
      
    }
  //if no candidates follow rule
  return false;

}


int main()
{
  string p = "abba";
  string mystr = "boygirlgirlboy";
  
  unordered_map<char,string> hMap;

  if (isPatternRule(mystr,p,0,0,hMap))
    cout << "True: Rule is followed" << endl;
  else
    cout << "False: Rule NOT followed" << endl;
}

