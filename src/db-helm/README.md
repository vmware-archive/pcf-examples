
### Demo/Example Script

```bash
helm install .

kubectl get deployments
kubectl expose deployment "<deployment name>" --type=NodePort

kubectl get services
export host="<service ip & port value>"

echo "Create bucket and user"
curl -u admin_db:monkey123 -X POST http://$host/api/admin/bucket/my_bucket
curl -u admin_db:monkey123 -X PUT http://$host/api/admin/bucket/my_bucket/credentials -d '{
	"username": "user",
	"password": "pass"
}'

echo "Empty bucket contents"
curl -u user:pass -X GET http://$host/api/bucket/my_bucket/

echo "Add to bucket"
curl -u user:pass -X PUT http://$host/api/bucket/my_bucket/some_val -d 'Test value'

echo "New bucket contents"
curl -u user:pass -X GET http://$host/api/bucket/my_bucket/
```
