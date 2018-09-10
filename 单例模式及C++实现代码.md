https://www.cnblogs.com/cxjchen/p/3148582.html

	这样就可以了，保证只取得了一个实例。但是在多线程的环境下却不行了，因为很可能两个线程同时运行到if (instance == NULL)这一句，导致可能会产生两个实例。于是就要在代码中加锁。
	Singleton* getInstance()
	{
	    lock();
	    if (instance == NULL)
	    {
	       instance = new Singleton();
	    }
	    unlock();

	    return instance;
	}
	但这样写的话，会稍稍映像性能，因为每次判断是否为空都需要被锁定，如果有很多线程的话，就爱会造成大量线程的阻塞。于是大神们又想出了双重锁定。
	Singleton* getInstance()
	{
	    if (instance == NULL)
	    {
		lock();
		if (instance == NULL)
		{
			instance = new Singleton();
		}
		unlock();
	    }

	    return instance;
	}
	这样只够极低的几率下，通过越过了if (instance == NULL)的线程才会有进入锁定临界区的可能性，这种几率还是比较低的，不会阻塞太多的线程，但为了防止一个线程进入临界区创建实例，另外的线程也进去临界区创建实例，又加上了一道防御if (instance == NULL)，这样就确保不会重复创建了。


	class Singleton{
	public:
		static Singleton* getInstance();

	private:
		Singleton();
		//把复制构造函数和=操作符也设为私有,防止被复制
		Singleton(const Singleton&);
		Singleton& operator=(const Singleton&);

		static Singleton* instance;
	};

