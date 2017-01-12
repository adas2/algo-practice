#include <iostream>
#include <unordered_map>
#include <string>

using namespace std;

void find_duplicates(string str)
{
  //ASCII char to Count MAP
  unordered_map<char, int> char_map;
  //read the string characters and create hash_map time=O(n)
  for(int i=0; i<str.size(); ++i)
    {
      //how do you figure upper case vs lower case?
      if((str[i]>='A') && (str[i]<='Z'))
	{
	  //uppercase letter already present
	  if(char_map.find(str[i]+32) != char_map.end())
	    char_map[str[i]+32]++; //create entry for lower case version only
	  else //seen first time
	    char_map[str[i]+32] = 1;
	}
      else //all other characters
	{
	  //already present
	  if(char_map.find(str[i]) != char_map.end())
	    {
	      char_map[str[i]]++;
	    }
	  else //character seen for the first time
	    {
	      char_map[str[i]] = 1;
	    }
	}
    }

  //print out duplicates based on count time = O(n)
  for (auto it=char_map.begin(); it!=char_map.end(); ++it)
    {
      if ((*it).second > 1)
	{
	  cout << "Duplicate found: " << (*it).first << " With Count: " << (*it).second << endl ;
	}
    }
}

int main()
{
  //string s;
  cout << int('A') << " : "<< int('Z') <<  endl;
  cout << "String : " << endl;
  //cin >> s;
  for (std::string line; std::getline(std::cin, line);) {
    std::cout << line << std::endl;
    find_duplicates(line);
  }
  //find_duplicates(s);
  
  return 1;
}
