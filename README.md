First phase benchmark runs around 1.17s.

Second phase benchmark runs around 1.13s with 6 consumers. The time is similar as we are not really constrained by the speed of the consumers, but by the reading of the file.
We could load the whole file into memory, but that depends on the size of the file.

The third phase of the benchmark runs at around 1.1s with 4 producers and 2 consumers.
 Since the json decoder reads data directly from the buffer into the variable, there is really no way to speed it up concurrently. 
The entire read would have to lock and unlock resources. 
Any speedup would only come from accessing the file, but I think it would be marginal.
This is why buffio was used, but it is much more tedious as we now have to split the file properly.
This approach does speed up the reads, because they do not need to synchronise on file access and can read at the same time.
We have to calculate how many bytes to read, make sure to recover the } characters, etc. 
More explanation in the code.


Make file provides the commands to run the repo. 
The code is meant to be run from the root of the repo, to avoid this we would need to allow passing the path to the pl-172 file.