class Solution {
    public int strStr(String haystack, String needle) {
        if(needle.isEmpty())
            return 0;
        if(haystack.length()<needle.length())
            return -1;
        for(int i=0;i<=haystack.length()-needle.length();i++)
        {
            int j;
            for(j=0; j<needle.length();j++)
            {
                if(haystack.charAt(i+j)!=needle.charAt(j))
                    break;
            }
            if(j==needle.length())
                return i;
        }
        return -1;
    }
}