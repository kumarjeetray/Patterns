/*
    123456789
     1234567
      12345
       123
        1
*/
#include<stdio.h>

void main()
{
	int n=5,i,j,c=1;
	c=n-1;
	for(i=1;i<=n;i++)
	{
		for(j=1;j<=c;j++)
		{
			printf(" ");
		}
		c++;
		for(j=1;j<=2*(n-i)+1;j++)
		{
			printf("%d",j);
		}
		printf("\n");
	}
}
