/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

//Merge two binary trees:
/***
struct tnode{
    struct tnode *left;
    struct tnode *right;
    int val;
};
***/

#include<iostream>
#include<cstdlib>

using namespace std;

class TreeNode{
    //struct tnode *root;
public:
    TreeNode *left;
    TreeNode *right;
    int val;
    TreeNode(int v, TreeNode *l, TreeNode *r) : val(v), left(l), right(r) {}
    TreeNode(int v) : val(v), left(NULL), right(NULL) {}
    
};

//result in root1
void merge_tree(TreeNode* &root1, TreeNode* &root2)
{
    //root2 null, root1 remains the same
    if(!root2)
    {
        return;
    }
    //root1 is NULL, root2 non-NULL, root1 becomes root2
    if(!root1)
    {
        root1 = root2;
        return;
        //return root1;
    }
        
    //both non-NULL
    root1->val += root2->val;
  
    merge_tree(root1->left, root2->left);
    merge_tree(root1->right, root2->right);
    
    return;
}

void printInorder(TreeNode *root)
{
  if(!root)
    return;
  printInorder(root->left);
  cout << " " << root->val << endl;;
  printInorder(root->right);
}


int main()
{
    TreeNode *root1 = new TreeNode(1, 
				new TreeNode(2, new TreeNode(4), new TreeNode(5)), 
				new TreeNode(3));
    TreeNode *root2 = new TreeNode(1, 
				new TreeNode(1), 
				new TreeNode(1, new TreeNode(4) , new TreeNode(5)));
    
    //printInorder(root1);
    //printInorder(root2);
        
    //TreeNode *result=NULL;
    merge_tree(root1, root2);
    printInorder(root1);
    
}

