#include <iostream>
#include <unordered_map>

#define max(a,b) (a>b)?a:b

using namespace std;

//original infinite wrap-around string = "......abcdefg....xyzabcd....xyzabc..."

int findWrapAroundSubStr(string p)
{
  //base case emoty string zero matches
  if (!p.size())
    return 0;
  
  cout << "Input: " << p << endl;

  //initialize hashMap with ASCII alphabets and longest substr ending at char
  unordered_map<char,int> hMap;
  hMap.clear();
  //first entry
  hMap.insert(make_pair(p[0],1));
  
  int sub_size=1; 
  //traverse the string to find the 
  for(int i=1; i<p.size(); ++i)
    {
      //diff in ASCII value between contiguois chars is ==1 
      if ((int(p[i])-int(p[i-1])==1)||(p[i]=='a' && p[i-1]=='z'))    
	sub_size++;
      else //reset subsize
	sub_size=1;
      
      //update entry in map
      if(hMap.find(p[i])!=hMap.end())
	{
	  //with max substr len found
	  hMap[p[i]]=max(hMap[p[i]], sub_size);
	}
      else //entry not found, insert
	{
	  //insert substr len for the char
	  hMap.insert(make_pair(p[i], sub_size));
	}
    }

  //traverse the map to find the total unique substr count
  int result=0;
  for(unordered_map<char,int>::const_iterator it=hMap.begin(); it!=hMap.end(); ++it)
    {
      result += (*it).second;
    }
  
  return result;

}


int main()
{
  string pattern;
  cout << "Enter a string: " << endl;
  cin >> pattern;
  cout << "result = ";
  cout << findWrapAroundSubStr(pattern);
  cout << endl;
}
