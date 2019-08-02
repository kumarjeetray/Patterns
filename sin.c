#include<stdio.h>
#include<math.h>

void main()
{
	int n,i,x1,j;
	float x,sign,sinx,fact;
	printf("Enter the numbers\n");
	scanf("%d %f",&n,&x);
	x1=x;
	x=x*(3.14159/180);
	sinx=1;
	sign=1;
	for(i=2;i<=n;i++)
	{	
		fact=1;
		for(j=1;j<=i;j++)
		{
			fact=fact*j;
		}	
		sinx=sinx+(pow(x,3)/(fact+2)*sign);
		sign=sign*(-1);
	}
	printf("The sum series of sin is : %.4f\n",sinx);
	printf("The value of sin(%d)using : %f\n",x1,sin(x));
}
