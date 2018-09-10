https://blog.csdn.net/q_l_s/article/details/52369065

饿汉：

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

懒汉：

	class CSingleton    
	{    
	private:    
	    CSingleton()      
	    {    
	    }    
	public:    
	    static CSingleton * GetInstance()    
	    {    
		static CSingleton instance;     
		return &instance;    
	    }    
	};    
