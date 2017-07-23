/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

// Given a string with brackets validate whether they are properly nested
// Alternate version: normalize a string with uneven parenthesis

#include <string>
#include <vector>
#include <iostream>
#include <deque>

using namespace std;

bool validate_bracket(string str)
{
    //check null case
    if(str.size()==0)
       return false;

//define a stack/vector/deque of opening and closing brackets
// define type as char val '(' or '{' or '['
    deque<char> br_stack;
    
    for(int i=0; i<str.size(); ++i)
    {
        //Note: To make this cleaner code can have a map of openers and corresponding closers
        //openers
        if(str[i]=='(' || str[i]=='{' || str[i]=='[')
        {
            //add to stack
            br_stack.push_back(str[i]);
        }
        //closers
        if(str[i]==')')
        {
            if(br_stack.back()=='(')
            {
              //match, pop
                br_stack.pop_back();
            }
            else //unmatched bracket
                return false;
        }
        //check the top of stack, if match pop
        if(str[i]=='}')
        {
            //
            if(br_stack.back()=='{')
            {
              //match, pop
                br_stack.pop_back();
            }
            else //unmatched bracket
                return false;
        }
        if(str[i]==']')
        {
            //
            if(br_stack.back()=='[')
            {
              //match, pop
                br_stack.pop_back();
            }
            else //unmatched bracket
                return false;
        }
    }
    //if stack empty string traversed, brackets are all matched
        //check if unmatched bracket left in stack
    return br_stack.empty();
        
    
    
}



//"{ [ ] ( ) }" should return true
//"{ [ ( ] ) }" should return false
//"{ [ }" should return false
int main()
{
    string test="{ [ }";
    if(validate_bracket(test))
        cout << "TRUE" << endl;
    else
        cout << "False" << endl;
}