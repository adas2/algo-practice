#include <iostream>
#include <vector>
#include <unordered_map>
//#include <climits>

using namespace std;

int countSubarrayMultiple(vector<int> arr, int k)
{
  //error case
  if (arr.size()==0 || k<0)
    return 0;

  //total count of occurences
  int total_cnt = 0;
    
  //create array with modulus(k) of running array sum
  vector<int> mod_sum;
  //intialize
  mod_sum.resize(arr.size());
  int sum = arr[0];
  mod_sum[0]=sum%k;
  for(int i=1; i< arr.size(); ++i)
    {
      sum += arr[i];
      //store modulus k of the cumultaive sum
      mod_sum[i] = sum%k;
    }

  //print mod array
  for (auto& x: mod_sum)
    cout << x << " ";
  cout << endl;
  
  //define hash maps to store the mod(kof ) subarray and corresponfing freq
  unordered_map<int,int> modMap;
  //insert null count i.e. when no elements added modulus is 0 and count 1 
  modMap.insert(make_pair(0,1));

  //iterate through the cumulative sum array lookup hashtable
  for(int i=0; i<mod_sum.size(); ++i)
    {
      if(modMap.find(mod_sum[i])==modMap.end())
	{//insert
	  //first occurence hence frequency counter=1
	  modMap.insert(make_pair(mod_sum[i],1)); 
	}
      else{
	//else increment total_count and increase the frequency
	total_cnt += modMap[mod_sum[i]];
	modMap[mod_sum[i]]++;
	//modMap.find(mod_sum[i]).second++;
      }
    }
  
  return total_cnt;

}

int main()
{
  vector<int> test={4,7,-2,1,18};
  int k = 5;

  int result = countSubarrayMultiple(test, k);

  cout << "Result: " << result << endl;
}
