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
// #include <deque>
#include <stack>
#include <unordered_map>

using namespace std;

bool validate_bracket(string str)
{
    // check null case
    if (str.size() == 0)
        return false;

    // define a stack/vector/deque of opening and closing brackets
    //  define type as char val '(' or '{' or '['
    deque<char> br_stack;

    for (int i = 0; i < str.size(); ++i)
    {
        // Note: To make this cleaner code can have a map of openers and corresponding closers
        // openers
        if (str[i] == '(' || str[i] == '{' || str[i] == '[')
        {
            // add to stack
            br_stack.push_back(str[i]);
        }
        // closers
        if (str[i] == ')')
        {
            if (br_stack.back() == '(')
            {
                // match, pop
                br_stack.pop_back();
            }
            else // unmatched bracket
                return false;
        }
        // check the top of stack, if match pop
        if (str[i] == '}')
        {
            //
            if (br_stack.back() == '{')
            {
                // match, pop
                br_stack.pop_back();
            }
            else // unmatched bracket
                return false;
        }
        if (str[i] == ']')
        {
            //
            if (br_stack.back() == '[')
            {
                // match, pop
                br_stack.pop_back();
            }
            else // unmatched bracket
                return false;
        }
    }
    // if stack empty string traversed, brackets are all matched
    // check if unmatched bracket left in stack
    return br_stack.empty();
}

/***/
// alternate version which outputs string after parenthesis normalization
string normalize_string_brackets(string str)
{
    // check null or error case
    if (str.size() == 0)
        return "-1";

    // define a stack of opening and closing brackets
    //  define type as char  '(' or '{' or '['

    unordered_map<char, char> br_map = {make_pair('(', ')'), make_pair('{', '}'), make_pair('[', ']')};

    // stack entry has char and corresponding index
    stack<pair<char, int>> br_stack;

    int i = 0;
    while (i < str.size())
    {

        // check if symbol in bracket_map
        if (br_map.find(str[i]) != br_map.end())
        {
            // push to stack, char and index
            br_stack.push(make_pair(str[i], i));
        }
        // if closer is found
        else if (str[i] == ')' || str[i] == '}' || str[i] == ']')
        {
            // stack should not be empty --> tricky!
            if (!br_stack.empty() && (str[i] == br_map[br_stack.top().first]))
            {
                // matched: good case pop the bracket from stack
                br_stack.pop();
            }
            else
            {
                // empty stack or unmatched bracket
                // bad case: erase bad bracket;
                str.erase(i, 1);
            }
        }

        ++i;
    }

    // delete the unmatched brackets left in the stack
    while (!br_stack.empty())
    {
        int indx = br_stack.top().second;
        str.erase(indx, 1);
        br_stack.pop();
    }

    return str;
}

//"{ [ ] ( ) }" should return true
//"{ [ ( ] ) }" should return false
//"{ [ }" should return false
int main()
{
    string test = "{ [ }";

    if (validate_bracket(test))
        cout << "TRUE" << endl;
    else
        cout << "False" << endl;

    // Normalize the string ")ab(()" -> "ab()"

    string test2 = ")ab(()";
    string test3 = "{a(b)[}";
    string test4 = "((((()())";

    cout << "Final: " << normalize_string_brackets(test3) << endl;
}
