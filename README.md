## Backend for web-shop 'Masterium' (ITMO project)

---

Launch:
make build && make run

Stop:
make stop

---

API:

Get all item groups:<br/>
GET /api/catalog/groups
<br/><br/>

Get item group by id:<br/>
GET /api/catalog/groups/\<id\>
<br/><br/>

Get item groups by catgory id:<br/>
GET /api/catalog/categories/\<id\>/items
<br/><br/>

Get items by item group id:<br/>
GET /api/catalog/groups/\<id\>/items
<br/><br/>

Get all items:<br/>
GET /api/catalog/items
<br/><br/>

Get item by id:<br/>
GET /api/catalog/items/\<id\>
<br/><br/>

Get categories:<br/>
GET /api/catalog/categories
<br/><br/>

Get category by id:<br/>
GET /api/catalog/categories/\<id\>

---

Database schema
<img src="db/schema.png">

Todo:
- Authorization
- Carts

Remarks:
Tried to use a clean architechture