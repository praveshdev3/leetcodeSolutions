class Solution {
    public int maxScore(String s) {
         int num_one=0;
        for(int i=0;i<s.length();i++)
        {
            if(s.charAt(i)=='1')
                num_one++;
        }
        int max=0;
        int num_zero=0;
        int Omax=Integer.MIN_VALUE;
        for(int i=0;i<s.length()-1;i++)
        {
            if(s.charAt(i)=='0')
            {
                num_zero++;
                max=num_zero+num_one;
            }
            else
            {
                num_one--;
                max=num_zero+num_one;
            }
            if(Omax<max)
                Omax=max;
        }
        return Omax;
    }
}