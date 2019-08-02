/*
 567895678
  4567456
   34534
    232
     1 
	 
	 */
	#include<stdio.h>

void main()
{
	int n=5,i,j,l,y=n,x=n,k;
	for(i=1;i<=5;i++)
	{	
		for(j=1;j<=i;j++)
		{
			printf(" ");
		}
		for(k=1;k<=n;k++)
		{
			printf("%d",x);
			x++;
		}
		for(l=2;l<=n;l++)
		{
			printf("%d",y);
			y++;
		}
		--n;
		x=n;
		y=n;
		printf("\n");
	}
}
