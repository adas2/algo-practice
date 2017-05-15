#include <iostream>
#include <vector>
#include <algorithm>
#include <unordered_map>

#define diff(a,b) (a>b)?(a-b):(b-a)

using namespace std;

bool mycompare(int a, int b)
{
  return a<b;
}

//Part (i) dist is at most k between two elements
// Time = O(nlogn)
//Space = O(1)
bool isDiffK(vector<int> arr, int k)
{
  //if one or less element return false
  if (arr.size()<=1)
    return false;

  // otherwise sort the elements and 
  //sort(arr.begin(), arr.end(), less);
  sort(arr.begin(), arr.end(), mycompare);

  //check if there are adjacent elements which differ by k or less, return true
  //if no adj elements found then return false
  for (int i=0; i<arr.size()-1; ++i)
    {
      if(arr[i+1]-arr[i]<=k)
	{
	  cout << arr[i] << " " << arr[i+1] << endl; 
	  return true;
	}
    }

  return false; 
}

//Part(ii) find identical elements not more that D dist apart
//Time = O(n) Space = O(n)
bool isDistEqualD(vector<int> arr, int d)
{
  //if one or less element return false
  if (arr.size()<=1)
    return false;

  //create a map storing each element and its last occurred index
  unordered_map<int,int> posMap;
  int dist=d+1;
  //traverse the array
  for(int i=0; i<arr.size(); ++i)
  {
	if(posMap.find(arr[i])!=posMap.end() )
	{
		dist=i-posMap[arr[i]];
		if(dist<=d)
		{
			cout << "match with dist at " << posMap[arr[i]] << " " << i << endl;
			return true;
		}
		//else update index
		posMap[arr[i]]=i;

	}
	else // elem seen first time, create entry
	{
		posMap.insert(make_pair(arr[i],i));

	}
  }

  //no dist match found
  return false;
 
}

//Part (iii): find elems that diffre my at most k and are d or less dist apart
//Time = ? Space = ?
bool isDiffDistKD( vector<int> arr, int k, int d)
{
	//if one or less element return false
	if (arr.size()<=1)
	  return false;

	//O(nlogn) approach:
  	//traverse the array using a moving window
	//store the window elements in a balanced binary search tree, so that sorted order can be maintained 
	//check if the nearest elements in btree differ by k or less

	//O(n) approach
	//create a hash map for bins in which keys are k diff aparts: e.g. nums 0-2 will fall in bin 0, 3-5 in bin 1 etc.
	//for negative integers -2 to -1 in bin (-1); -5to -3 bin (-2) etc.
	//the value to be stored in each bin is the actual number and its arr index  
	unordered_map<int,pair<int,int> > binMap;
	int key;
	//traverse the array and fill the hash map as you go
	for(int i=0; i<arr.size(); ++i)
	{
		//calc hash key
		if(arr[i]>0) //+ve
			key = arr[i]/k;
		else //-ve num
			key = (arr[i]/k)-1;
		cout << "key: " << key << endl;
		//check if value is present for key+1 or key-1
		if(binMap.find(key+1)!=binMap.end() )
		{
			//check if the value and dist is in range
			cout << arr[i] << " " << i << " " << binMap[key+1].first << " " << binMap[key+1].second << endl;
			//cout << "dist: " << i-binMap[key+1].second << endl;
			//int temp = diff(binMap[key+1].second, i);
			//cout << temp << " " << d << endl;
			if((diff(binMap[key+1].first, arr[i])<=k) && (diff(binMap[key+1].second,i)<=d))
				return true;
		}
		if(binMap.find(key-1)!=binMap.end() )
		{
			//check if the value and dist is in range
			cout << arr[i] << " " << i << " " << binMap[key-1].first << " " << binMap[key-1].second << endl;
			//cout << "dist: " << i-binMap[key-1].second << endl;
			if((diff(binMap[key-1].first, arr[i])<=k) && (diff(binMap[key-1].second, i)<=d))
				return true;
		}
		//if no match is adjacent bins check current bin, if entry found check diff
		if(binMap.find(key)!=binMap.end() )
		{
			cout << arr[i] << " " << i << " " << binMap[key].first << " " << binMap[key].second << endl;
			if (diff(binMap[key].second, i)<=d)
				return true;
			else //update the value
				binMap[key]=make_pair(arr[i],i);
				
		}
		else
		{
			//create new entry
			binMap.insert(make_pair(key, make_pair(arr[i], i)));

		}


	}
	//no match found
	return false;


}


int main()
{
  vector<int> test = {7, 2, -4, 18, 9, 5, 12, -1 };
  int k=3;
  int d=2;
  
  if (isDiffK(test, k))
    cout << "true"  <<  endl;
  else 
    cout << "false"  <<  endl;
  
  if (isDistEqualD(test, d))
    cout << "true"  <<  endl;
  else 
    cout << "false"  <<  endl;
  
  if(isDiffDistKD(test, k, d))
    cout << "true"  <<  endl;
  else 
    cout << "false"  <<  endl;


  return 0;
}



