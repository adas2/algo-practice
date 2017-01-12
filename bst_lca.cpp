#include <iostream>

//find the lowest common ancestor in a BST; no parent node pointre given

using namespace std;

//<template class T>
struct t_node{
  int val;
  struct t_node *left;
  struct t_node *right;	
};

typedef struct t_node *tree;

class Solution{
public:
  bool addNode(int val, tree &s);
  void initTree();
  void lowest_common_ancestor(tree root, int val1, int val2 );
  //void printInorder(tree s);
 tree root;
};

void Solution::initTree() {
  root = NULL;
}

bool Solution::addNode(int val, tree &s)
{
	if (s==NULL)
	{	
		struct t_node *t = new (struct t_node);
		t->val = val;
		t->left=t->right=NULL;
		s=t;
		return true;	
	}
	if(val < s->val)
		return addNode(val, s->left);
	else
		return addNode(val, s->right);
}


void Solution::lowest_common_ancestor(tree root, int val1, int val2 )
{
	//null tree case
	if (!root)
	  return;
	
	if((root->val < val1) && (root->val < val2))
	{
		lowest_common_ancestor(root->right, val1, val2);
	}

	else if((root->val>val1) && (root->val>val2))
	{
		lowest_common_ancestor(root->left, val1, val2);;
	}

	/****
	//val matches any of the given val1 or val2 then root is the lowest common node
	if(root->val==val1 || root->val==val2)
	{
		cout << "LCA: " << root->val << endl;
		return;
	}
	//val lies between val1 and val2
	if((root->val < val1 ) && (root->val > val2 ))
	{
		cout << "LCA: " << root->val << endl;
		return;
	}
	if((root->val > val1 ) && (root->val < val2 ))
	{
		cout << "LCA: " << root->val << endl;
		return;
	}
	****/
	/**Note all 3 cases above have the same action and there are no other cases possible hence removve condition and just print root->val*
	 */
	else{
		cout << "LCA: " << root->val << endl;
		return;
	}
}


int main()
{
	Solution sol;
	sol.initTree();
	sol.addNode(6, sol.root);
	sol.addNode(3, sol.root);
	sol.addNode(9, sol.root);
	//sol.addNode(1, sol.root);
	sol.addNode(5, sol.root);
	sol.addNode(8, sol.root);
	sol.addNode(11, sol.root);
	sol.addNode(4, sol.root);
	sol.addNode(7, sol.root);
	sol.addNode(10, sol.root);
	//sol.addNode(12, sol.root);
	//cout << "In Order Traversal:" << endl;
	//sol.printInorder(sol.root);
	//cout << "Leaf of the Tree:" << endl;
	//sol.printLeaves(sol.root);
	cout << "Printing LCA of  nodes" << endl;
	//sol.printPeripheryLeft(sol.root);
	//sol.printPeripheryRight(sol.root->right);
	//sol.printTreePerimeter(sol.root);
	sol.lowest_common_ancestor(sol.root, 10, 8);
	
}



