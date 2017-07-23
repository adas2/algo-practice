/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
#include<vector>
#include<iostream>
#include<unordered_map>

#define DECK_SIZE 8

using namespace std;

//check if shuffled card deck is a single riffle of original_deck? 
bool isShuffledDeckaRiffle(vector<int> shuffled_deck, vector<int> orig_deck)
{
    if(shuffled_deck.size()!=DECK_SIZE || orig_deck.size()!=DECK_SIZE)
        return false; //error case
    
    //cards with numbers 1...DECK_SIZE
    //populate the map with index of each card from irginal deck in the vector
    unordered_map<int, int> card_map;
    for(int i=0;i<DECK_SIZE; ++i)
    {
        //card value, index in deck 
        card_map.insert(make_pair(orig_deck[i], i));
    }
    
    //counters to keep tracks of the last index encountered
    int left_idx=0;
    int right_idx=DECK_SIZE/2;
    //traverse shuffled_deck to check the order
    for(int j=0; j<DECK_SIZE; ++j)
    {
        //for cards in the shuffled deck, indices wrt original deck should be increasing from 1 to 26, AND 27 to 52
        if(card_map.find(shuffled_deck[j])!=card_map.end())        
        {
            //index from orig_deck
            int index  = card_map[shuffled_deck[j]];
            //if index false in the left half
            if(index<(DECK_SIZE/2) && index <=left_idx)
            {
                left_idx++;
            }
            //index in right half 
            else if(index>=(DECK_SIZE/2) && index <=right_idx)
            {
                right_idx++;
            }
            else //out of order
                return false;
        }
        else //error since card valued couldn't be found in shuffled deck
            return false;
    }
    
}

int main()
{
    vector<int> test1={1,2,3,4,5,6,7,8};
    vector<int> test2={1,5,2,6,3,7,8,4};
    
    if(isShuffledDeckaRiffle(test2,test1))
        cout << "RIFFLE" << endl;
    else
        cout << "NO RIFFLE" << endl;
    
}