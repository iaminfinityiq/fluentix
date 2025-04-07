**Here are the syntaxes that we have right now**<br><br>
**1. Data types**
There are 3 data types available: `int`, `double` and `bool` (`int` and `double` are in range of 64 bits)
You can do operations with `numbers`, but not `booleans`
Here is an example of an `int`: -5
Here is an example of a `double`: 3.14

**2. Variables**
Variables are kind of like a `container`because each of them store a value.
To assign a value to a variable, we use:
```
<var_name> = <value>
```
This syntax also works when updating the values.

To get the value of the variable, we simply write the variable name:
```
<var_name>
```

We can also do operation between variables, like:
```
<var1> + <var2>
```

To name a variable, the first letter of a variable must me a formal English letter, from A through Z or from a through z, or an underscore. Other characters of the variable must be letters, underscores, or digits.

Here is a full program involving variables:
```
a = 5
b = 6

a / b
```

**3. Special operators**
There are many different kinds of operators we got in Fluentix. I will list all of them just in case:
- `+`: Addition
- `-`: Subtraction
- `*`: Multiplication
- `/`: Division, throws a `MathError` if you try to divide a number with `0`
- `%`: Modulus operator, throws a `MathError` if you try to mod a number with `0`. Returns the remainder of dividing a number to another number
- `==`: Compares if 2 objects are equal to each other...
- `!=`: Compares if 2 objects are different to each other...
- `>`: Compares if the first number is greater than the second number
- `<`: Compares if the first number is smaller than the second number
- `>=`: Compares if the first number is greater than equal the second number
- `<=`: Compares if the first number is smaller than or equal to the second number
- `|x|`: Absolute value of a number, the value would be in between of the two pipes (`|`)
- `x!`: Factorial of a number, the value would stand before the exclamation mark(s) (`!`). Read more at https://en.wikipedia.org/wiki/Factorial and https://mathworld.wolfram.com/Multifactorial.html. Only works for non-negative integers. `0! = 1`

The order of operation is performed using `PEMDAS` or `BODMAS`, which follows the basic rules of operation. You can include left parentheses and right parentheses to change the precedence of some operation (`(`, `)`).
We also got some special operator that you might not have seen in other language (by syntax), which is the absolute value operator and factorial operator

**4. Line seperations**
A Fluentix code consists of many lines, or probably none, but there are some things to note:
1. Each line could end in a semicolon (`;`) or not.
2. There can't be multiple statements on a line, expressions are no exception (except for some special cases)

**5. If statements**
An if statement is denoted by this structure. These syntaxes differ through file extensions.
- For `.flu` file extensions, we got:
```
if [condition] [->/do]
[tab][your code, each line of code in the if statement must be seperated by tabs]
[unless/elseif/elif/else if] [condition] [->/do] (these can have an optional amount of it)
[same for if part]
else [->/do] (these are optional, but there can only be 0 or 1 of it)
[same for if part]
```
- For `.fl` file extensions, we got:
```
if [condition] [->/do] {
[your code, this time, no need for tabs]
} [unless/elseif/elif/else if] [condition] [->/do] (these can have an optional amount of it) {
[same for if part]
} else [->/do] {
[same for if part]
}
```

Note that you can switch the style of `.fl if statements` as so, so that it may suit you better:
```
if [condition] [->/do]
{
[your code]
}
[unless/elseif/elif/else if] [condition] [->/do] (these can have an optional amount of it]
{
[your code]
}
else [->/do]
{
[your code]
}
```
