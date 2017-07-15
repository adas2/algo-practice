#include <iostream>
//#include <list>


using namespace std;

struct Node{
     struct Node *next;
     struct Node  *rand;
};
 
typedef struct Node *list;


void printList(list head)
{
	while (head)
	{
		cout << "Loc: " << head << " Rand: " << head->rand << " --> ";
		head=head->next;
	}
	cout << "NULL" << endl;
}

//clone list l1 to list l2
void cloneList (list* l1 , list* l2)
{    
     struct Node *curr= *l1;
     //struct Node *newList=NULL;

     //NULL list
     if (!curr)
     {
	*l2=NULL;
	return;
     }

    //first pass to create empty cloned list
     while (curr->next)
     {
         //hold the next poiter
         struct Node *temp = curr->next;
         //create new node
         struct Node *newNode = new (struct Node);
         newNode->rand=NULL;
          
         newNode->next=temp;
         curr->next=newNode;
         //point next pointer of original list to the new node
	curr=temp;
     }
     //last node of original list
     struct Node *newNode = new (struct Node);
     newNode->rand=NULL;
     newNode->next=NULL;
     curr->next = newNode;

     
     //second pass to fill the random pointers for the cloned list
     curr = *l1;
     //struct Node *curr2 = newHead;
     while(curr && curr->next)
     {
         curr->next->rand = curr->rand->next;
         curr=curr->next->next;
     }
     
     //third pass to change back original pointers, break the cloned list
     curr = *l1;
     *l2=curr->next;
     while(curr->next)
     {
	struct Node* temp = curr->next;
	curr->next=curr->next->next;
	curr=temp;
     }
  
}


int main()
{
	//create rand list
	list l1=NULL,l2=NULL;
	struct Node* a1 = new (struct Node);
	struct Node* a2 = new (struct Node);
	struct Node* a3 = new (struct Node);
	a1->rand=a1;
	a1->next=a2;
	a2->rand=a3;
	a2->next=a3;
	a3->rand=a2;
	a3->next=NULL;
	
	l1=a1;

	printList(l1);

	cloneList(&l1,&l2);

	printList(l2);

}
