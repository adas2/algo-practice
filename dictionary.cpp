#include <iostream>
#include <string>

#define CHAR_TO_INDEX(c) ((int)c-(int)'a')
#define MAX 26

using namespace std;

//gloabl dictionary structure


class trieNode{
public:
  //default
  trieNode() {
    for (int i=0; i<MAX; ++i)
      {
	this->alph[i]= NULL;
      }
    this->isWord = false;
    this->cnt = 0;
  }

  /***
  //custom constructor
  trieNode(char c, bool isword){
    
  }
  ***/
  
  trieNode *alph[MAX];
  bool isWord;
  int cnt;
};

class Trie{
public:
  //default constructor
  Trie(trieNode* node)
  {
    this->root=node;
  }
  trieNode *root;
  //procedures for trie
  void insertStr(string s);
  
  bool searchStr(string t);
  
  int startsWithStr(string prefix);
  void printTrie(string &prefix, trieNode *head);
};

//insert
void Trie::insertStr(string str)
{
  int index;
  //int i=0;
 
  if (!root || !str.length()){
    return;
  }
  //temp pointer
  trieNode *currNode = this->root; 
 
  for (int i=0; i<str.length(); ++i)
    {
      //cout << "insert: " <<  str[i] << endl;
      index = CHAR_TO_INDEX(str[i]);
      if(!currNode->alph[index])
	{
	  //create new Node for the alphabet
	  currNode->alph[index]= new trieNode();
	}
      //move to the child
      currNode = currNode->alph[index];
      //increment count for child 
      currNode->cnt++;
    }
  currNode->isWord = true;
}

void Trie::printTrie(string &prefix, trieNode *curr)
{
  //trieNode *curr = this->root;
  //while (!curr){
  if(curr->isWord){
    cout << prefix << endl;
    return;
  }
  int i;
  for(i=0; i< MAX; ++i){
    //check all non-NULL nodes and any
    if(curr->alph[i])
      {
	prefix.push_back(char('a'+i));
	//cout << i << ":" << curr->alph[i]->cnt << " --> ";
	//increment to child and recurse
	//curr = curr->alph[i];
	printTrie(prefix, curr->alph[i]);
	prefix.pop_back();
      }
  }
  //cout << endl;
  //}
}

//search word in dictionary
//return the TRUE if it is only a valid word, false otherwise
bool Trie::searchStr(string word)
{
  int i=0;
  int idx;
  trieNode *curr = this->root;
  while(i<word.length())
    {
      idx = CHAR_TO_INDEX(word[i]);
      if(!curr->alph[idx]) //character not found
	return false;
      //char found, move to next char and child in trie
      curr=curr->alph[idx];
      ++i;
    }
  //last node should indicate a word
  if(curr->isWord)
    return true;
  else // string found but not a word
    return false;
}

//find # of words that starts wiht the prefix 
//return # words prefix , 0 else
int Trie::startsWithStr(string prefix)
{
  int i=0;
  int idx;
  trieNode *curr = this->root;
  //traverse the prefix
  while(i<prefix.length())
    {
      idx = CHAR_TO_INDEX(prefix[i]);
      if(!curr->alph[idx]) //character not found
	return 0;
      //char found, move to next char and child in trie
      curr=curr->alph[idx];
      ++i;
    }
  //end of prefix
  //check if a word return zero
  if(!curr->isWord)
    return curr->cnt;
  //else
  return 0;
}


int main()
{
  //trieNode *head = new trieNode;
  Trie dict(new trieNode);
  
  string w1="test";
  string p;
  dict.insertStr(w1);
  dict.insertStr("ten");
  dict.insertStr("pen");
  dict.printTrie(p, dict.root);
  if (!dict.searchStr("test")) cout << "NOT FOUND" << endl; else cout << "FOUND" << endl;
  if (!dict.searchStr("te")) cout << "NOT FOUND" << endl; else cout << "FOUND" << endl;
  cout << "Number of words starts with prefix: " << dict.startsWithStr("te") << endl;
  
  
  return 0;
}








