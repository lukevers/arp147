# math

**This was *not* written by the Engo project. The original author decided to remove the original project from GitHub, we were given permission to reupload the code as long as we did not credit them.**


This library is a copy of the standard math library but for float32 (more common in computer graphics and games. Most of the function are simply casting the arguments as `float64` and forwarding to the original math library (something you would have to do anyway, but slowly we are translating the original function to the `float32` version.  
ALL tests are passing, meaning whatever program you we're running using the 64 bit version should still be correct using this library (using the IEEE floating point definition of "correct").  
I only care about i386 and amd64, but if I receive pull request for other architecture I will merge them (I can't test it so please send me proof that all test pass and benchmark are better :)

## rewriten functions
The benchmark is based on my macbook pro, with an i386 processor but a x86_64 operating system. benchmark may produce different results on different computers. However if you see an architecture for which casting to float64, calling the standard lib then casting back to flaot32 is faster then my implementation please tell me.  
✓ = Test passing.  
✗ = Not implemented yet.  
? = Implemented but not tested.  
\* = has no ASM implementation.


|Abs|status|lux math|casting|std math|accel|  
|---|---|---:|---:|---:|---:|  
|function name|✓, ✗, ? or \* |benchmark for float32 using lux math package|benchmark for casting a var to float64, doing the op then casting back|benchmark for float64 using standard math package|accelleration between casting and lux math|

|Abs|status|lux math|casting|std math|accel|  
|---|---|---:|---:|---:|---:|
|software|✓|||||
|amd64|✓|2.10 ns/op|2.37 ns/op|2.08 ns/op|+13%|
|386|✓|2.02 ns/op|2.24 ns/op|2.02 ns/op|+11%|

|Sqrt|status|lux math|casting|std math|accel|
|---|---|---:|---:|---:|---:|
|software|✓|||||
|amd64|✓|2.33 ns/op|4.70 ns/op|0.33 ns/op|+102%|
|386|✓|2.33 ns/op|5.14 ns/op|4.74 ns/op|+121%|

|Acosh|status|lux math|casting|std math|accel|  
|---|---|---:|---:|---:|---:|  
|software|✓|43.5 ns/op|75.8 ns/op|40.3 ns/op|+74%|

|Asinh|status|lux math|casting|std math|accel|  
|---|---|---:|---:|---:|---:|  
|software|✓|47.3 ns/op|81.3 ns/op|46.1 ns/op|+72%|

|Gamma|status|lux math|casting|std math|accel|  
|---|---|---:|---:|---:|---:|  
|software|✓|25.1 ns/op|38.5 ns/op|24.0 ns/op|+53%|


|template|status|lux math|casting|std math|accel|  
|---|---|---:|---:|---:|---:|  
|software|✓| ns/op| ns/op| ns/op|%|
|amd64|✓| ns/op| ns/op| ns/op|%|
|i386|✓| ns/op| ns/op| ns/op|%|


