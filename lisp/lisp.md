## allocator

```rust
fn init(a: std.mem.Allocator) !*atom {
    return try a.create(atom); //创建一个对象
}
```

when linking against libc, zig exposes this allocator with std.