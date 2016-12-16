#include <iostream>
#include <stdlib.h>
#include <stdint.h>

using namespace std;

class Solution {
public:
  void fillTable()
  {
  
    for (unsigned int i=1; i<=256; i++)
      {
	Solution::hTable[i] = (i&1) + Solution::hTable[i/2];
	//cout << "hT: " << hTable[i] << endl;
      }
  }
  
  int hammingWeight(uint32_t n) {
    fillTable();
    int s = sizeof(n);
    int count=0;
    cout << "Size: " << s << endl;
    unsigned char* p= (unsigned char*)&n; 
    for (int j=0; j<s; j++)
      {
	cout << Solution::hTable[*(p+j)] << endl;
	count+= Solution::hTable[*(p+j)];
      }

    return count;
      
  }
 
private:
  static unsigned int hTable[256];
};

unsigned int Solution::hTable[256];


int main()
{
  uint32_t num;
  Solution sol;
  cout << "Enter the number of caluclate hammign wieght: " << endl;
  cin >> num;
  cout << "Hamming Weight: " << sol.hammingWeight(num) << endl;
}
