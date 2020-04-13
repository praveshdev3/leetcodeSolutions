class Solution {
    public int maxSubArray(int[] nums) {
int max= Integer.MIN_VALUE, max_new= 0;  
        // if(nums.length==1)
        //     return nums[0];
    for (int i = 0; i < nums.length; i++)  
    {  
        max_new= max_new + nums[i]; 
           
         if (max < max_new)  
            max = max_new;
        if (max_new < 0)  
             max_new = 0;
    }  
    return max;
    }
}