# poptest2
### Source of book tutorial
https://andrew-sledge.gitbooks.io/the-unofficial-pop-book/content/

## Have a docker mysql with 

```
sudo docker run --name=mysql -e MYSQL_ROOT_PASSWORD=root -p=3306:3306 -d mysql:5.5
``` 

## Generating stuff 
```
soda g config -t mysql
soda create -a
soda generate model user title:string first_name:string last_name:string bio:text -e development
soda migrate up -e development
soda generate fizz add_location_column
```
### add the below line in the migration file ( fizz ) just generated
```
add_column("users", "location", "string", {"size": 100, "default": "New York, NY"})
```
### run migrate to apply
```
soda migrate up -e development
soda generate fizz add_image
```
### add the below line in the migration file ( fizz ) just generated
```
add_column("users", "image", "string", {"default": "https://upload.wikimedia.org/wikipedia/commons/thumb/0/04/MarvelLogo.svg/1200px-MarvelLogo.svg.png"})
```
### run migrate to apply
``` 
soda migrate up -e development
``` 
### run the below command to generate favorite_food table
```
soda generate model favorite_food user:uuid food:text -e development
soda migrate up -e development
```

