# greenjade-task


Before run the commands below please add values to config/development.yml file with values which I sent in email.

---
to run app you have 2 options:
 - with regular way just run go run main.go in your terminal;
 - or run docker-compose up
By default the app running on port :5001

## api doc
https://documenter.getpostman.com/view/1163851/UVsSL2pg

The data stored in AWS dynamodb. I attached sample images how viewed stored data in AWS interface:
![alt text](https://greenjadetask.s3.eu-central-1.amazonaws.com/Screen+Shot+2022-03-24+at+12.12.03.png)

![alt text](https://greenjadetask.s3.eu-central-1.amazonaws.com/Screen+Shot+2022-03-24+at+12.12.24.png)
 

##### The time estimation for each part which mentioned in the task:

- Part1: 7hours
- Part2: 2hours
- Part3: 5hours

#### Notes:
###### Part 2
-  we can make additional validation that set max limit dimensions count
- Maybe while we validate the 3 options we can check also is it possible to out of maze. so if there is no 0 value in dimensions we return error from validation
- there was only one 2 value in dimension

###### Part 3
    I think we must use searching algorithms here. 
    While I thinked about the part 3 and research it on internet there is 
    bfs and dfs algorithms which exactly use for find path for out of maze.
    
