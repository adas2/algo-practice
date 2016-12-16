#include <iostream>
#include <string>
#include <unordered_map>

using namespace std;

#define ALPHA_SIZE 26

//global hash table
static unordered_map<string, string> hTable;
string alphabet("ABCDEFGHIJKLMNOPQRSTUVWXYZ");

void createHashMap()
{
  //fill in a hash table with alphabet codes
  // 0->A, 1->B,.....,25-Z
  hTable.clear();
  for (int i=0; i<26; i++)
    {
      string idx = to_string(i);
      //hTable.emplace(idx, alphabet[i]);
      hTable[idx]=alphabet[i];
    }
}

void findAllPermutations(string str, string res)
{
  //base case
  if (str.size() < 1)
    {
      cout << res << endl;
      //res.clear();
      return;
    }

  
  //case : 1 digit letter found in table
  if (hTable.find(str.substr(0,1)) != hTable.end())
    {
      //cout << "Substr: " << str.substr(0,1) << endl;
      res += hTable[str.substr(0,1)];
      //cout << "res: " << res << endl;
      //string nS(res);
      findAllPermutations(str.substr(1,str.size()-1), res);
      
    }
  
  // backtrack the result buffer by deleting the last letter 
  // such that case 2 can run on unmodified result buffer
  res.pop_back();
  

  //case : two 2 digit letters are found
  if(str.size()>=2 && hTable.find(str.substr(0,2))!=hTable.end())
    {
      //cout << "Substr: " << str.substr(0,2) << endl;
      res += hTable[str.substr(0,2)];
      //cout << "res: " << res << endl;
      //string nS(res);
      findAllPermutations(str.substr(2, str.size()-2), res);
    }


  //return;
}


int main()
{
  
  unordered_map<string,string>::iterator it;
  createHashMap();
  //for (auto& x: hTable)
  /**/
  for (it=hTable.begin(); it!=hTable.end(); ++it)
    cout << (*it).first << ": " << (*it).second << endl;
  /**/

  cout << "Enter the string" << endl;
  string inp;
  cin >> inp;
  string result("");

 
  
  findAllPermutations(inp, result);

}
