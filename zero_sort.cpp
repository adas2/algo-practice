/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
#include <iostream>
#include <vector>
#include <climits>

using namespace std;

int removeZeros(vector<int> &numbers)
{
 //corner case
  if(numbers.size()==0)
    return INT_MIN;
  
  //traverse the array from two sides
  int left=0;
  int right=numbers.size()-1;
  int cnt=0;
  
  //O(n) --> traverses all elements once
  while(left<right)
  {
       
    if(numbers[left]==0 && numbers[right]!=0 && left<right)
    {
    
      int temp = numbers[left];
      numbers[left]=numbers[right];
      numbers[right]=temp;
    }
    
    if(numbers[left]!=0)
      left++;
    
    if(numbers[right]==0)
      right--;
    
    //swap the left and right values
  }
  
  //count number of zeros
  //O(n) 
  for(int i=0; i<numbers.size(); ++i)
  {
    if(numbers[i]==0)
      cnt++;
  }
  /**/
  
  return cnt;
}

//test case {1,0,2,0,4,3,0}

int main()
{
    vector<int> test={1,0,2,0,4,3,0};
    //vector<int> test={0};
    removeZeros(test);
    for(auto x: test)
        cout << x << " ";
    cout << endl;
}

