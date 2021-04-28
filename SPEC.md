# syntax

- variable declarations begin with the keyword DECADE
- - variables are statically typed, but we just don't care
- - names are \`name name\`
- - there are no types, only sizes.
- - structs are type RAVIS and their value is a function pointer
- - - accessing struct members uses keyword RAVIS
- function declarations begin with the keyword JAM
- - returns with [MOKKA value]
- - if no return, it'll leak into other functions
- built in functions should begin with the ! symbol
- all statements are enclosed in `[]` similarly to lisp
- polish notation is used
- everything is done as a function call. Making a variable, getting a list field etc.

- [DECADE \`name\` type value] is variable
- [JAM name [ params ] [ body ] type] is function
- [MOKKA value] is return
- [DECADE name RAVIS funcptr] is struct
- [JAMUNTIL [DECADE name type value] [JAMLENGTH 0 x]] is for loop
- [GAINTIME value1 value2] is addition
- [LOSETIME value1 value2] is subtraction
- [CRUNCHTIME value1 value2] is power
- [LAZYTIME value1 value2] is root
- [OPATTR value1 value2] is XOR
- [SIMATTR value1 value2] is AND
- [BUILDUP value1 value2] is OR
- [KEEPJAMMING condition] is while loop
- [PUBGAME variable value] is set value
- [PLAYGAME variable] is get value
- [FINISHED condition] is if statement, skip next instruction if not condition
- [VOTEOVER value1 value2] is greater than
- [JAMOVER value1 value2] is equal to
- [JAMRUNNING value1 value2] is less than
- [NOJAM value] is not
- builtin functions are [!print text], [!timeleft]
- 

# technical details

- arrays are of dynamic length and should be implemented using lists
- imports are handled the same way as in c - just copying the content of the files. Function prototypes are made automatically.
- interpreter/complier should check to see if the Decade Jam is over or not and if it is have it fail to compile/interpret
- variables are stored on the DECADE-- stack, which is on the heap
- the DECADE-- stack is of undefined size


# Standard Library
- random function should begin with !roll followed by d(any number of dice) example !roll d1 will roll a 1
- some sort of GUI library
- networking library
