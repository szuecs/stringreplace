For everyone that doesn't want to deal with complex escape sed syntax
and wants to find and replace regular expressions in files.

@ is the magic separator. Need something else? Create a PR or change
it in the file.

```
% go build -o stringreplace main.go
% cat file1
([.])?(:[0-9]+)?)$")
% ./stringreplace '(:[@foo' file1
% cat file1
([.])?foo0-9]+)?)$")
```
