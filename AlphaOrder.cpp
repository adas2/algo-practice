/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

#include <iostream>
#include <string>

using namespace std;

//dictionary for alphabets
char alpha[27] = {' ', 'a','b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
                    'q', 'r', 's', 't', 'u', 'v','w', 'x', 'y', 'z'};

//recursive function to find the count of strings generate using the alphabet mapping
// use a predefine alphabet set a=1,...z=26
//Inputs: numeric string, k = len of string porcessed, candidate string , cnt of strings 
//return: count of strings formed
int alphaOrderCount(string str, int k, string cnd, int cnt )
{
    if (k==str.size())
    {
        //end of str reached, one complete string created
        cout << "String: " << cnd << endl;
        return 1;
    }
    
    int total=0;
    //check all possible alphabet candidates: 1 digit, 2 digits
    int cnd1, cnd2;
    cnd1 = str[k] - '0';
        
    //append to cnd string
    cnd += alpha[cnd1];
    total += alphaOrderCount(str, k+1, cnd, cnt);
    //revert back by discarding last char
    cnd=cnd.substr(0, cnd.size()-1);
    
    //check 2-digit alpha possible
    if(k<str.size()-1)
    {
        cnd2 = (str[k]-'0')*10+(str[k+1]-'0');
        if(cnd2<=26)
        {
            cnd += alpha[cnd2];
            total += alphaOrderCount(str, k+2, cnd, cnt);
        }
    }
    return total;
    
}

//can be done in DP with following formula:
//      DP[i] = DP[i-1]+DP[i-2] + 1; where DP[i]==> # of strings till the ith char in str


int main()
{
    string test="1234566";
    cout << "Total number of strings : " << alphaOrderCount(test, 0, "",0) << endl; 
}