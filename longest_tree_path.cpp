#include <iostream>
#include <vector>
#include <string>

#define max(a,b)  (a>b)?a:b

using namespace std;

class TreeNode{
public:
  TreeNode(string s): val(s), left(NULL), right(NULL), l_wt(0), r_wt(0)
  {
  }
  TreeNode(string s, TreeNode *ltree, TreeNode *rtree, int lw, int rw)
  {
    val=s;
    left=ltree;
    right=rtree;
    l_wt=lw;
    r_wt=rw;
  }
  int getLwt()
  {
    return l_wt;
  }
  int getRwt()
  {
    return r_wt;
  }
  int findLongestPath(TreeNode *root, int &max_height);
  
  
private:
  string val;
  int l_wt; //left path weight
  int r_wt; //right path weight
  TreeNode *left; //left child
  TreeNode *right; //right child 
};

//Find the longest path through any two nodes in the tree 
//NOTE: doesnt have to be through root
int TreeNode::findLongestPath(TreeNode *root, int &max_height)
{
  //NULL case no path
  if(!root)
  {
    max_height=0; //NULL node
    return 0;
  }
  //case 1: longest path is through root i.e. height of the left subtree + height of rsubtree 
  // + corresponding path weights
  //case 2: longest path is through left subtree or right subtree
  int lht, rht;
  int lpath=findLongestPath(root->left, lht);
  int rpath=findLongestPath(root->right, rht);
  max_height = max((lht+root->l_wt),(rht+root->r_wt)); 
  
  int max_path = max(lpath,rpath);
  int curr_path = (lht+root->l_wt)+(rht+root->r_wt);
  max_path=max(max_path, curr_path);
  cout << root->val << " " << max_height << " " << max_path << endl;

  return max_path;
}


int main()
{
  TreeNode *root = new TreeNode("a", new TreeNode("b", new TreeNode("d", new TreeNode("e"), new TreeNode("f"), 100,100), NULL, 10,0), new TreeNode("c"), 20, 10);  
  int h;
  cout << "Longest Path: " << root->findLongestPath(root, h);
  cout << endl;
  cout << "height of the tree = " << h << endl; 
  return 1;
}

