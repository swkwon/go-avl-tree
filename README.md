# go-avl-tree
go-avl-tree는 goroutine으로 부터 안전한 generic tree 자료구조 입니다.
# 시작하기
```
go get github.com/swkwon/go-avl-tree@latest
```
```
    myTree := New[int, string]()
```
`New`함수로 tree객체를 생성 합니다. 위 예제는 경우 key의 자료형은 int, value의 자료형은 string입니다.
`key`의 자료형은 크기비교가 가능한 타입은 모두 사용할 수 있습니다. `key`의 자료형은 `cmp.Ordered` 타입은 모두 가능합니다.
참고로 `cmp.Ordered` 인터페이스는 아래와 같이 정의되어 있습니다.
```
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}
```
## 삽입
```
    myTree.Put(1, "hello")
```
이미 존재하는 key라면 value는 덮어씌워 집니다.
## 가져오기
```
    resultList := myTree.Gets(1, 100)
```
트리에서 key를 찾을 경우 파라메터로 여러개의 key를 입력할 수 있습니다. 결과는 결과 구조체의 slice로 리턴하게 되는데, 결과구조체는 key, value, 찾은 여부의 flag가 있습니다.
```
type GetResult[K cmp.Ordered, V any] struct {
	Key     K
	IsExist bool
	Value   V
}
```
## 삭제
```
	myTree.Delete(1, 10)
```
삭제의 경우도 여러개의 key를 파라메터로 입력할 수 있습니다. 노드를 삭제할 때마다 tree의 밸런스를 다시 맞춥니다.