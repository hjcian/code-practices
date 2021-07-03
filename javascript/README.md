# javascript syntax notes

## Public/Private instance/class fields/methods
> [Public class fields](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Classes/Public_class_fields)
> [Private class features](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Classes/Private_class_fields)

- class level 的 static field and method 只能被 `Class` 直接呼叫；instance 無法呼叫/存取 static methods/fields
- private field/method 以 `#` 前綴，並且只能在 class 內被使用