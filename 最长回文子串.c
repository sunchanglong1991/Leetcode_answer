

#include <stdio.h>
#include <string.h>
#include <stdlib.h>


int Max(int len1, int len2)
{
    return len1>len2?len1:len2;
}

int findchildstring(char *s, int left, int right)
{
    int  len =  strlen(s);
    if(len <= 1)
        return len;
    int L = left;
    int R = right;
    while(L >= 0 && R < len && s[L] == s[R])
    {
        L--;
        R++;
    }
    printf("R = %d;  L = %d\n", R, L);
    return R-L-1;
}

char *printchildstring(char *s, int start, int size)
{
    int len = strlen(s);
    if(len < start+size)
        return "";
    char * buf = malloc(size+1);
    memset(buf, '\0', size+1);
    strncpy(buf, s+start, size);
//    buf[size] = '\n';
    return buf;
}

char *longestPalindrome(char *s)
{
    int len = strlen(s);
    if(len <= 1)
        return s;
    int i, start =0, end=0, max = 0, mlen = 0;
    for(i=0; i<len; i++)
    {
        int len1= findchildstring(s, i,i);
        int len2= findchildstring(s, i, i+1);
        mlen = Max(Max(len1, len2), mlen);
        printf("len1 = %d; len2 = %d; mlen = %d\n",len1, len2, mlen);
        if(mlen > (end-start+1))
        {
            start = i - (mlen-1)/2;
            end = i + mlen/2;
        }
    }
    return printchildstring(s, start, mlen);
}


int main()
{
    char str[] = "babad";
    char str_2[] = "bb";
    printf("%s --> %s\n", str, longestPalindrome(str));
}


