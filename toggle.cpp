#include <iostream>
#include <vector>
#include <math.h>

#define toggle(a) (a==0)?1:0


using namespace std;

class Solution {
public:
    int bulbSwitch(int n) {
      //define array of bulb states as 0/1
      vector<int> bulbs(n+1, 1);
  
      //find the number of perfect squares from 1, 2,.....n
      //instead o0f counting the floor of the square root of n gives the num
      return floor(sqrt(n));
    }
};

int main()
{
  cout << " Neter the no. of bulbs" << endl;
  int nb;
  cin >> nb;
  Solution s;
  cout << "The no. of bulbs still on: " << s.bulbSwitch(nb) << endl;
  return 0;
}
