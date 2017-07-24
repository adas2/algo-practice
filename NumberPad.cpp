/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
#include<iostream>
#include<string>
//#include<unordered_map>
#include<vector>

using namespace std;

//global map/vector containing the numeric code to letter mapping
//0 and 1 have no string attached -:)
vector<string> num_pad={"","","abc","def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};


//recursive call input: numeric string, position in code, candidate string, total number of strings
void print_numberpad_strings(string code, int k, string &str, int &cnt)
{
    //error case
    if(code.size()==0)
        return;
    
    //termination case, whole code covered
    if(k==code.size())
    {
        cnt++;
        //print the candidate
        cout << str << endl;
        return;
    }
    
    //start of string
    if (k==0)
        str="";
    //create all possible candidates
    //if '0' or '1' or '-' return or insert ' '
    int val = code[k]-'0' ; 
    if((val>1) && (val<=9))
    {
        //create all possible strings
        for(int i=0; i<(num_pad[val]).size(); ++i, cnt)
        {
            //add char to candidate str
            str += num_pad[val][i];
            print_numberpad_strings(code, k+1, str, cnt);
            //backtrack IMPORTANT!
            //strip the last character from string
            str = str.substr(0,str.size()-1);
        }
    }
    else // 0 or 1 or else break
    {
        return;
        //insert blank space
        //str += " ";
        //print_numberpad_strings(code, k+1, str);
    }
    
    
}

int main()
{
    string test="7";
    //start with empty string
    
    string s;
    int count =0;
    print_numberpad_strings(test, 0, s, count);
    cout << "total= " << count << endl;
}
