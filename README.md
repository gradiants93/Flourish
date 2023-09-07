### CRUD Commands
 - `deleteCustomer --customerId`    
   - delete customer and all related orders
 - `deleteOrder --orderId`        
   - delete order by id
 - `getAllCustomer`     
   - Gets all customer info
 - `getAllOrder`        
   - Get all orders
 - `getCustomer --customerId`        
   - Gets customer info via id
 - `getOrderByCustomer --customerId` 
   - gets all orders by customer id
 - `getOrderByID --orderId`       
   - get single order by id
 - `help`               
   - Help about any command
 - `insertCustomer --name --email`  
   - Insert a new customer 
 - `insertOrder --customerId --quantity --totalPrice`        
   - insert a new order
 - `updateCustomer --customerId --name --email`     
   - update customer name and email by id 
 - `updateOrder --orderId --customerId --date --quantity --totalPrice`
   -  update order by id

### Write-up
1. Setup project space and install go, cobra
   -  45 mins
2. Setup database and schema
   - 30 mins
   - Needed to install mysql and had some syntax errors
3. Create first command - get customer by id
   - 45 mins
   - Most of my time was spent learning how to code in go. The syntax is different from what I am used to. Looked through official documentation, guides, and stackoverflow for reference and error code resolution.
4. Create further commands
   - 2 hours 30 mins
   - Ran into difficulties with how I had named `order` table. I tested functionality of each command as I went to make sure they were behaving as expected. Also spent some time trying to implement flags before reverting to passing naked arguments.
5. README
   - 45 mins
6. Went back and added flags to each function. Had to read a few guides on how to do this. I also had to take out validation that arguements were passed due to this change. Next step would be figuring out how to confirm that values were passed to flags.
   - 1 hour
   
### TODOS & Next Steps
1. Add in unit tests for each CRUD function
2. Add more validation for values passed to flags
3. Create end-to-end test
4. Potentially change to use some kind of cli library