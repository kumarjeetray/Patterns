/* 
    555555555
     4444444
      33333
       222
        1
	*/
#include<stdio.h>

void main()
{
	int n=5,i,j,c=1,x;
	c=n-1;
	x=n;
	for(i=1;i<=n;i++)
	{
		for(j=1;j<=c;j++)
		{
			printf(" ");
		}
		c++;
		for(j=1;j<=2*(n-i)+1;j++)
		{	
			printf("%d",x);
		}
		x--;
		printf("\n");
	}
}
