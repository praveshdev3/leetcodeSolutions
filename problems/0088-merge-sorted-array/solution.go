func merge(nums1 []int, m int, nums2 []int, n int)  {
    i:=m-1
    j:=n-1
    for i>=0 && j>=0{
        if nums1[i] > nums2[j]{
            nums1[i+j+1] = nums1[i] 
            i--
        }else{
            nums1[i+j+1] = nums2[j] 
            j--
        }
    }

    for j>=0{
        nums1[j] = nums2[j] 
        j--
    }

}
