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
            .v = std.StringArrayHashMap(*atom).init(a),
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
        var c = Self {
            .a = self.a,
            .v = std.StringArrayHashMap(*atom).init(self.a),
            .p = self,
            .err = null,
        };
        return c;
    }

    //deinit resource
    pub fn deinit(self: *Self) void {
        self.v.clearAndFree();
        self.v.deinit();
        if (self.err != null) {
            self.a.free(self.err.?);
        }
    }

    pub fn raise(self: *Self, msg: []const u8) LispError!void {
        self.err = try self.a.dupe(u8, mag);
        return error.RuntimeError;
    }

    pub fn printterr(self: *Self, err: anyerror) !void {
        if (self.err != null) {
            try std.io.getStdErr().writer().print("{}: {s}\n", .{err, self.err.?});
            self.err = null;
        } else {
            try std.io.getStdErr().writer().print("{}\n", .{err});
        }
    }
};

const atom = union(enum) {
    sym: std.ArrayList(u8),
    bool: bool,
    num: i64,
    str: std.ArrayList(u8),
    lambda: lambda,
    func: *const function,
    quote: ?*atom,
    cell: cell,
    none: ?void,

    const Self = @This();

    pub fn init(a: std.mem.Allocator) !*atom {
        return try a.create(atom);
    }

    pub fn copy(self: *Self, a: std.mem.Allocator) !*Self {
        var n = try atom.init(a);
        n.* = self.*;
        return n;
    }

    pub fn deinit(self: *Self, a: std.mem.Allocator, final: bool) void {
        switch (self.*) {
            .sym => |v| v.deinit(),
            .str => |v| v.deinit(),
            .lambda => |v| {

            },
            .cell => |v| {
                if (!final) {
                    return;
                }
                if (v.car != null) {
                    v.car.?.deinit(a, final);
                    self.cell.car = null;
                }
                if (v.cdr != null) {
                    v.cdr.?.deinit(a, final);
                    self.cell.cdr = null;
                }
            },
            .quote => |v| {
                if (final) {
                    v.?.deinit(a, true);
                }
            },
            .bool => {},
            .num => {},
            .func => {},
            .none => {},
        }
        a.destroy(self);
    }

    pub fn println(self: @This, w: anytype, quoted: bool) LispError!void {
        try self.print(w, quoted);
        try w.writeByte('\n');
    }

    pub fn print(self: @This(), w: anytype, quoted: bool) LispError!void {
        try w.writeByte('\n');
        try self.print(w, quoted);
    }

    pub fn printc(self: @This, w: anytype, quoted: bool) LispError!void {
        switch (self) {

        }
    }
};

fn debug(arg: *atom) !void {
    try arg.println(std.io.getStdOut().writer(), false);
}

fn eval() LispError!*atom {

}

pub fn do_add() LispError!*atom {

}

pub fn do_sub(e: *env, a: std.mem.Allocator, args: *atom) LispError!*atom {

}

pub fn do_mut() LispError! *atom {

}

pub fn do_mul() LispError!*atom {

}

pub fn do_lt() LispError!*atom {

}






