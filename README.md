## Backend for web-shop 'Masterium' (ITMO project)

---

Launch:
make build && make run

Stop:
make stop

---

API:

Get all item groups:<br/>
GET /api/item_groups
<br/><br/>
Get item group by id:<br/>
GET item_groups/\<id\>
<br/><br/>
Get items by item group id:<br/>
GET item_groups/\<id\>/items
<br/><br/>
Get all items:<br/>
GET /items
<br/><br/>
Get item by id:<br/>
GET /items/\<id\>

---

Database schema
<img src="db/schema.png">

Todo:
- Authorization
- Carts

Remarks:
Tried to use a clean architechture