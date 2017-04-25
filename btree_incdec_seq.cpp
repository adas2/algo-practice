#include <iostream>
#include <climits>

#define max(a,b) a>b?a:b

using namespace std;

/* Definition for a binary tree node.*/
struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode(int x) : val(x), left(NULL), right(NULL) {}
  TreeNode(int x, TreeNode *l, TreeNode *r) : val(x), left(l), right(r) {}
};
/****/

void printInorder(TreeNode *root)
{
  if(!root)
    return;
  printInorder(root->left);
  cout << " " << root->val << endl;;
  printInorder(root->right);
}

//return type: LIC, LDS
pair<int,int> findMaxIncDecPath(TreeNode *root, int &max_len )
{
  //base case
  if (!root)
    {
      return make_pair(0,0);
    }
  
  pair<int,int> val_left = findMaxIncDecPath(root->left, max_len);
  pair<int,int> val_right = findMaxIncDecPath(root->right, max_len);
  
  //LIS and LDs of left and right subtree
  int LIS_left = val_left.first;
  int LDS_left = val_left.second;
  int LIS_right = val_right.first;
  int LDS_right = val_right.second;
  
  cout << "Root:  " << root->val << " " << LIS_left << LIS_right << LDS_left << LDS_right << endl;

  int LIS_root=0, LDS_root=0;
  if(root->left!=NULL && root->val < root->left->val )
    LIS_root=LIS_left+1;
  if(root->left!=NULL && root->val < root->right->val )
    LIS_root=max(LIS_root,(LIS_right+1));
  
  //for LDS
  if(root->left!=NULL && root->val > root->left->val )
    LDS_root=LDS_left+1;
  if(root->left!=NULL && root->val > root->right->val )
    LDS_root=max(LDS_root,(LDS_right+1));

  //if leaf node
  if(!root->left && !root->right)
    {LIS_root=1; LDS_root=1;}

  //check if seq_len exceeds max
  int curr_len=LIS_root+LDS_root-1; //root counted twice
  
  if (curr_len>max_len)
    {
      max_len=curr_len;
    }

  return make_pair(LIS_root,LDS_root);
}




int main()
{
  TreeNode *root = new TreeNode(4, 
				new TreeNode(12, new TreeNode(1), new TreeNode(2)), 
				new TreeNode(7,new TreeNode(6, new TreeNode(10), new TreeNode(5)) , new TreeNode(8)));

  printInorder(root);

  int max_len = 0;

  pair<int,int> result=findMaxIncDecPath(root, max_len );

  cout << "Max len: " << max_len << endl;
  cout << "Root LIS: " << result.first << " Root LDS: " << result.second << endl;

  return 1;
}

