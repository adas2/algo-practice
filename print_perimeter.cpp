#include <iostream>

// print the perimeter of a binary tree
//	        root
//	         /\
//	        V  V
//       left  right
//
//

using namespace std;

//<template class T>
struct t_node
{
  int val;
  struct t_node *left;
  struct t_node *right;
};

typedef struct t_node *tree;

class Solution
{
public:
  void printTreePerimeter(tree s);
  bool addNode(int val, tree &s);
  void initTree();
  void printInorder(tree s);
  void printLeaves(tree s);
  void printPeripheryLeft(tree s);
  void printPeripheryRight(tree s);
  // private:
  // root of the binary tree
  tree root;
};

void Solution::initTree()
{
  root = NULL;
}

bool Solution::addNode(int val, tree &s)
{
  if (s == NULL)
  {
    struct t_node *t = new (struct t_node);
    t->val = val;
    t->left = t->right = NULL;
    s = t;
    return true;
  }
  if (val < s->val)
    return addNode(val, s->left);
  else
    return addNode(val, s->right);
}

void Solution::printInorder(tree s)
{
  if (!s)
    return;

  printInorder(s->left);
  cout << " " << s->val << endl;
  ;
  printInorder(s->right);
}

void Solution::printLeaves(tree s)
{
  // cout << "Processing node: " << s->val << endl;
  // leaf nodes of a tree
  if (!s)
    return;

  if (!s->left && !s->right)
  {
    cout << " " << s->val << endl;
    return;
  }
  // if(!s->left)
  printLeaves(s->left);
  // if(!s->right)
  printLeaves(s->right);
}

void Solution::printPeripheryLeft(tree s)
{
  if (!s)
    return;
  // don't print the leftmost node (as it will be covered by the leaf nodes)
  if (s->left)
  {
    cout << " " << s->val << endl;
    printPeripheryLeft(s->left);
  }
  else if (s->right)
  {
    cout << " " << s->val << endl;
    printPeripheryLeft(s->right);
  }
}

void Solution::printPeripheryRight(tree s)
{
  // don't print the rightmost node (as it will be covered by the leaf nodes)
  if (!s)
    return;
  if (s->right)
  {
    printPeripheryRight(s->right);
    cout << " " << s->val << endl;
  }
  else if (s->left)
  {
    printPeripheryRight(s->left);
    cout << " " << s->val << endl;
  }
}

void Solution::printTreePerimeter(tree s)
{
  if (!s)
    return;
  cout << " " << s->val << endl;
  printPeripheryLeft(s->left);
  printLeaves(s);
  printPeripheryRight(s->right); // skip the root node since its already covered
}

int main()
{
  Solution sol;
  sol.initTree();
  sol.addNode(6, sol.root);
  sol.addNode(3, sol.root);
  sol.addNode(9, sol.root);
  // sol.addNode(1, sol.root);
  sol.addNode(5, sol.root);
  sol.addNode(8, sol.root);
  sol.addNode(11, sol.root);
  sol.addNode(4, sol.root);
  sol.addNode(7, sol.root);
  sol.addNode(10, sol.root);
  // sol.addNode(12, sol.root);
  cout << "In Order Traversal:" << endl;
  sol.printInorder(sol.root);
  cout << "Leaf of the Tree:" << endl;
  // sol.printLeaves(sol.root);
  cout << "Printing perimeter nodes" << endl;
  // sol.printPeripheryLeft(sol.root);
  // sol.printPeripheryRight(sol.root->right);
  sol.printTreePerimeter(sol.root);
}
