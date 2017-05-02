#include <iostream>
#include <string>
//#include <climits>
#include <vector>

using namespace std;

/****
//utility function that returns the count of the last substring match with given len form the map
// if no match found returns 1
int searchLastSubstr(string s, int sub_len, vector<vector<int> > sMap)
{
  int s_len = s.size();
  for(int i=; i<(s_len-sub_len); ++i)
    {
      if()
    }
}

*****/



string findRepeatingSubstr(string s)
{
  int len=s.size();  

  if(!len){
	  cout << "Empty String" << endl;
	  return s;
  }

  //int max = INT_MIN;
  int start_idx=0;
  int sub_len=0;
  
  //Using DP appraoch:
  //create a 2D map M storing the following at M[i][j]:
  // count of occurences of sub-string sub ending at j with len i in the original string s
  //e.g. is s = "abcd"; M[2][3]= numbe of occurences of s[1..3] i.e. "cd" seen so far in the string
  // i.e. M[2][3]=1; Similarly, in string s ="cdbcd" M[2][4]=2; 2 occurences of cd 

  vector<vector<int> > M;
  M.resize(len+1); //takes into account 0th size string
  //initialize with all 0's
  for(int i=0; i<=len; ++i)
    M[i].resize(len+1, 0);

  //fill the matrix using DP
  //for substrings ranging in size form 1 to len-1
  for(int i=1; i<=len; ++i)
    {
      //from lenght of oiginal string from 0th index to last i.e. len-1
      for(int j=0; j<len; ++j)
	{
	  //Rules:
	  // if M[i][j] =  (M[i][k] + 1)  ==> if s[j-i..j] is seen before ending at index k in string
	  // 1;             ==> otherwise
 
	  //cout << "Checking substring that ends at " << j << " of len " << i << endl; //": " << s.substr(j-i,i) << endl; 
	  if(j>=i-1)
	  	M[i][j]=1;
	  /****/
	  //search backwards: this can be elliminated if using a hashMap with O(n2) spaces
	  //strings can overlap hence decrement by 1; if they don't overlap decrement k by i
	  //for(int k=j-i; k>=0; k-=i )
	  for(int k=j-i; k>=0; --k )
	    {
	      //cout << "\tString to search: starts at " << k << " of len " << i << endl;
	      if(s.substr(k,i)==s.substr(j-i+1,i)){
		cout <<"Matched string: " << s.substr(j-i+1,i) << endl;
		M[i][j]=M[i][k+i]+1;
		break; //break out this last for loop and move on the next size substr
	      }
	      //else
	    }
	  /******/	  
	}
    }

  //traverse the map in reverse fashion to check the longest substr thats repeated 
  for(int i=len; i>0; --i)
    {
      for(int j=len-1; j>=0; --j)
	{
	  if (M[i][j] > 1)
	    {
	      //max = M[i][j];
	      start_idx=j-i+1;
	      sub_len=i;
  	      return s.substr(start_idx, sub_len);
	    }
	}
    }
 // cout << "Max repeated substr: " << max << endl;
  cout << "No sub string repeated "  << endl;
  return s;

}


int main()
{
  //string test="pqabcdpqbabcpq";
  //string test="aabbaddabad";
  string test="abababa";
  cout << "Original string: " << test << endl;
  cout << findRepeatingSubstr(test) << endl;
  return 1;
}
