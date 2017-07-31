#include <iostream>
#include <vector>
#include <cmath>

//max size of the 2D vector
#define MAX 3
#define max(a,b) a>b?a:b

using namespace std;

/**/
/* Definition for a binary tree node.*/
struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode(int x) : val(x), left(NULL), right(NULL) {}
  TreeNode(int x, TreeNode *l, TreeNode *r) : val(x), left(l), right(r) {}
};
/****/

int findTreeHeight(TreeNode *root);

//pass the level and result array
void zigzagTraverse(TreeNode *root, int level, vector<vector<int> > &tArr)
{
  if(root==NULL)
    return;
  
  if(level%2==1) //odd level
    {
      //tArr[level].push_front(root->val);
      tArr[level].insert(tArr[level].begin(), root->val);
    }
  else //even level
    {
      tArr[level].push_back(root->val);
    }
  //increase level
  level++;

  
  //recurse in child subtrees
  zigzagTraverse(root->left, level, tArr);
  zigzagTraverse(root->right, level, tArr);
  
  return;
}


//class Solution {
//public:
vector<vector<int> > zigzagLevelOrder(TreeNode* root) 
{
  //define a 2D square matrix to hold the result
  
  //initialize result array
  vector<vector<int> > result;

  //set the size as the height of the tree
  int h=findTreeHeight(root);
  result.resize(h);
  
  zigzagTraverse(root, 0, result);

  return result;
}

int findTreeHeight(TreeNode *root)
{
  if (!root)
    return 0;
  int l_h=0, r_h=0;
  
  if(root->left)
    l_h =findTreeHeight(root->left);
  if(root->right)
    r_h =findTreeHeight(root->right);
  
  return (max(l_h,r_h)+1);

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
  TreeNode *root = new TreeNode(3, new TreeNode(9, new TreeNode(8), new TreeNode(10)), new TreeNode(20, new TreeNode(15), new TreeNode(7)));

  printInorder(root);

  vector<vector<int> > result = zigzagLevelOrder(root); 

  cout << "height: " << findTreeHeight(root) << endl;

  //PRINT
  for (int i=0; i<findTreeHeight(root); ++i)
    {
      cout << "{ " ;
      for (int j=0; j<result[i].size(); ++j)
	cout << result[i][j] <<", ";
      cout << "}" << endl;
    } 
  return 1;
}
