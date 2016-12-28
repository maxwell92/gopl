echo "POST"
curl -XPOST -d "{\"name\":\"Justin\"}" localhost:8081/hello  

echo "PUT"
curl -XPUT -d "{\"name\":\"Justin\"}" localhost:8081/hello  

echo "DELETE"
curl -XDELETE -d "{\"name\":\"Justin\"}" localhost:8081/hello  

echo "PATCH"
curl -XPATCH -d "{\"name\":\"Justin\"}" localhost:8081/hello  

echo "GET"
curl -XGET localhost:8081/hello  

