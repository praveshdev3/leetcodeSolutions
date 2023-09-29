func minPartitions(n string) int {
   var min_number byte = '0' 
   for i := 0; i < len(n); i++ {
       if n[i] > min_number {
           min_number = n[i]
       }
   }
   return int(min_number - '0')
}