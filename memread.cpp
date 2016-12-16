#include <iostream>
#include <string>
#include <stdio.h>


using namespace std;

#define BLK_SIZE 4096
#define DISK_NUM 1

void *Read4kBlock(int disk, int blk_num, void *buffer)
{
  cout << "Reading disk " << disk << " block# " << blk_num << endl; 
}

void *mymemcpy(void * dest, void *src, size_t num)
{
  cout << "Mem copying " << num << " bytes from " << src << " to " << dest << endl;
}


void *ReadBlock(int disk, unsigned long offset, unsigned long len, void *buffer)
{
  unsigned long blk_idx, blk_off, len_read, size_toread, buf_off;
  blk_idx = offset/BLK_SIZE;
  blk_off = offset%BLK_SIZE;
  size_toread = len;
  void *tbuf = new char[BLK_SIZE]; 
  //void *tbuf = malloc(BLK_SIZE*sizeof(char)); 
  buf_off = 0;
  
  while(size_toread!=0)
    {
      len_read = BLK_SIZE-blk_off; //size availbale to read from current block
      len_read = (size_toread>len_read)?len_read:size_toread; //min of reaming size to read and size available to read from current block
      Read4kBlock(disk, blk_idx, tbuf);
      mymemcpy((buffer+buf_off), tbuf, len_read);
      blk_idx++;
      size_toread -= len_read;
      blk_off = 0;
      buf_off = len_read;
    }
  
}


int main()
{
  unsigned long offset, nsize;
  cout << "Print the offset to read from and num bytes" << endl;
  cin >> offset >> nsize;
  void *mbuf = new char[nsize+1];
  ReadBlock(DISK_NUM, offset, nsize, mbuf);  

  return 0;
}



