# Golang-appwrite-demo
golang demo using appwrite

<!-- 'Made with Appwrite' badge -->
<a href="https://appwrite.io/" target="_blank" style="position: fixed;right: 18px;bottom: 18px;z-index: 999;">
  <img style="width: 160px;" src="https://appwrite.io/images-ee/press/badge-pink-button.svg" alt="Built with Appwrite">
</a>
<br> <br>
This is a guide about how to use appwrite rest api with golang. 
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
