When the elements are the same as the candidate element, votes are incremented 
whereas when some other element is found (not equal to the candidate element), 
we decreased the count. This actually means that we are decreasing the priority 
of winning ability of the selected candidate, since we know that if the 
candidate is in majority it occurs more than N/2 times and the remaining 
elements are less than N/2. We keep decreasing the votes since we found some 
different element(s) than the candidate element. When votes become 0, this 
actually means that there are the equal  number of votes for different elements
, which should not be the case for the element to be the majority element. 
So the candidate element cannot be the majority and hence we choose the present 
element as the candidate and continue the same till all the elements get 
finished. The final candidate would be our majority element. We check using 
the 2nd traversal to see whether its count is greater than N/2. 
If it is true, we consider it as the majority element.

Suppose that the input contains a majority element x, whose number of copies 
is more than half of the input length. Then the algorithm should return x as 
its final value m. Boyer and Moore prove the following stronger statement: 
For the final values m and c of the algorithm, on a sequence of length n, 
it is always possible to partition the input values into c copies of m and 
(n − c)/2 pairs of unequal items. From this it follows that no element x ≠ m 
can have a majority, because x can have at most half of the elements in the 
unequal pairs and none of the remaining c copies of m. Thus, if there is a 
majority element, it can only be m

To prove the existence of this partition, Boyer and Moore proceed by induction 
on the length of the input sequence. At each step, they use the induction 
hypothesis to find a partition of the same type for the subsequence of the 
input with its last item removed, with its value of m. If the final input 
value also equals m, it is added to the set of c copies of m. If it is unequal 
to m, and c was positive, then one of the c copies of m is removed from this 
set of copies and paired with the final value. And if c was zero, 
then the final value can be added into the (previously empty) set of copies 
of m, even if this step causes m to change. Thus, in all cases, the partition 
described by Boyer and Moore continues to exist.
