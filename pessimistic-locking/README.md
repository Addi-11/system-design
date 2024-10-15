## Pessimitic Locking in GO
Running the following code on threads will give you inconsistent results.
```
func incCount(){
	count++;
	wg.Done()
}

func doCount(){
	for i:=0; i < 1000000; i++{
		wg.Add(1)
		go incCount()
	}
}
```
**Reason**: <br>
Count++ (increment operation) is not atomic in nature. Micro-operations in registers
1. movl %rdi, %tmp  LOAD
2. addl $1, %tmp    ADD
3. movl %tmp, %rdi  STORE

**Solution**: <br>
Lock and Unlock the operation (count++) using Mutex, so that only one thread runs the operation at a time.
```
var mu sync.Mutex

mu.Lock()
count++;
mu.Unlock()
```
#### Reference:
[Pessimitic Locking with Mutex](https://www.youtube.com/watch?v=4F-WiPFrPsA)