#include <iostream>
//#include <cstdbool>
#include <list>
#include <queue>

#define MAX_NODES 10

using namespace std;

typedef struct gnode{
  int adj_v;
  int weight;
  //struct gnode *next;
}g_node;


class Graph{
public:
  //add edge from node u to node v 
  void addEdge(int u, int v, int wght);
  //initialize()
  void initGraph(int num_v, const bool gdirected);
  //start DFS traversal from node
  void DFSvisit(int node);
  //start BFS from v
  //void BFSvisit(int v);
  void printGraph();
  int numNodes();
  int numEdges();
  void printParents();
  void topSort();
  void topSort2();


private:
  int V; //num nodes
  int E; // edges
  int parent[MAX_NODES]; //parent nodes
  list<g_node> edges[MAX_NODES];
  bool directed;
  int degree[MAX_NODES];
  int in_degree[MAX_NODES];
  bool visited[MAX_NODES];
  bool processed[MAX_NODES];
  //stack used for DFS
  list<int> st;
  //private function to initialize DFS visited array
  void DFSinit();
  //DFS visit util from source s
  void DFSUtil(int s);
  //topological sort util
  void TopSortUtil(int node, list<int> &node_list);
  
};

//initialize this graph
void Graph::initGraph(int num_v, const bool gdirected)
{
  V=num_v;
  E=0;
  directed = gdirected;
  for(int i=0; i<MAX_NODES; i++)
    {
      //parent[i] = -1;
      degree[i] = 0; //undirected case
      in_degree[i] = 0; //directed case
      edges[i].clear();
    }
  
}

int Graph::numNodes()
{
  return V;
}

int Graph::numEdges()
{
  return E;
}

//manually add each edge from user
void Graph::addEdge(int x, int y, int w)
{
  //g_node p = (g_node*)malloc(sizeof(g_node));
  //p->
  g_node t;
  t.adj_v = y;
  t.weight = w;
  edges[x].push_front(t);
  degree[x]++;
  in_degree[y]++;
  
  if (!directed)
    {
      //addEdge(y, x);
      g_node q;
      q.adj_v = x;
      q.weight = w;
      edges[y].push_front(q);
      degree[y]++;
      //E++;
    }
  
  E++;
    
}


//print grapg as read
void Graph::printGraph()
{
  if (directed)
    cout << "\n\nDirected Graph:" << endl;
  else
    cout << "\n\nUndirected Graph:" << endl;

  //for each vertex
  for (int i = 0; i< V; i++)
    {
      cout << "Node: " << i;
      //cout << " In-Degree="<< in_degree[i] << " Out-Degree=" << degree[i] << endl;
      if(!edges[i].empty())
	{
	  cout << " Edges -->" ;
	  list<g_node>::iterator it;
	  for(it=edges[i].begin(); it != edges[i].end(); it++)
	    {
	      cout << " ("  << (*it).adj_v ;
	      cout << ", " << (*it).weight << ")";
	    }
	}
	cout << endl;
    }
 
}

void Graph::DFSinit()
{
  st.clear();
  for(int i=0; i<V; i++)
    {
      visited[i]=false;
      processed[i]=false;
      parent[i]=-1; // initialize parents to -1 
    }
}

//Depth First Search: uses the call stack instead of exclusive stack declaration
void Graph::DFSUtil(int s)
{

  visited[s] = true;

  cout << "Visited vertex: " << s << endl;
  list<g_node>::iterator it;
  for(it=edges[s].begin(); it!=edges[s].end(); it++)
    {
      int y = (*it).adj_v;
      if(!visited[y])
	{
	  parent[y]=s;
	  //cout << "Process edge: " << s << y << endl; 
	  DFSUtil(y);
	}
      else if (!processed[y]) //visited but not processed
	{
	  //this is a back edge to already visited node
	  //process edge s->y
	  //cout << "Process edge: " << s << y << endl; 
	  if(parent[y]!=s)
	    cout << "Back edge causing cycle: " << s << y << endl;
	}
	  
    }
  processed[s]=true;
  //cout << "Finished processing vertex: " << s << endl;
  
}


void Graph::DFSvisit(int s)
{
  	cout << "\n\nDFS visit of Graph from node " <<  s << endl;
  	//initialize the visited array
	DFSinit();
	//recursively call DFS for all nodes while they are unvisited
	DFSUtil(s);
	//if unvisited nodes call DFS on those
	for(int i=0; i<this->numNodes(); ++i)
		if(!visited[i])
			DFSUtil(i);

}

//Optimal topological sort based on DFS
void Graph::topSort2()
{
	cout << "\n\nTopoligcal Sort Ordering" << endl;
	//similar to DFS, but instead of printing the vertex first it will store the traversal in a stack
	list<int> top_list;

  	//initialize the visited array
	DFSinit();
	
	//while there are unvisited nodes in graph
	for(int i=0; i<this->numNodes(); ++i)
		if(!visited[i])
			TopSortUtil(i, top_list);

	//when all nodes are traversed the stack is printed
	for(auto x: top_list)
		cout << x << " ";
	cout << endl;
	return;
}

void Graph::TopSortUtil(int node, list<int> &node_list)
{
	list<g_node>::iterator it;
	for(it=edges[node].begin(); it!=edges[node].end(); ++it)
	{
		//edge
		int y=(*it).adj_v;
		//if not visited
		if(!visited[y])
			TopSortUtil(y, node_list);
	}
	//if all adjacent nodes visited 
	//add the starting node to front to list
	node_list.push_front(node);
	visited[node]=true;
	return;

}



/*****
//Bread First Search: iterative
void Graph::BFSvisit(int s)
{
  //initialize stakc, visited, processed, same as DFS
  DFSinit();

  //insert first node in stack
  st.push_front(s);
  visited[s] = true;
  //process(s)

  while(!st.empty())
    {
      int x=st.front();
      cout << "Strated processing vertex: " << x << endl;
      list<g_node>::iterator it;
      for(it=edges[x].begin(); it!=edges[x].end(); it++)
	{
	  int y = (*it).adj_v;
	  if(!visited[y])
	    {
	      //queue implementation, always push back
	      st.push_back(y);
	      parent[y]=x;
	      cout << "Process edge: " << x << y << endl; 
	      visited[y] = true;
	      //DFSvisit(y);
	    }
	  else if (!processed[y])
	    {
	      //disocvered but not processed
	      cout << "Process edge: " << x << y << endl; 
	
	    }
	}
      processed[x]=true;
      st.pop_front();
      cout << "Finished processing vertex: " << x << endl;
    }
  
  
}
*****/

void Graph::printParents()
{
  cout << "Parent Graph:" << endl;
  for (int j=0; j<V; j++)
    cout << parent[j] << " --> " << j << endl;
}

//topological sort of Directed Graph
void Graph::topSort()
{
  cout << "Printing top sort" << endl;
  
  //if undirected return
  if(!directed)
    {
      cout << "Undirected graph!!" << endl;
      return;
    }
  //maintain queue of vertices
  list<int> tsq;
  tsq.clear();
  int count = 0;
  
  // push vertices with in degree 0
  for (int i=0; i<V; i++)
    {
      if (in_degree[i] == 0)
	{
	  tsq.push_back(i);
	}
    }

  if(tsq.empty())
    {
      cout << "Graph has cycles!" << endl;
      return;
    }

  //traverse all vertices
  while(!tsq.empty())
    {
      int temp = tsq.front();
      cout << tsq.front() << " , ";
      tsq.pop_front();
      //modify in-degree for adj edges in-degree array
      list<g_node>::iterator it;
      for(it = edges[temp].begin(); it!=edges[temp].end(); it++)
	{
	  if((--in_degree[(*it).adj_v])==0)
	    {
	      tsq.push_back((*it).adj_v);
	    }
	  //if()
	}
    }
  
  cout << endl;

}



int main()
{
  int num_v=6, num_e=6;
  Graph G;
  //default to directed graph
  const bool g_directed = true;
  G.initGraph(num_v, g_directed);

  G.addEdge(5,0, 0);
  G.addEdge(5,2, 0);
  G.addEdge(4,1, 0);
  G.addEdge(4,0, 0);
  G.addEdge(2,3, 0);
  G.addEdge(3,1, 0);

  G.printGraph();
  G.DFSvisit(5);
  G.topSort2(); 

/****
  cout << "Enter the num vertices and num edges" << endl;
  cin >> num_v >> num_e;
  int x, y, w;
  cout << "Enter graph by its edges and weight " << endl;
  int i=1;
  while (i<=num_e)
    {
      cin >> x >> y >> w;
      G.addEdge(x,y, w);
      i++;
    }

  cout << "Graph Initialized with " << G.numNodes() << " nodes and " << G.numEdges() << " edges" << endl;
 ****/
  

  /*******
  cout << "DFS traversal of graph:" << endl;
  G.DFSinit();
  G.printParents();

  G.BFSvisit(0);
  G.printParents();
  ****/

  //G.topSort(); 
  
}

