import java.util.*;
class Solution {
    public int findMaxLength(int[] nums) {
        HashMap<Integer,Integer> hs = new HashMap<>();
        hs.put(0,-1);
        int sum = 0;
        int max = -1;
        for(int i = 0; i < nums.length; i++)
            if(nums[i] == 0)
                nums[i] = -1;
        for(int j = 0; j < nums.length; j++){
            sum += nums[j];
            
            if(hs.containsKey(sum)){
                int go = hs.get(sum);
                if(max < j - go)
                    max = j - go;
            }
            else
                hs.put(sum,j);
        }
        if(max == -1)
            return 0;
        return max;
    }
    
}