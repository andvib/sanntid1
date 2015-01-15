from threading import Thread

i = 0

def threadOne():
	global i	
	for x in range(1000000):
		i = i + 1

def threadTwo():
	global i	
	for y in range(1000000):
		i = i - 1

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
