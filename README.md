# 设计模式

## 运行

```
go test
```

## 理解

本项目用 Go 语言实现了 20 个 GoF 设计模式，另有 Singleton, Prototype, Memento 三个模式被排除在外。

## 分类

GoF 一书根据设计模式的目的和范围两个原则进行分类，分成三类：创建型、结构型、行为型。如下表所示。

| | Creational(3) | Structural(7) | Behavioral(10) |
|:-|:-|:-|:-|
| Class | Factory-Method | Adapter | Interpreter |
| Object | Abstract-Factory<br>Builder | Adapter<br>Bridge<br>Composite<br>Decorator<br>Facade<br>Flyweight<br>Proxy | Chain<br>Command<br>Iterator<br>Mediator<br>Observer<br>State<br>Strategy<br>Visitor |

## 扩展

从扩展的角度看，变化的需求可以通过改变不同的子类或联合类实现，当然也可以同时改变子类和联合类实现。在面对需求的变化时，我们有时需要扩展对象的功能，有时需要扩展对象的数量，又或者需要合理分派各种不同的功能实现。那么同时考虑扩展目的和扩展方式，我们可以对这些设计模式进行如下分类。

| | Dispatch(6) | Object(10) | Group(4) |
|:-|:-|:-|:-|
| Interface | Facade<br>Flyweight | Template-Method<br>Abstract-Factory<br>Factory-Method | Composite<br>Iterator |
| Association | State<br>Strategy<br>Interpreter | Adapter<br>Proxy<br>Decorator<br>Chain | Mediator |
| I + A + I | Command<br>Observer | Builder<br>Bridge | Visitor |

#### 分派

Dispatch 类的模式需要 broker 协调实现需求的各种子类。子类之间通过不同的配合方式实现功能：

0. Strategy 模式里的每个请求都由某个子类完全实现
0. Interpreter 模式里的每个请求是由子类之前相互调用实现
0. State 模式里是由各个请求驱动子类的切换完成功能

Command 模式不光可以扩展各种实现功能的子类，还可以扩展 broker 的功能。Observer 模式可以理解为一种间接的协调。

#### 对象

模板方法代表通过子类扩展功能的基本思路，基类最大程度的实现代码复用。Bridge 将两个模板方法合在一起，前者扩展设计，后者扩展实现。Proxy 和 Decorator 模式则是通过关联方式实现对象的扩展。

#### 集合

Composite 解决一个集合里的对象组合问题，Iterator 模式解决集合的遍历问题，Mediator 模式通过一个 Director 集成了相关对象。Visitor 解决集合对象的访问问题，并支持各种不同访问功能的扩展。