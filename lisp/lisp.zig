const std = @import("std");

const cell = struct {
    car: ?*atom,
    cdr: ?*atom,
};

const lambda = struct {
    e: ?*env,
    cell: cell,
};

const ref = ?*atom;

const function = struct {
    name: []const u8,
    //function sign
    ptr: *const fn (*env, std.mem.Allocator, *atom) LispError ! *atom,
};

const env = struct {
    a: std.mem.Allocator,
    v: std.StringArrayHashMap(*atom),
    p: ?*env,
    err: ?[]const u8,

    const Self = @This();
    pub fn init(a: std.mem.Allocator) Self {
        return Self {
            .a = a,
            .v = std.
            .p = null,
            .err = null
        };
    }

    pub fn get(self: *Self, key: []const u8) !?*atom {
        var e: *env = self;
        while (true) {
            if (e.v.get(key)) |ev| {
                return ev;
            }

            if (e.p == null) {
                break;
            }
            e = e.p.?;
        }
        try e.raise("invalid symbol");
        unreachable;
    }

    pub fn child(self: *Self) Self {
        
    }
}

