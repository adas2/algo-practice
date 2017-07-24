/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

#include<iostream>
#include<vector>

using namespace std;

void print_vector(vector<int> col)
{
    cout << "[ " ;
    for (auto elem : col) {
        cout << elem << " ";
    }
    cout << "]" << endl;
}

//Note: no specific lexicographic order is achieved here
void print_subsets(vector<int> arr, int k, vector<int> &cnd)
{
    //termination case:
    if(k==arr.size())
    {
        print_vector(cnd);
        return;
    }
        
    //create candidate sets
    //take a number OR drop the number, 2 possibilities
    
    //DON'T TAKE
    print_subsets(arr, k+1, cnd);
    
    //TAKE:
    cnd.push_back(arr[k]);
    print_subsets(arr, k+1, cnd);
    
    //backtrack (IMPORTANT)
    cnd.pop_back();
    
         
    return;
    
}

int main()
{
    vector<int> test={1,2,3};
    vector<int> cnd;
    print_subsets(test, 0, cnd);
}

