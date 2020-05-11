package LruCache_test

import (
    "testing"
)

func TestLruCache(t *testing.T) {
    lru := Constructor(3)
    lru.Put(1,1)
    lru.Put(2,2)
    lru.Put(3,3)
    lru.Put(4,4)
    lru.Get(4)
    lru.Get(3)
    lru.Get(2)
    lru.Get(1)
    lru.Put(5,5)
    lru.Get(1)
    lru.Get(2)
    lru.Get(3)
    lru.Get(4)
    lru.Get(5)
    t.Log(lru.m)
}


/**
 * https://leetcode-cn.com/problems/lru-cache/
 * 采用双链表的思路实现O(1)
 */
type LRUCache struct {
    len     int
    count   int
    head    *Cache
    tail    *Cache
    m       map[int]*Cache
}


type Cache struct {
    head    *Cache
    tail    *Cache
    value   int
    key     int
}


func Constructor(capacity int) LRUCache {

    lru := LRUCache{
        len:capacity,
        count:0,
        head:nil,
        tail:nil,
        m:make(map[int]*Cache, capacity),
    }
    
    return lru
}


func (this *LRUCache) Get(key int) int {
    v, ok := this.m[key]
    if ok  {

        if this.head == v {
            return v.value
        }
        if this.tail == v {
            this.tail = v.head
            this.tail.tail = nil
        } else {
            v.tail.head = v.head
            v.head.tail = v.tail
        }
        v.tail = this.head
        v.head = nil
        this.head.head=v
        this.head = v

        return v.value
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    v, ok := this.m[key]
    if ok {
        v.value = value
        if this.head == v {
            return
        }

        if this.tail == v {
            this.tail = v.head
            this.tail.tail = nil
        } else {
            v.tail.head = v.head
            v.head.tail = v.tail
        }
        v.tail = this.head
        v.head = nil
        this.head.head=v
        this.head = v
    } else {
        if this.count < this.len {
            cache := &Cache{
                head:nil,
                tail:nil,
                value:value,
                key:key,
            }
            if this.count == 0 {
                this.tail = cache
            } else {
                this.head.head = cache
                cache.tail = this.head
            }
            this.head = cache
            this.count++
            this.m[key] = cache
        } else {
            delete(this.m, this.tail.key)

            cache := this.tail
            cache.key = key
            cache.value = value
            this.m[key] = cache

            if this.len > 1 {
                this.tail = this.tail.head
                this.tail.tail = nil
                this.head.head = cache
                cache.tail = this.head
                this.head = cache
                this.head.head = nil
            }

        }
    }
}