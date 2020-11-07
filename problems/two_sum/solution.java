class Solution {
    public int[] twoSum(int[] nums, int target) {
      Map<Integer,Integer> maps=new HashMap<>();
        for(int i=0;i<nums.length;i++)
        {
            int complement=target-nums[i];
            if(maps.containsKey(complement))
            {
                return new int[]{maps.get(complement),i};
            }
            maps.put(nums[i],i);
        }
        throw new IllegalArgumentException("No two sum solution");
    }
}