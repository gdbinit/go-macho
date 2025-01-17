enum MyEnum: String {
    case A = "test"
    case B
    case C
}

struct Name {
    var a: MyEnum
    var b: MyEnum
}

var name = Name(a: .A, b: .B)

func greet(person: String) -> String {
    let greeting = "Hello, " + person + "!"
    return greeting
}