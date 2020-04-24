import java.util.*;
class Solution {
    public boolean checkValidString(String s) {
      Stack <Character> bracket=new Stack<Character>();
        for(int i=0;i<s.length();i++)
        {
            if(s.charAt(i)=='(' || s.charAt(i)=='*')
                bracket.push(s.charAt(i));
            else{
                if(bracket.isEmpty()){
                    return false;
                }
                else{
                    int count=0;
                    while(bracket.peek()=='*'){
                        count++;
                    bracket.pop();
                    if(bracket.isEmpty())
                        break;
                }
                if(bracket.isEmpty()&&count==0)
                    return false;
                if(bracket.isEmpty())
                    count--;
                else
                    bracket.pop();
                while(count>0){
                    bracket.push('*');
                    count--;
                }
            }
        }
    }
   int count=0;
        while(!bracket.isEmpty()){
            if(bracket.peek()=='('&&count==0){
                return false;
            }
            else if(bracket.peek()=='*'){
                count++;
                bracket.pop();
            }
            else if(bracket.peek()=='('){
                count--;
                bracket.pop();
            }
        }
        return true;
}
}