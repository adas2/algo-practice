/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

#include <iostream>
#include <vector>

using namespace std;


//Find intersection between two sorted arrays
//Time=O(m+n) Space=O(1)
void print_instersection(vector<int> arr1, vector<int> arr2)
{
    //error case
    if(arr1.size()==0 || arr2.size()==0)
        return;
    
    //traverse through both arrays in sorted order and see if elements are interesecting
    //No extra space required
    
    int i=0, j=0;
    while(i<arr1.size() && j<arr2.size())
    {
        if(arr1[i]==arr2[j])
        {
            cout << arr1[i] << endl;
            i++; j++;
        }
        else if(arr1[i]<arr2[j])
        {
            i++;
        }
        else
        {
            j++;
        }
    }
    //if(i==arr1.size)
    return;
}

int main()
{
    vector<int> test1= {1,3,5,7,9};
    vector<int> test2= {-3,2,4,5,6,7,10,25};
    
    print_instersection(test1, test2);
    
}