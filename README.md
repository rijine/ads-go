# Readme

## Graphql
### Schema Changes
run ```go run cmd/gen/bson.go```

### File Uploads 
Using postman [form-data]


#### Single file
``` 
POST
{ 
   operations: {"query":"mutation ImageUpload($picture: Upload!) {\n  uploadImage(picture: $picture, kind:\"post\")\n}","variables":{"picture":null}}
    map : { "0":  ["variables.picture"]}
    0: <File>
}
``` 
#### Multiple Files
``` 
POST
{ 
   operations: {"query":"mutation($req: [Upload!]) {\n  uploadImages(pictures: $req)\n}\n", "variables": { "req": [ null, null ] }}
    map : { "0": ["variables.req.0"], "1": ["variables.req.1"] }
    0: <File>
    1: <File>
}
``` 
