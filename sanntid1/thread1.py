from threading import Thread, Lock

i = 0
lock = Lock()

def threadOne():
	global i	
	for x in range(1000000):
		lock.acquire()
		i = i + 1
		lock.release()

def threadTwo():
	global i	
	for y in range(1000000):
		lock.acquire()
		i = i - 1
		lock.release()

def main():
	global i

	thread_1 = Thread(target = threadOne, args = (),)
	thread_2 = Thread(target = threadTwo, args = (),)

	thread_1.start()
	thread_2.start()

	thread_1.join()
	thread_2.join()
	
	print i

main()
