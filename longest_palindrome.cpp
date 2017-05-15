#include <iostream>
#include <string>

using namespace std;


bool isPalindrome(string s, int start, int end)
{
	cout << "isPali " << start << end << endl;
	cout << s.substr(start, (end-start)+1) << endl;
	while(start<=end)
	{
		if(s[start]!=s[end])
			return false;
		++start; --end;
		//else
	}
	return true;	
}

//Better way O(n^2)
string longestPalindrome2(string s)
{
	int n=s.size();
	if (n<=0)
		return "";
	if(n==1)
		return s;

	int max_len = 1, len;
	int i, j, k, start_i=0, end_j=0;
	//even palindrome case
	// start with i=0, j=1 expand in both ways
	for(i=0; i<n-1; i++) //O(n)
	{
		k=i;
		j = i+1;
		len=0;
		while(k>=0 && j<=n-1 && s[k]==s[j] ) //O(n)
		{
			len +=2;
			k--; j++;
		}
		if (len>max_len)
		{
			max_len = len;
			start_i=k+1;
			end_j = j-1;
		}

	}

	//odd palindrome case
	for(i = 1; i<n-1; ++i)  //O(n)
	{
		k=i-1;
		j=i+1;
		len = 1;
		while(k>=0 && j<=n-1 && s[k]==s[j] ) //O(n)
		{
			len +=2;
			k--;j++;
		}
		if (len>max_len)
		{
			max_len = len;
			start_i=k+1;
			end_j = j-1;
		}

	}

	cout << "Max plaindrome len: " << max_len << " " << start_i << " " << end_j << endl;
	return s.substr(start_i, max_len);
}




//Naive way T=O(n^3)
string longestPalindrome(string s) {
	if(s.size()<=0)
		return "";
	int max_len = 0, start_idx=-1;
	//int i=0, j=s.size()-1;
	for(int i=0;i<s.size();++i)
	{
		for(int j=s.size()-1;j>=i;--j){
			//if(s[i]!=s[j])
			//	j--;
			if(s[i]==s[j] && isPalindrome(s, i,j))
			{
				cout << i << j << endl;
				int len = (j-i)+1;
				if (max_len<len){
					max_len=len;
					start_idx = i;
				}
			}

		}

	}
	if(max_len>0)
		return s.substr(start_idx, max_len);
	else
		return ""; //no palindrome
}


int main()
{
	string test;
	cin >> test;	
	cout << "Original string: " << test << endl;
	cout << "Longest palindrome: " << longestPalindrome2(test);
	cout << endl;
	return 0;
}



