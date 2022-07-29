# My 'C-like' Interpreter

This project is based on the book [*Writing An Interpreter In Go*](https://interpreterbook.com/) by *Thorsten Ball*.

[![Writing An Interpreter In Go](waiig.png)](https://interpreterbook.com/)

Examples:

```
>>> int x = 10;
10
>>> float y = 3.50;
3.50000
>>> string s = "a string";
a string
>>> char c = 'c';
c
>>> x;
10
>>> 1 + 1;
2
>>> s;
a string
>>> int add(int a, int b) { return a + b; }
int add(int a, int b) { return a + b; }
>>> add(3, 5);
8
>>> 
Ctrl + D to exit
```
