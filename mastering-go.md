# 前言

```
go mod init filename // 初始化项目
```

```
go mod tidy // 加载依赖
```

# 1 入门

Go的标准库提供了100多个包，以支持常见功能，如输入、输出、排序以及文本处理

## 1.1 命令行工具

var声明定义了变量。变量会在声明时直接初始化，如果变量没有显式初始化，则被隐式地赋予其类型的*零值*

Go语言只有for循环这一种循环语句，for有三种形式

```
for initialization; condition; post { 
    // zero or more statements
}
```

```
for condition { // while loop
    // ...
}
```

```
for { // loop
    // ...
}
```

## 1.2 查找重复行

`Printf`的转换

```
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数。
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true或false
%c          字符（rune） (Unicode码点)
%s          字符串
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
```

