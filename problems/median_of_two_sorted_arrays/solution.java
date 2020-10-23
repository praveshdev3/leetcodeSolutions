class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        if(nums1.length>nums2.length)
        {
            int temp[]=nums1;
            nums1=nums2;
            nums2=temp;
        }
        int lo=0;
        int hi=nums1.length;
        int combined=nums2.length+nums1.length;
        
        while(lo<=hi)
        {
            int partX=(lo+hi)/2;
            int partY=(combined +1)/2 -partX;
            int leftX = getMax(nums1, partX);
            int rightX = getMin(nums1, partX);
            int leftY = getMax(nums2,partY);
            int rightY = getMin(nums2, partY);
            
            if(leftX<=rightY && leftY<=rightX)
            {
                if(combined%2==0)
                    return (Math.max(leftX,leftY)+Math.min(rightX,rightY))/2.0;
                else
                    return Math.max(leftX,leftY);
            }
            if(leftX>rightY)
            {
                hi=partX-1;
            }else
            {
                lo=partX+1;
            }
        }
        return -1;
    }
        private int getMax(int[] nums,int partition)
        {
            if(partition==0)
            {
                return Integer.MIN_VALUE;
            }
            else
                return nums[partition-1];
        }
        private int getMin(int[] nums,int partition)
        {
            if(partition==nums.length)
            {
                return Integer.MAX_VALUE;
            }
            else
                return nums[partition];
        }
    
}