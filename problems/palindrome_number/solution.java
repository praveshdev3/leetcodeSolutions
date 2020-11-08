class Solution {
    public boolean isPalindrome(int x) {
     int reversed=0;
        if(x==0)
            return true;
        if(x<0||x%10==0)
            return false;
        while(x>reversed)
        {
            int temp=x%10;
            x/=10;
            reversed=reversed*10+temp;
        }
        if(x==reversed||x==reversed/10)
            return true;
        else
            return false;
    }
}