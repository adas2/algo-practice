/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

#include <iostream>
#include <climits>
#include <vector>
#include <string>
#include <cstdlib>
#include <algorithm>
#include <sstream>

#define CHUNK 18
//max value for a chunk
#define MAXV 999999999999999999 

using namespace std;

long convertStringtoLong(string s) {
    if (s.size() <= 0)
        return -1;
    //strto
    return strtol(s.c_str(), 0, 10);

}

//takes big num as a string and converts to array of long integers
//returns size of long array

vector<long> processBingNum(string num) {
    
    //reverse the num string
    reverse(num.begin(), num.end());
    //Read string by 19 digits each
    int cnt = 0;
    int idx=0; //start from left to right
    int len = (num.size() / CHUNK) ;
    len = (num.size()%CHUNK)==0?len:len+1;
    cout << "Len: " << len << endl;
    vector<long> arr(len,0);
    while (cnt < len) {
        //process string to long
        string temp=num.substr(idx, CHUNK);
        reverse(temp.begin(), temp.end());
        arr[cnt] = convertStringtoLong(temp);
        cout << arr[cnt] << "  ";
        cnt++;
        idx +=CHUNK;
       
    }
    cout << endl;
    cout << "Created array of len: " <<  len << endl;
    
    return arr;
}
//add two long vectors, return resultant long vector
vector<long> addBigNum(vector<long> a1, vector<long> a2)
{
    int len1=a1.size();
    int len2=a2.size();
    
    int min =(len1<len2)?len1:len2;
    int max =(len1<len2)?len2:len1;
    
    vector<long> result(max+1,0);
    
    for(int i=0; i<min; ++i)
    {
        result[i] += (a1[i]+a2[i]);
        //check if carry is there?
        
        //if((result[i]<a1[i])||(result[i]<a2[i]))
        if(result[i]>MAXV)
        {
            result[i+1] +=1; //add carry
            result[i] -= (MAXV+1); //normalize base
        }
    }
    //add remaining part
    for(int i=min; i<max; ++i)
    {
        if(len1>len2)
            result[i]+=a1[i];
        else
            result[i]+=a2[i];
    }
    for(int j=0; j<=max; ++j)
    {
        cout << result[j] << "  ";
    }
    cout << endl;
    return result;
}

//now convert long array to string 
string convertLongToString(vector<long> num)
{
    string outp="";
    stringstream ss;
    //start from end to first
    for (int i=num.size()-1; i>=0; --i)
    {
        ss << num[i];
        outp += ss.str();
    }
    return outp;
}

int main()
{
    //cout << LONG_MAX << endl;
  string test ="123456765432198761111111111111111111";   
  string test2="1212121212121212121212121212121";  
  string test3="999999999999999999999999999999999999";  
  
  vector<long> n1=processBingNum(test);
  vector<long> n2=processBingNum(test3);
  
  vector<long> final=addBigNum(n1, n2);
  
  cout << convertLongToString(final) << endl;
}