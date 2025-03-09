First phase benchmark runs around 1.17s

Second phase benchmark runs around 1.13s with 6 consumers
Time is similar since we are not really bound by the speed of consumers but by the reading of file.
We could load the entire file into memory but then again it depends on the size of the file.

Third phase benchmark runs at around 1.1s with 4 producers 2 consumers
Since json decoder reads data directly from the buffer to variable, there is really no way to speed it up concurently. 
The entire read would need to lock resources, and then unlock it again. 
Any boost would be only from the access to the file, but I think it would be marginal.
Due to that, buffio was used, but it involes way more hassle since we need to now split the file properly.
This approach however speed ups the reads since they do not need to synchronize on file access and can read at the same time.
We need to calculate how many bytes to read, make sure to recover the } signs etc. 
More explanation in the code.

Off course given the size of the data we could just load the entire file into memory first.. but I'm not sure what the concurency would showcase them ;).

Make file provides the commands to run the Repo. 
The code is ment to be run from the root of the repo, to avoid this we would need to allow passing of path to pl-172 file.