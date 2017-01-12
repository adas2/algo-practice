#include <iostream>
#include <string>
#include <unordered_map>
#include <list>
#include <bitset>

using namespace std;

class StrUtils{
public:
  static string reverse(string str);
  static string reverse(string &str, int start_indx, int end_idx);
  static string reverseWords(string sentence);
  static int findSubstr(string target, string body);
  static string IntegerToString(int num);
  static int StringToInt(string str);
  static void allSubsets(string str);
  static void rotate(string &str, int k);
  static string floatToString(float r, int precision);
  static void allPermutations(string &str, int start, int end);
};


string StrUtils::reverse(string str)
{
	//null case or zero lenght case
	if (str.empty())
		return "";
	
	int i=0, j=str.length()-1;
	while (i<j)
	{
		char temp = str[i];
		str[i] = str[j];
		str[j] = temp;
		++i; --j;
	}
	return str;
}


string StrUtils::reverse(string &str, int start, int end)
{
	
	//null case or zero lenght case
	if (str.empty())
		return "";
	int i=start, j=end;
//	if(i>j)
	

	if(i < 0)
		i=0;
	if(j >= str.size() )
		j=str.size()-1;
	
	while (i<j)
	{
		char temp = str[i];
		str[i] = str[j];
		str[j] = temp;
		++i; --j;
	}
	return str;
		

}


string StrUtils::reverseWords(string sentence)
{
	if(sentence.empty())
		return "";
	string temp_str = reverse(sentence);
	string result_str = "";
	//cout << temp_str  << endl;
	int i=0; int prev=0;
	while(i < temp_str.size() )
	{
	 	//check for " " chars, note the index
	 	if(temp_str[i] == ' ' || temp_str[i] == '\0')
		{
			reverse(temp_str, prev, i-1);
			prev=i+1;
		}
		i++;
	}
	reverse(temp_str, prev, i-1);
	return temp_str;

}

//naive implementation time= O(m*(n-m+1))
int StrUtils::findSubstr(string target, string body)
{
  int m=target.length();
  int n=body.length();
  if((n<m) || !n)
    return -1;
  cout << "M: " << m << " N: " << n << endl; 

  int j, idx=-1;
  //change starting position of body from left to right
  for (int i=0; i<=n-m ; ++i)
    {
      //start matching the target from left to right
      for (j=0; j<m; ++j)
	{
	  //mismatch happens 
	 if(target[j]!=body[i+j])
	    {
	      break;
	    }
	 //i++;
	}
      //if length m already matched 
      if (j==m){
	idx = i;
	cout << "Found a match at: " << idx << endl;
      }
    }
  return idx;
}


/***
//Boyer moore bad character rule implementation
int StrUtils::findSubstr(string target, string body)
{
  //empty case
  if (body.empty() || target.empty())
    return -1;

  //preprocess target to hold position of index
  int i;
  int j=target.size()-1;
  unordered_map<char, list<int>> pos;
  while(j>=0){
    pos[target[j]].push_back(j);
    --j;
  }
  
  int idx=-1;
  i=j=target.size()-1;
  cout << "Size: " << body.size() << endl;

  while((j>=0) && (i<body.size()))
    {
      //cout << "Case: " << body[i] <<  " " << target[j] << endl;
      //start from end  there is a match,
    if(body[i] == target[j])
      {
	--i; --j; 
	cout << "Match: " << i <<  " "<< i << endl;
      }
    else //mismatch at body[i] and target[j]
      {
	//shift pattern to last occurent of bad charcatre in target
	//while(idx > j){
	cout << "mismatch: " << i << " " << j << endl; 
	  if(pos.find(target[j])!=pos.end())
	    {
	      idx = pos[body[i]].front();
	      cout << "Next index: " << idx << endl; 
	      //pos[target[j]].pop_front();
	    }
	  
	  //shift string to align with previously unmatched character
	  i += (j-idx);
	  
	  //start matching from end of string
	  j=target.size()-1;
	  cout << "New i: " << i << endl;
	  cout << "New j: " << j << endl;
      }
    }
  //Target is traveresed
  if(j==-1)
    {
      return i+1;
    }

  return -1;

}

*****/



string StrUtils::IntegerToString(int num)
{
  bool sign = false; 
  //int i=0;
  string str = "";
  if (num<0){
    num *= -1;
    sign = true;
  }
  while(num){
    int digit = num%10;
    //add the digit to string as char
    str += ('0'+digit);
    num /= 10;
  }
  
  if(sign)
    return  ("-" + reverse(str));
  else
    return reverse(str);
  
}

int StrUtils::StringToInt(string str)
{
  if (!str.size())
    return -1; //error case
  bool sign;
  int val=0;
  int i=0;
  //check for sign
  if (str[0] == '-')
    {
      sign = true;
      ++i;
    }
  //check for leading zeros
  while (str[i] == '0' )
    ++i;
  //calculate number
  while(str[i] != '\0')
    {
      int new_val = val*10 +(str[i] - '0');
      ++i;
      val = new_val;
    }
  if (sign)
    return (val*-1);
  else
    return val;
}

//print all possible subsets of astring
void StrUtils::allSubsets(string str)
{
  string result;
  int len = str.size();
  //ietrate for all posisble subset cases
  for (int i=0; i < (1 << len); ++i)
    {
      result = "";
      //ietrate through characters of the string
      for(int j=0; j<len; j++)
	{
	  //check whether bit position is 
	  if(i & (1 << j))
	    result += str[j]; 
	}
      cout << result << endl;
    }
}

void StrUtils::rotate(string &str, int k)
{
  //length of string
  int len = str.size();
  //revrese the entire string
  reverse(str, 0, len-1);
  reverse(str, 0, k-1);
  reverse(str, k, len-1);
  return;
}

string StrUtils::floatToString(float r, int precision)
{
  	 //convert the float to whole number and fraction part
	int num = int(r); //can use floor also
	float frac = r - int(r);
	int result = 1;
	//calculate the 10 power precision
	while (precision != 0) {
		result *= 10;
		--precision;
	}
	int num_2 = int(frac*result); //integer part after multiplying with precision
	if (num<0) num_2 *=-1; //this correct extra negative sign for -ve flotas
	if(!num_2)
		return IntegerToString(num);
	else
		return (IntegerToString(num)+"."+IntegerToString(num_2) );
}

void StrUtils::allPermutations(string &str, int start, int end)
{
  if(start==end)
    {
      //recursion end print string
      cout << str << endl;
      return;
    }
  
  //iterate over the string
  for (int i=start; i <= end ; ++i)
    {
      //fill the first position by swapping with ith element
      char temp = str[i];
      str[i] = str[start];
      str[start] = temp;
      //cout << "swapped: " <<  i << ": " << start << " " << str << endl;
      //permute for rest of the string
      allPermutations(str, start+1, end);

      //backtrack to the original string 
      temp = str[start];
      str[start] = str[i];
      str[i] = temp;
      //cout << "bactracked: " << i << ": " << start << " " << str << endl;
     
    }
  
}

int main()
{
	StrUtils u;
	string mStr = "abc";
	cout << StrUtils::reverse(mStr) << endl;
	cout << u.reverse("") << endl;
	//cout << StrUtils::reverse(NULL) << endl;
	cout << u.reverse(" abc") << endl;
	//cout << StrUtils::reverse("dog cat", 4, 6) << endl;
	//cout << StrUtils::reverse("", 0,0) << endl;
	//cout << StrUtils::reverse("abc", 2, -2) << endl;
	//cout << StrUtils::reverse("abcde", -10, 30) << endl;
	cout << StrUtils::reverseWords("the quick brown") << endl;
	//cout << u.reverse(" abc") << endl;
	cout << StrUtils::findSubstr("Pan", "Japan") << endl;
	//cout << StrUtils::findSubstr("abb", "aaaababab bba ab ababb") << endl;
	/***
	cout << StrUtils::IntegerToString(-123) << endl;
	cout << StrUtils::StringToInt("-0345") << endl;
	StrUtils::allSubsets("abc");
	string test = "abcdefghijklmnopqrst1234567890";
	StrUtils::rotate(test, 7);
	cout << test << endl;
	cout << StrUtils::floatToString(230.12345, 2) << endl;
	string test2 = "abc";
	StrUtils::allPermutations(test2, 0, 2);
	***/
}


