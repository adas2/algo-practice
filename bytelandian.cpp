#include <stdio.h>
#include <math.h>
#include <iostream>
#include <map>

class Solution
{
public:
  static unsigned int find_max(unsigned int val1, unsigned int val2);
  static unsigned int find_exch_value(unsigned int val);
  static void fill_value_map();
    

private:
  static std::map<unsigned int, unsigned int> value_map;
  static std::map<unsigned int, unsigned int>::iterator it;

};

//definiton of the map
std::map<unsigned int, unsigned int> Solution::value_map;
std::map<unsigned int, unsigned int>::iterator Solution::it;

unsigned int Solution::find_exch_value(unsigned int val)
{
  //unsigned int val_1 = floor(val/2);
  //unsigned int val_2 = floor(val/3);
  //unsigned int val_3 = floor(val/4);
  
  //std::cout << (val_1+val_2+val_3) << std::endl;
  it = Solution::value_map.find(val);
  if (it!=Solution::value_map.end())
    return Solution::value_map[val];
  else
    {
      std::cerr << "Error case";
      return 0;
    }

  //return (val_1+val_2+val_3);
}

unsigned int Solution::find_max(unsigned int val1, unsigned int val2)
{
  if (val1 < 0 || val2 < 0)
    {
      std::cerr << "Error Case";
      return 0;
    }
   
  if (val2 > val1)
    return val2;
  else
    return val1;
  
}


void Solution::fill_value_map()
{
  //Fill the lookup table before hand for 10000 entries
  
  

  Solution::value_map[0] = 0;
  Solution::value_map[1] = 1;

  for (unsigned int n=2; n < 10000 ; n++)
    {
      Solution::value_map[n]= Solution::find_max(n, (Solution::value_map[floor(n/2)]+Solution::value_map[floor(n/3)]+Solution::value_map[floor(n/4)]));
    }

  return;

}




int main()
{
  
  //fill the map
  Solution::fill_value_map();

  unsigned int coin_value;
  
  std::cout << "Enter the Value of bylandian ccoin: " << std::endl;

  std::cin >> coin_value;

 
  std::cout << "Max value: " << Solution::find_exch_value(coin_value) << std::endl;

}

