#include <iostream>
#include <string>
#include <vector>
#include <sstream>

using namespace std;


class Sentence{
public:
  static string reverse(string &str)
  {
    //need to copy into a char arr?

    //traverse from both sides
    int i=0, j=str.size()-1;
    //terminate when i>j
    while (i<=j)
      {
	char temp=str[i];
	str[i] = str[j];
	str[j] = temp;
	i++;
	j--;
      }
    return str;
  }
  static string reverse_words(string &str)
  {
    //define delimeter
    stringstream ss(str);
    //ss.str(str);
    string ostr, temp;
    int pos = 0;

    while (getline(ss, temp, ' '))
      {
	//cout << "Temp: " << reverse(temp) << endl;
	ostr += reverse(temp);
	ostr += " ";
      }

    return ostr;
  }
  
};


int main()
{
  cout << "Enter a sentence: " << endl;
  string s, r_s, r_w;
  getline(cin, s);
  cout << "Entered: " << s << endl;
  r_s = Sentence::reverse(s);
  cout << "Reversed sentence: " << r_s << endl;
  cout << "Reverse words: " << Sentence::reverse_words(r_s) << endl;
}
