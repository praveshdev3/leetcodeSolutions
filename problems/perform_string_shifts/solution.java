class Solution {

    public String stringShift(String s, int[][] shift) {
        int n = shift.length;
        int Right = 0;
        int Left = 0;
        for(int i = 0; i < n; i++){
            if(shift[i][0] == 0){
                Left += shift[i][1];
            }
            if(shift[i][0] == 1){
                Right += shift[i][1];
            }
        }
        if(Right > s.length()){
            while(Right > s.length()){
                Right = Right - s.length();
            }
        }
        if(Left > s.length()){
            while(Left > s.length()){
                Left = Left -  s.length();
            }
        }
        if(Right > Left)
        {
            int d = Right - Left;
            return leftRotate(s, s.length() - d);
        }
        else{
            int d = Left - Right;
            return leftRotate(s, d);
        }
    }
        static String leftRotate(String str, int d) 
    { 
            String ans = str.substring(d) + str.substring(0, d); 
            return ans; 
    } 
}