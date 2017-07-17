/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

#include <iostream>
#include <vector>
#include <float.h>

using namespace std;

//given len of subarray k
double  findMaxSubArrayAvg(vector<int> arr, int k)
{
    //corner cases
    if (arr.size()<k)
        return DBL_MIN;
    
    //maintain running avg and max avg variables
    double  max_avg=0;
    double  curr_avg=0;
    int sum=0;
    int end=0; //index of the last max subarray last element 
    //sum the first k elements
    for(int i=0; i<k ; ++i)
    {
        sum += arr[i];
    }
    curr_avg=max_avg=(double)sum/k;;
    
    //create the moving window with k elements
    for(int i=k; i<arr.size()-1; ++i)
    {
        //keep track of largest avg; 
        curr_avg = double((curr_avg*k)-arr[i-k]+arr[i])/k;
        if(curr_avg > max_avg)
        {
            max_avg = curr_avg;
            end = i;
        }
    }
    cout << "Max avg subarray from " << end-k+1 << " to " << end << endl;
    
    return max_avg;
}


int main()
{
    vector<int> test= {5,-2,1,5,3,1,1,1};
    vector<int> test2={1, 12, -5, -6, 50, 3};
    cout << "Max avg subarray is "<< findMaxSubArrayAvg(test2, 4) << endl;
}
