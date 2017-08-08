## my-service

todo: come up with clever name

* Key value store with a service broker & api


### service-broker
* catalog: single plan
* create-instance
    * create bucket
    * store in some "main" bucket
* bind-instance
	* generate creds
	* store in some "main" bucket
* ubund-instance
    * remove creds from "main" bucket
* delete-instance
    * destroy bucket

### db
* boltdb - filesystem based key value store. multi tenant but no HA
* API (for devs, given out by broker)
    * get to bucket (with auth header)
    * put to bucket (with auth header)
* API for admin (used by SB)
    * create bucket instance
    * create bucket creds...?