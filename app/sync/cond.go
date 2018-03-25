package sync


/**
条件变量相关，条件变量总要与互斥量组合使用
sync.Cond类型的方法集合：wait、singal、broadcast，wait方法会自动地对与该条件变量关联的那个锁进行解锁，并且使它所在的goroutine阻塞
 */


