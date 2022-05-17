### find .class file

#### example
```java
package java;

public class HelloWorld{

    public static void main(String[] args) {
        System.out.print("Hello World");
    }

}
```
一个简单的 HelloWorld 类在加载前也需要加载其超类, 即<mark> java.lang.Object </mark>.  
除此之外,包括打印输出的 System 类也需要进行加载等等...  

#### 类路径
Java 虚拟机并没有规定如何去寻找类,不同的虚拟机可能有不同的实现.  
按照类路径搜索的先后顺序,可以分为一下:  
- 启动类路径(bootstrap classpath)
- 拓展类路径(extension classpath)
- 用户类路径(user classpath)

启动类路径默认对应jre\lib目录，Java标准库（大部分在rt.jar里）位于该路径。  
扩展类路径默认对应jre\lib\ext目录，使用Java扩展机制的类位于这个路径。  
我们自己实现的类，以及第三方类库则位于用户类路径。  
可以通过-Xbootclasspath选项修改启动类路径，不过通常并不需要这样做
用户类路径的默认值是当前目录，也就是“.”。可以设置CLASSPATH环境变量来修改用户类路径，但是这样做不够灵活，所
以不推荐使用。更好的办法是给java命令传递-classpath（或简写为-cp）选项。-classpath/-cp选项的优先级更高，可以覆盖CLASSPATH环境变量设置.  
-classpath/-cp选项既可以指定目录，也可以指定JAR文件或者ZIP文件.还可以同时指定多个目录或文件，用分隔符分开即可。分隔符因操作系统而异。在Windows系统下是分号，在类UNIX（包括Linux、Mac OS X等）系统下是冒号。例如在Windows下:  
> java -cp path\to\classes;lib\a.jar;lib\b.jar;lib\c.zip ...  

从Java 6开始，还可以使用通配符（*）指定某个目录下的所有JAR文件，格式如下：
> java -cp classes;lib\* ...  


