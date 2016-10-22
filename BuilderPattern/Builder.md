## BuilderPattern

Separate the constrcution of complex object from its representation so that the same constrcution process can create different representations.

(将一个复杂对象的构建与它的表示分离.使得同样的构建过程可以创建不同的表示)

### Roles

* 导演类: 负责安排已有模块,由Builder开始建造.e.g. Director

* 产品类: 实现模板方法模式,即模板方法和基本方法.e.g. BenzModel

* 抽象建造者:规范产品的组建,一般为父类/interface,由子类去实现。 e.g. CarBuilder

* 具体建造者:实现抽象类定义的所有方法,并且返回一个组建好的对象。e.g. BenzBuilder
