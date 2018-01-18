#include <iostream>
#include <vector>
#include <unordered_map>
//#include <climits>

//Incorrect solution: check

using namespace std;

int countSubarrayMultiple(vector<int> arr, int k)
{
  //error case
  if (arr.size()==0 || k<0)
    return 0;

  //print array 
  cout << "Input array" << endl;
  for (auto& x: arr)
    cout << x << " ";
  cout << endl;

  //total count of occurrences
  int total_cnt = 0;
    
  //create array with modulus(k) of running array sum
  vector<int> mod_sum(arr.size());

  int sum = 0;
  //mod_sum[0]=sum%k;
  for(int i=0; i< arr.size(); ++i)
    {
      sum += arr[i];
      //store modulus k of the cumulative sum
      mod_sum[i] = sum%k;
    }

  //print modulus array
  cout << "Modulus %k array" << endl;
  for (auto& x: mod_sum)
    cout << x << " ";
  cout << endl;
  
  //define hash maps to store the mod(k) of subarray and corresponding freq
  unordered_map<int,int> modMap;

  //insert null count i.e. when no elements added modulus is 0 and count 1 
  modMap.insert(make_pair(0,1));

  //iterate through the modulus sum array lookup hashtable
  for(int i=0; i<mod_sum.size(); ++i)
    {
      //check if 0 or (k-mod_sum[i]) is already in the hahsMap
      int key = mod_sum[i];
      
      if(modMap.find(key)==modMap.end())
        {
         //insert first occurrence hence frequency counter=1
         modMap.insert(make_pair(key,1)); 
        }
        else 
        {
         //else increment total_count and increase the frequency
         total_cnt += modMap[key];
         //total_cnt += 1 ;
         modMap[key]++;
        }
      //cout << "Count: " << total_cnt << endl;
  }  
  return total_cnt;
}


int main()
{
  //vector<int> test={4,7,-2,1,18};
  vector<int> test={15, -5, 3,5,3,9,1};
  int k = 10;
  
  int result = countSubarrayMultiple(test, k);

  cout << "Num subarray sum that are multiple of : " << k << endl;
  cout << result << endl;
}





