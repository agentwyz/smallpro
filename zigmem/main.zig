//this code is about memeory allocator

//the first example 
//1.Automatically handled by the complier
//2. very fast allocation
//3. very fast cleanup
//-----------------
//-toal memory fixed size fixed lifetimes
pub fn sumRange(strat: u32, end: u32) u32 {
    var result = 0;
    var curr = strat;

    while (curr <= end) : (curr += 1) {
        result += 1;
    }
    return result;
}

//simplest allocation strategy, 简化分配策略
//leverage the OS
const PageAllocator = struct {
    //request a memory from the operating system 
    //and in a POSIX system that's a memory map 
    pub fn alloc(self: *@This(), size: u32) []u8 {          //slow(syscall)
        const mem = std.os.mmap(
            alignForward(size, page_size)
        ) catch {
            return error.OutOfMemory;
        }
        return mem[0..size];
    }
    //upon freeing it we just call memory
    pub fn free(self: *@This(), mem: []u8) void {
        return std.os.munmap(mem);
    }
}

//bump allocator
//very fast allocation
//control lifetimes via buffer 控制声明周期通过buffer
//-Fixed total memory
//cannot free memory
const FixedBufferAlloactor = struct {
    buffer: []u8,
    end_idx: usize,

    pub fn alloc(self: *@This(), size: u32) []u32 {
        const new_idx = self.end_idx + size;
        if (new_idx > self.buffer.len) {
            return error.OutOfMemory;
        }
        //返回一片内存
        const res = self.buffer[self.end_idx..new_idx];
        self.end_idx = new_idx;
        return res;
    }
}



//bumap allocator with expandable memory
//Arena allocator
//expandable total memory
//manual lifetime
//cannot free individual memory, but we can free we request

//create a memory, create two memory, create three memory


//what about free()、general purpose
//free list
const FreeListAllocator = struct {
    root: ?*Node,

    fn find(self: *@This(), size: u32) ?[] {
        var iter = self.root;
        while (iter) |node| : (iter = node.next) {
            if (node.size == size) {
                self.remove(node);
                return node.buffer();
            }
        }
        return null;
    }

    pub fn free(self: *@This(), mem: u32) void {
        const node = Node.init(mem);
        //添加一个
        self.prepend(node);
    }
}









