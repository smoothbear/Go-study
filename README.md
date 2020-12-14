# Go-study

이 레포지토리는 Go 언어를 공부하면서 배운 내용을 정리한 레포지토리입니다.

Go란?

![Go-Character](https://1.cms.s81c.com/sites/default/files/2019-02-27/1_Gogophercolor-180.png)

2009년 구글에서 일하는 로버트 그리즈머, 롭 파이크, 켄 톰프슨이 개발한 프로그래밍 언어.

-----------

## Go의 OOP(Object Oriented Programming)

Go는 객체 없이도 객체 지향 프로그래밍의 원리를 사용할 수 있다.

Go는 전통적인 OOP 언어들이 객체를 지니는 반면, 값을 가진다.

<Java의 객체>

``` java
class SomeObject {
  private String str;
  
  SomeObject(String str) {
    this.str = str;
  }
}

SomeObject object1 = new SomeObject("Hello World!");
SomeObject object2 = object1;
```

위 코드에서는 object1과 object2가 모두 같은 SomeObject의 인스턴스를 가리키고 있다.

```go
import (
	"fmt"
)

type SomeStruct struct {
  Field string
}

func main() {
  value := SomeStruct(Field: "Hello World!")
  copy := value
  copy.Field = "Hello"
}
```

위 코드에서 copy.Field의 값을 재설정할때, value.Field는 값을 변경하지 않는 것을 볼 수 있다. 같은 인스턴스를 참고하고 싶을 땐, C와 유사하게 포인터를 사용한다.

--------

## 캡슐화 (Encapsulation)

전통적으로 class 기반의 OOP에서는 캡슐화가 private/public 같은 접근 제한자를 통하여 이루어졌다. Go에서는 패키지(package) 수준에서 캡슐화가 이루어졌다. 하지만 Go에서는 노출되지 않는 요소는 첫 문자를 소문자로 하여 표시하며, 해당 패키지 내에서만 접근이 가능하게 한다.

```go
package main

func Expose() {} // 패키지 밖으로 노출되는 함수
func hide() {} // 패키지 내부에서만 사용 가능한 함수
```

-----

## 다형성 (Polymorphism)

Go에서는 인터페이스(interface)가 암시적으로 충족되며, 인터페이스 또한 타입이다.

**인터페이스가 암시적으로 충족되며** -> 인터페이스의 모든 메서드가 어떤 타입의 메서드 집합에 모두 포함되어 있을 경우 해당 타입은 인터페이스를 만족한다. Go에는 *implements* 키워드가 없다.

**인터페이스는 또한 타입이다** -> 어떤 타입이 한 인터페이스를 만족하면, 그 타입은 또한 타입이 만족하는 인터페이스에 의해 제한되는 모든 타입을 만족한다.

-----



### 추가

- Go의 가비지 컬렉션 :  https://engineering.linecorp.com/ko/blog/go-gc/

