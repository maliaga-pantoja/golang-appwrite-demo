# Golang-appwrite-demo
<a href="https://appwrite.io/" target="_blank" style="position: fixed;right: 18px;bottom: 18px;z-index: 999;">
  <img style="width: 160px;" src="https://appwrite.io/images-ee/press/badge-pink-button.svg" alt="Built with Appwrite">
</a>

Golang demo using appwrite. This is a guide about how to use appwrite rest api for database resource with golang. 
## AppWrite
First, you need to run appwrite. I recommend to you to follow the official [appwrite guide](https://appwrite.io) in the getting start section.
Its a shortcut ussing docker:
```
docker run -it --rm \
    --volume /var/run/docker.sock:/var/run/docker.sock \
    --volume "$(pwd)"/appwrite:/usr/src/code/appwrite:rw \
    --entrypoint="install" \
    appwrite/appwrite:0.10.4
```
## How to
run the script
```
/bin/sh run.sh
```
## What i need to modify?
You need to to modify env vars set in run.sh using your own
- BASEPATH
- PORT
- PROJECT_ID
- TOKEN
## How it works ?
It project use the appwrite provided apis. Fow now, only database endpoint are implement in this project. <br>
### Procedures
- Create new collection
- List all avaiable collections
- Create a new document using the id of inserted collection
## Tips
Is required to set payload for create collection in the format below
```
{
  "data": {
    "key": "value3"
  } 
}
```
The data key is mandatory. *key* is the registered key in Collection.CollectionCreateInputRule. <br> 
If you modify the value of key, modify it in the create collection payload too. <br>

### Example
```	
inputRule := Collection.CollectionCreateInputRule{
  Label:    "label",
  Key:      "myKey",
  Type:     "text",
  Default:  "no set",
  Required: false,
	Array:    false,
}

{
  "data": {
    "myKey": "value4"
  } 
}
```
The value of Label is for web interface