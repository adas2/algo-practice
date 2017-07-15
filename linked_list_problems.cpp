#include <iostream>
#include <string>
#include <cstdlib>

using namespace std;

class Node{
public:
  //Constructor
  Node(string str, Node *next){
    this->value = str;
    this->next = next;
    //prev = prev;
  }

  string val(){
    return this->value;
  }

  Node *nex(){
    return this->next;
  }

  void setVal(string val){
    this->value= val;
  }
  void setNext(Node *nex){
    this->next = nex;
  }
  
private:
  string value;
  Node *next;
  //Node *prev;

};

class List{
public:
  //constructor
  List(Node *head){
    this->head = head;
  }
  List (Node *head, Node* tail){
    this->head = head;
    this->tail = tail;
  }
  void printList();
  void reverseList();
  void insertBeforeHead(string value);
  void insertAfterTail(string value);
  string valueAt(int index); 
  void deleteNodeAt(int index);
  bool isCyclePresent();
  //for test purposes
  void create_cycle(){
    this->tail->setNext(this->head);
  }
  void createCycle(Node *rand)
  {
	this->tail->setNext(rand);	
  }
  string getMiddleValue();
  string getNthValueFromEnd(int n);
  Node* getNthPtr(int n);
  void reverseListRecursive();
  void detectAndRemoveCycle();

private:
  Node *head;
  Node *tail;
};

void List::printList()
{
  Node *curr = this->head;
  while (curr)
    {
      cout << curr->val() << "-->" ;
      curr = curr->nex();
    }
  cout << "NULL" << endl;
  //if(this->tail)
  //cout << "tail: " << this->tail->val() << endl;
}

void List::insertBeforeHead(string value)
{
  Node *nd = new Node (value, this->head);
  this->head = nd; 
  //update tail if last node
  if (!nd->nex())
    this->tail = nd; 
}

void List::insertAfterTail(string value)
{
  //create last node
  Node *nd = new Node(value, NULL);
  //if tail is not NULL attach the node to original tail
  if(!this->tail)
  	this->tail->setNext(nd);
  //update tail
  this->tail = nd;
  //if only element update head 
  if (!this->head)
    this->head = nd;
}

string List::valueAt(int index){
  int i=0;
  Node *curr = this->head;
  while ((i<index) && curr){
    curr = curr->nex();
    ++i;
  }
  if (curr)
    return curr->val();
  else //error case
    return "ERROR";
}

void List::deleteNodeAt(int index)
{
  int cnt=0;
  Node *curr=head;
  //case when index 0: first element gets deleted
  if(index==0){
    //curr = head;
    this->head = curr->nex();
    delete curr;
    if(!head)
      tail=head; //NULL list
    
    return;
  }

  //Else: case when index non-zero
  while((cnt < (index-1)) && curr){
    curr= curr->nex();
    cnt++;
  }
  //node previous to node to be deleted
  Node *temp1 = curr->nex();
  if(!temp1)
    return; //ERROR case when index is gerater than List size
  Node *temp2 = curr->nex()->nex();
  curr->setNext(temp2);
  delete temp1;
  if(!temp2) //case when index is the last node
    tail = curr;
  
  return;
}

bool List::isCyclePresent()
{
  //if NULL  
  if(!this->head)
    return false;
  //case when 1 node points to itself
  //if(this->head->nex()==this->head)
  //return true;

  //rest cases:
  Node *n1, *n2;
  n1 = head;
  n2 = head->nex();
  
  //iterate till end of the list if reached; check for n2 reaching NULL or before
  while(n1 && n2 && (n2->nex())){
    //cout << "n1: " << n1->val()<< "\tn2: " << n2->val()<< endl;
    if(n1==n2)
      return true;
    n1 = n1->nex();
    n2 = n2->nex()->nex();
  }
  
  return false;
}

string List::getMiddleValue()
{
  if (!this->head)
    return "ERROR";
  
  Node *p1, *p2;
  p1 = p2 = this->head;

  //cout << "head: " << this->head->val() << "tail: " << this->tail->val() << endl;

  //iterate till p2 reaches end or one before
  while ((p2!=this->tail) && (p2->nex()!=this->tail))
    {
      p1 = p1->nex();
      p2 = p2->nex()->nex();
    }
  return p1->val();
}
//return pointer to the nth value
Node* List::getNthPtr(int n)
{
  //can't be less than tail i.e. 1
  if (n<1)
    return NULL;

  int cnt = 1;
  Node *curr = this->head;

  //iterate till cnt is reached
  while ((cnt!=n) && curr)
    {
      curr = curr->nex();
      ++cnt;
    }

  return curr;
}


//one pass function
string List::getNthValueFromEnd(int n)
{
  //can't be less than tail i.e. 1
  if (n<1)
    return "ERROR";

  int cnt = 1;
  Node *curr = this->head;

  //iterate till cnt is reached
  while ((cnt!=n) && curr)
    {
      curr = curr->nex();
      ++cnt;
    }
  //if end of list is reached
  if(!curr)
    return "ERROR";

  //cnt == n; set a new pointer
  Node *p = this->head;

  //complete the traversal
  while(curr!=this->tail){
    curr = curr->nex();
    p = p->nex();
  }
  return p->val();
}

//Recursive version
void List::reverseListRecursive()
{
  //return if head ==NULL or only node 
  if((!this->head) || (!this->head->nex()))
    return;

  //else
  Node *first = this->head;
  Node *rest = this->head->nex();
  //crate rest list
  List *rst = new List(rest);
  
  //rcursive call
  rst->reverseListRecursive();
  first->nex()->setNext(first);
  first->setNext(NULL);
  this->head=rst->head;
  
}

void List::detectAndRemoveCycle()
{
	if(!this->head)
		return;
	//first detect if there is a cycle
	Node *sp=this->head;
	Node *fp=sp->nex();
	Node *p=NULL; //pointer to where slow and fast p's merge

	while(sp && fp && fp->nex())
	{
		if(sp==fp)
		{
			//cycle detected at p
			p=sp;
			break;
		}
		sp=sp->nex();
		fp=fp->nex()->nex();
	}
	if(!sp || !fp || !fp->nex())
		return; //not a cyclic linked list
        else
            cout << "List has cycle detected at " << p << endl;
	//count the number of nodes in the cycle
	
        Node *p2=p->nex();
	
        int cnt =1;
	while(p2!=p)
	{
		p2=p2->nex();
		cnt++;
	}
        cout << "Count of loop nodes: " << cnt << endl;
        
	//detect the first and last node of cycle
	int num=1;
	p=p2=this->head;
        Node *prev=NULL; //prev node pointer
        //set the pointer p2 to cnt+1 th node
	while(num!=cnt+1)
	{
		p2=p2->nex();
                num++;
	}
        //let n be the num nodes in the noncyclical part
        //p2 is thus at (cnt+1)-n element in the cycle
        //start moving p from head and p2 from cnt+1; after n hops they will meet at start of cycle
        /**/
        while(p!=p2)
        {
            p=p->nex();
            prev=p2;
            p2=p2->nex();
            
        }
        /**/
        
        //change the last cycle node to NULL
        prev->setNext(NULL);

	return;

}

int main()
{
  Node *tl = new Node("f", NULL);
  Node *hd = new Node("a", new Node("b", new Node("c", new Node("d", new Node("e", tl)))));
  
  List *lt = new List(hd, tl);
  //lt->insertAfterTail("w");
  lt->insertBeforeHead("p");
  /***
  List *lt = new List(NULL, NULL);
  lt->printList();
  
  lt->insertBeforeHead("q");
  lt->insertBeforeHead("r");
  lt->insertAfterTail("t");
  
  lt->printList();
  cout << "Value: " << lt->valueAt(0) << endl;
  lt->deleteNodeAt(3);
  ***/
  //test a cycle
  //lt->create_cycle();

  //create a cycle with 3rd node
  lt->createCycle( lt->getNthPtr(3) );

  //lt->printList();

  //if (lt->isCyclePresent())
    //cout << "List has a cycle" << endl;
  
  lt->detectAndRemoveCycle();
  lt->printList();
  
  //cout << "Middle value: " << lt->getMiddleValue() << endl;

  //cout << "3rd last element: " << lt->getNthValueFromEnd(3) << endl;;

  //lt->reverseListRecursive();
  //lt->printList();
}
