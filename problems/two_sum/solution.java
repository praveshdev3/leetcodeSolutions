class Solution {
    public int[] twoSum(int[] nums, int target) {
        int arr[]=new int[2];
        Map<Integer,Integer> res=new HashMap<>();
        for(int i=0;i<nums.length;i++)
        {
            int collission=target-nums[i];
            if(res.containsKey(collission))
            {
                arr[0]=res.get(collission);
                arr[1]=i;
                return arr;
            }
            res.put(nums[i],i);
        }
        return arr;
    }
}