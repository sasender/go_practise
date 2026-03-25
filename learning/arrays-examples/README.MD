### Go Arrays
*** Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value. ***

### Syntax

var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred

### 2. With the := sign:
*** Syntax
array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred ***

*** Array Size = Index Range Rule ***
arr1 := [5]int{1,2,3,4,5}
        ↑
     Size 5 = indices 0,1,2,3,4 ONLY

*** Visual Breakdown  ***
Index:  0  1  2  3  4
Values: 1  2  3  4  5
Size:     ↑
         [5]   

### Array = Copy data   
### Slice = Share data


