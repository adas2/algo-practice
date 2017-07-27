/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
/**

word list:   "peak", "perk", "park", "pa"
input:   

"peak" -> true
"p**k" -> true
"part" -> false
"pa" -> true

**/ 

#define ALPHA_SIZE 26

//Node 
  struct TrieNode{
      char c;
      struct TrideNode *alph[ALPHA_SIZE];
      bool isWord; //whether it is a workd 
  };

typedef struct TrieNode *Trie;

// Trie dict;
// dict->c='\0';
// for (int i=0; i<ALPHA_size; ++i)
//   dictionary->alph[] = NULL;

TrieNode *TrieNode(char ch)
{
  TrieNode *newNode = new (struct TrieNode);
  newNode->c=ch;
  for (int i=0; i<ALPHA_SIZE; ++i)
    newNode->alph[i]=NULL;
  newNode->isWord=false;
}


void insert(Trie &dict, string wrd)
{
  int i=0;
  TrieNode *curr=dict;
  while(i<wrd.size())
  {
    if(!curr->alph[i])
      curr->alph[wrd[i]] = new TrieNode(wrd[i]);
    curr=curr->alpha[wrd[i]];
    i++;
  }
  if(i==wrd.size())
    curr->isWord=true;
  
}

//void populat_dict(Trie &dict, vector<string> word_list);

//k is the index of the wrd 
bool searchWord(struct TrieNode *dict, string wrd, int &k)
{
  if(wrd.size()==0 || dict==NULL)
    return false;
  
  if(k==wrd.size() && dict->isWord)
    return true;
  
  struct TrieNode *curr=dict;
  
  //searching * char
  if(k<wrd.size() && wrd[k]=='*')
  {
    for(int i=0; i<ALPHA_SIZE; ++i)
      if(curr->alph[i])
        if(searchWord(curr->alph[i], k+1))
          return true;
    return false;
  }
  
  //searching regular char
  else if (k<wrd.size() && curr->alph[int(wrd[k])-int('a')])
  {
    
    return searchWord(curr->alph[int(wrd[k])-int('a')], wrd, k+1);
  }
  else //no node 
  {
    
   return false; 
  }
    
  
}
//test:
//wrd = "p**k"
//searchWord(dict,wrd,0);




