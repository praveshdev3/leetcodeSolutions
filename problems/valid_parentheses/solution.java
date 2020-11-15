import java.util.*;
class Solution {
    public boolean isValid(String s) {
        if(s.length()%2!=0) return false;
      Stack <Character> st=new Stack<>();
        int i=0;
        while(i!=s.length())
        {
            if(s.charAt(i)=='(' || s.charAt(i)=='[' || s.charAt(i)=='{')
                st.push(s.charAt(i));
            else if(s.charAt(i)==')' && !st.isEmpty() && st.peek()=='(')
                st.pop();
           else if(s.charAt(i)=='}' && !st.isEmpty()  && st.peek()=='{')
                st.pop();
            else if(s.charAt(i)==']'&& !st.isEmpty() && st.peek()=='[' )
                st.pop();
            else 
                return false;
            i++;
        }
        if(st.isEmpty())
            return true;
        else 
            return false;
    }
}