mutex是互斥锁, 即被锁住的临界区无法被除当前线程外的其它线程所访问

在下面的示例里, *死锁不一定会发生*

|Thread1    |  Thread2
------------|----------
|A.lock()   |  B.lock()
|B.lock()   |  A.lock()
|B.unlock() |  A.unlock()
|A.unlock() |  B.unlock()

如果Thread1先执行了A.lock(),然后Thread2执行了B.lock(),这个时候就发生了死锁。 代码mutex.go里的例子就说明了这个问题。

