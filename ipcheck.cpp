#include<iostream>
#include <string>
#include <stdlib.h>

using namespace std;


bool valid_ip_check(string ip)
{
  //edge case:
  //check if len is zero or more than 15
  if((ip.length()<=0)||(ip.length()> 15) )
    return false;
  //read the 4 parts of the ip
  string rest = ip;
  int dot_cnt = 0;
  
  while(rest.length()>0){
    string addr = "";
    //position of the first '.'
    size_t pos = rest.find(".");
    if(pos != string::npos)
      {
	//part before '.'
	addr = rest.substr(0,pos);
	//parts after '.'
	rest = rest.substr(pos+1);
	cout << "Part: " << addr << endl;
	dot_cnt++;
	
	//check for empty string addr
	if(addr.length() ==0)
	  return false;

	//check value
	int val = atoi(addr.c_str());
	if((val>255)||(val<0))
	  return false;
	//check addr for leading zero for nonzero value
	if ((val!=0) && addr[0]=='0')
	  return false;

      }
    else // no dots found, must be last address part
      {
	addr = rest;
	cout << "Part: " << addr << endl;
	
	//check for empty string addr
	if(addr.length() ==0)
	  return false;
	//check value
	int val = atoi(addr.c_str());
	if((val>255)||(val<0))
	  return false;
	//check addr for leading zeros for non-zero case
	if ((val!=0) && addr[0]=='0')
	  return false;

	break;
      }
    
  }
  if(dot_cnt!=3)
    return false;

  return true;

}

int main()
{
  string ip_addr;
  cout << "Enter IP address to check" << endl;
  cin >> ip_addr ; 
  if(valid_ip_check(ip_addr))
    cout << "Given IP is valid" << endl;
  else 
    cout << "IP Invalid" << endl;
}
