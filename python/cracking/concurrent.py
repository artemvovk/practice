"""Multithreaded and Multiprocessed"""
import queue
import threading


def fb_3s(work):
    while True:
        num = work.get()
        work.task_done()
        if num is None:
            return
        if not num%3 and num%5:
            print("Fizz")
            continue
        work.put(num)

def fb_5s(work):
    while True:
        num = work.get()
        work.task_done()
        if num is None:
            return
        if not num%5 and num%3:
            print("Buzz")
            continue
        work.put(num)

def fb_15s(work):
    while True:
        num = work.get()
        work.task_done()
        if num is None:
            return
        if not num%15:
            print("FizzBuzz")
            continue
        work.put(num)

def fb_rest(work):
    while True:
        num = work.get()
        work.task_done()
        if num is None:
            return
        if num%5 and num%3:
            print(num)
            continue
        work.put(num)

def mt_fizzbuzz(num):
    work = queue.Queue()
    threads = []
    threads.append(threading.Thread(target=fb_3s, args=(work,)))
    threads.append(threading.Thread(target=fb_5s, args=(work,)))
    threads.append(threading.Thread(target=fb_15s, args=(work,)))
    threads.append(threading.Thread(target=fb_rest, args=(work,)))
    for thread in threads:
        thread.start()
    for idx in range(num):
        work.put(idx)
    work.join()
    for thread in threads:
        work.put(None)
    for thread in threads:
        thread.join()
    return num
