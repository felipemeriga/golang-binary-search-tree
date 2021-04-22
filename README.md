# Golang Binary Search Tree

This is a Golang implementation of a binary search tree, with all the respective methods for searching, adding, max and min.

The binary tree is a data structure that has a lot of different usages among software development products,
and effectively this structure can bring the performance of your application to other levels, if used wisely.

Another reason that people needs to learn about binary trees, is they are the favorite data structure of coding tests
and interview, so if you are planning to do interviews in the future, it's nice to know about this data structure.

This repo has the functionalities of:

- Creating a binary tree of how many children nodes you want
- Print how the tree is currently structure
- Traverse using DFS(Inorder, Preorder, Postorder)
- Traverse using BFS
- Get the max value
- Get the minimum value
- Get Height

Additionally, this Golang app uses some features from the own language, like Locking functionalities,
to prevent different goroutines to race with each other, preventing another routines to write data, when a specific
one is consuming(read/write) this data.

This Golang mod has only a single dependency, [Genny](https://github.com/cheekybits/genny), which helps us for using 
generics with Golang.

# Running the App

First, you need to install the dependencies of the app, you can execute the following command in the root of the cloned repo:
```
go mod download
```

This will fetch the dependencies and place them on your ```GOPATH```.

The main code is located on ```main.go``` and the tests to execute and perform operations with binary trees,
are on ```main_test.go```.

So you need to execute the tests available in the test file, for example, for executing the test ```TestInsert```,
which is the main one for creating a simple tree, you can do:
```
go test ./ -run TestInsert -v
```

You can also run your app through an IDE, like Goland or VSCode, when using an IDE, remember to
enable modules integration configuration, so the modules that are installed through the command line, or even
directly through IDE, will be synced and indexed from GOPATH to your IDE, and you will be able to run all the tests.



