# GoLang: How to create package which can be referenced by main goapp.

## How to create package.?
 - Your go source code must be under the $GOPATH/src folder.
 - Go inside the source code folder and run ```$go install``` it add file under **pkg** with extension ```.a```  
 
## How to referencing the package in main/test go file?  
 - ```import (<complete path of package>)```  
 
## How to Run this go app?  
 - Install using ```go get``` as mentioned below, this will clone the project to **$GOPATH**  
   - ```$go get github.com/lovababu/golang-calculator```  
   
 - Installing **ginkgo and gomega** fallow this [https://onsi.github.io/ginkgo/](Link)  
 
 - ```$go test``` to run the test case.
 