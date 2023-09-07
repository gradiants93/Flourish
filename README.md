### CRUD Commands
 - `deleteCustomer <customer_id>`    
   - delete customer and all related orders
 - `deleteOrder <order_id>`        
   - delete order by id
 - `getAllCustomer`     
   - Gets all customer info
 - `getAllOrder`        
   - Get all orders
 - `getCustomer <customer_id>`        
   - Gets customer info via id
 - `getOrderByCustomer <customer_id>` 
   - gets all orders by customer id
 - `getOrderByID <order_id>`       
   - get single order by id
 - `help`               
   - Help about any command
 - `insertCustomer <name> <email>`  
   - Insert a new customer 
 - `insertOrder <customer_id> <qty_ordered> <total_price>`        
   - insert a new order
 - `updateCustomer <customer_id> <name> <email>`     
   - update customer name and email by id 
 - `updateOrder <cutomer_id> <customer_id> <date> <qty_ordered> <total_price>`
   -  update order by id

### Write-up
1. Setup project space and install go, cobra
   1. 45 mins
2. Setup database and schema
   1. 30 mins
   2. Needed to install mysql and had some syntax errors
3. Create first command - get customer by id
   1. 45 mins
   2. Most of my time was spent learning how to code in go. The syntax is different from what I am used to. Looked through official documentation, guides, and stackoverflow for reference and error code resolution.
4. Create further commands
   1. 2 hours 30 mins
   2. Ran into difficulties with how I had named `order` table. I tested functionality of each command as I went to make sure they were behaving as expected. Also spent some time trying to implement flags before reverting to passing naked arguments.
5. README
   1. 45 mins
   
### TODOS & Next Steps
1. Add in unit tests for each CRUD function
2. Create end-to-end test
3. Get flags working or use some kind of cli library