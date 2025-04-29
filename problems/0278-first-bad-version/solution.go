/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    // for i:=n-1; i>=1; i--{
    //     if !isBadVersion(i){
    //         return i+1
    //     }
    // }
    // return 1
    start,end := 1,n
    for start<=end{
        mid := (start+end)/2
        if isBadVersion(mid){
            if !isBadVersion(mid - 1) {
                return mid
            }
            end = mid-1
        }else{
            start = mid+1
        }
    }
    return -1
}
