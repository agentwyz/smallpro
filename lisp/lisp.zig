const std = @import("std");

const cell = struct {
    //lisp语言的两个基本操作
    car: ?*atom,
    cdr: ?*atom,
};

const lambda = struct {
    //lambda是函数表达式, 函数表达式
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
    cell: cell, //这个cell会
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
                if (!final) {
                    return;
                }
                if (v.cell.car != null) {
                    v.cell.car.?.deinit(a, final);
                    self.cell.car = null;
                }
                if (v.cell.cdr != null) {
                    v.cell.cdr.?.deinit(a, final);
                    self.cell.cdr = null;
                }
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
    
    //这个函数向命令行打印
    pub fn printc(self: @This, w: anytype, quoted: bool) LispError!void {
        switch (self) {
            .none => try w.writeAll("null"),
            .sym => |v| try w.writeAll(v.items),
            .str => |v| {
                if (quoted) {
                    try w.writeByte('"');
                    for (v.items) |c| {
                        switch (c) {
                            '\\' => try w.writeAll("\\\\"),
                            '"' => try w.writeAll("\\\""),
                            '\n' => try w.writeAll("\\n"),
                            '\r' => try w.writeAll("\\r"),
                            else => try w.writeByte(c),
                        }
                    }
                    //尝试输出有右边的"
                    try w.writeByte('"');
                } else {
                    try w.writeAll(v.items);
                }
            },
            .func => |v| try w.writeByte(v.name),
            .bool => |v| {
                if (v) {
                    try w.writeAll("T");
                } else {
                    try w.write("nil");
                }
            },
            .num => |v| try w.print("{}", .{v}),
            .lambda => |v| {
                try w.writeAll("(lambda");
                try v.cell.cdr.?.cell.car.?.cell.cdr.?.princ(w, quoted);
                try w.writeByte(' ');
                try v.cell.cdr.?.cell.car.?.princ(w, quoted);
                try w.writeByte(')');
            },
            .cell => |v| {
                try w.writeByte('(');
                try v.car.?.princ(w, false);
                try w.writeByte(' ');
                if (v.cdr == null) {
                    return;
                }
                var a = v.cdr;
                while (a != null) {
                    if (a.?.cell.car == null) {
                        break;
                    }
                    try a.?.cell.car.?.princ(w, quoted);
                    if (a.?.cell.cdr == null) {
                        break;
                    }
                    a = a.?.cell.cdr;
                    if (a == null) {
                        break;
                    }
                    try w.writeByte(' ');
                }
                try w.writeByte(')');
            },
            .qoute => |v| {
                try w.writeByte('\x27');
                try v.?.princ(w, quoted);
            },
        }
    }
};

fn debug(arg: *atom) !void {
    try arg.println(std.io.getStdOut().writer(), false);
}

//计算参数
fn eval(e: *env, a: std.mem.Allocator, root: *atom) LispError!*atom {
    var arg: ?*atom = root;

    return switch (arg.?.*) {
        atom.sym => |v| blk: {
            var p = e;
            while (true) {
                if (p.v.get(v.items)) |ev| {
                    break :blk try eval(e, a, ev);
                }
                if (p.p == null) {
                    break;
                }
                p = p.p.?;
            }
            try e.raise("invalid symbol");
        },
        
        atom.str => |v| blk: {
            var bytes = std.ArrayList(u8).init(a);
            try bytes.writer().writeAll(v.items);
            var na = try atom.init(a);
            na.* = atom{
                .str = bytes,
            };
            break :blk na;
        },

        atom.lambda => try arg.?.copy(a),

        atom.cell => blk: {
            var last = arg.?;
            while (true) {
                last = try switch(arg.?.cell.car.?.*) {
                    atom.lambda => {

                    },
                    atom.sym => {

                    }
                }
            }
        }
    }    
}

pub fn do_add(e: *env, a: std.mem.Allocator, args: *atom) LispError!*atom {
    //计算加法
    var arg = args;
    var num: i64 = 0;

    while (true) {
        var val = try eval(e, a, arg.cell.car.?);
        defer val.deinit(a, false);
        if (val.* == atom.num) {
            num += val.num;
        } else {
            try e.raise("invalid type for +");
        }
        if (arg.cell.cdr == null) {
            var na = try atom.init(a);
            na.* = atom{
                .num = num,
            };
            return na;
        }
        arg = arg.cell.cdr.?;
    }
    unreachable;
}

pub fn do_sub(e: *env, a: std.mem.Allocator, args: *atom) LispError!*atom {
    var arg = args;
    var val = try eval(e, a, arg.cell.car.?);
    //错误处理
    if (val.* != atom.num) {
        try e.raise("invalid type for");
    }
    var num: i64 = val.num;
    val.deinit(a, false);
    
    if (arg.cell.cdr == null) 
    {
        var na = try atom.init(a);
        na.* = atom{
            .num = num,
        };
        return na;
    }

    while (true) {
        arg = arg.cell.cdr.?;
        val = try eval(e, a, arg.cell.car.?);
        defer val.deinit(a, false);
        if (val.* == atom.num) {
            num -= val.num;
        } else {
            try a.raise("invalid type for -");
        }

        if (arg.cell.cdr == null) {
            var na = try atom.init(a);
            na.* = atom {
                .num = num,
            };
            return na;
        }
    }
    unreachable;
}

pub fn do_mat(e: *env, a: std.mem.Allocator, args: *atom) LispError! *atom {
    var arg = args;
    var num: i64 = 1;

    while (true) {
        var val = try eval(e, a, arg.cell.car.?);
        defer val.deinit(a, false);

        if (val.* == atom.num) {
            num *= val.num;
        } else {
            try e.raise("invalid type for *");
        }

        if (arg.cell.cdr == null) {
            var na = try atom.init(a);
            na.* = atom{
                .num = num,
            };
            return na;
        }
        arg = arg.cell.cdr.?;
    }
    unreachable;
}

pub fn do_mul(e: *env, a: std.mem.Allocator, args: *atom) LispError!*atom {
    var arg = args;
    var val = try eval(e, a, arg.cell.car.?);
    defer val.deint(a, false);

    if (val.* != atom.num) {
        try e.raise("invalid type for/");
    }

    var num: i64 = val.num;

    if (arg.cell.cdr == null) {
        var na = try atom.init(a);
        na.* = atom {
            .num = num,
        };
        return na;
    }
    
    while (true)
    {
        arg = arg.cell.cdr.?;
        val = try eval(e, a, arg.cell.car.?);
        if (val.* == atom.num) {
            num = @divTrunc(num, val.num);
        } else {
            try e.raise("invaild type for/");
        }
        val.deinit(a, false);
        if (arg.cell.cdr == null) {
            var na = try atom.init(a);
            na.* = atom {
                .num = num,
            };
            return na;
        }
    }
    unreachable;
}

pub fn do_lt(e: *env, a: std.mem.Allocator, args: *atom) LispError!*atom {
    var arg = args;
    var lhs = try eval(e, a, arg.cell.car.?);

    defer lhs.deinit(a, false);

    if (lhs.* != atom.num) 
    {
        try e.raise("invalid type for <");
    }
    arg = arg.cell.cdr.?;
    var rhs = try eval(e, a, arg.cell.car.?);
    defer rhs.deinit(a, false);

    if (rhs.* != atom.num)
    {
        try e.raise("invalid type for <");
    }

    var na = try atom.init(a);
    na.* = atom{
        .bool = lhs.num < rhs.num,
    };
    return na;
}

pub fn do_gt() LispError!*atom {
    var arg = args;
    var lhs = try eval(e, a, arg.cell.car.?);
    defer lhs.deinit(a, false);

    if (lhs.* != atom.num)
    {
        try e.raise("invalid type for >");
    }

    arg = arg.cell.cdr.?;
    var rhs = try eval(e, a, arg.cell.car.?);
    defer rhs.deinit(a, false);

    if (rhs.* != atom.num)
    {
        try e.raise("invalid type for >");
    }

    var na = try atom.init(a);
    na.* = atom{
        .bool = lhs.num > rhs.num,
    };

    return na;
}

pub fn do_eq(e: *env, a: std.mem.Allocator, args: *atom) LispError!*atom {
    var arg = args;
    var lhs = try eval(e, a, arg.cell.car.?);
    defer lhs.deinit(a, false);

    if (lhs.* != atom.num)
    {
        try e.raise("invalid type for =");
    }
    arg = arg.cell.cdr.?;
    var rhs = try eval(e, a, arg.cell.car.?);
    defer rhs.deinit();

    if (rhs.* != atom.num)
    {
        try e.raise("invalid type for =");
    }

    var na = try atom.init(a);
    
    na.* = atom{
        .bool = lhs.num == rhs.num,
    }

    return na;
}

pub fn do_mod(e: *env, a: std.mem.Allocator, args: *atom) LispError!*atom
{
    var arg = args;
    var lhs = try eval()
}

pub fn do_if() LispError!*atom 
{

} 

pub fn do_seteq() LispError!*atom
{

}

pub fn do_printc() LispError!*atom {

}

pub fn do_print() LispError!*atom {

}

var builtins = [_] function {
        .{ .name = "+", .ptr = &do_add },
    .{ .name = "-", .ptr = &do_sub },
    .{ .name = "*", .ptr = &do_mat },
    .{ .name = "/", .ptr = &do_mul },
    .{ .name = "<", .ptr = &do_lt },
    .{ .name = ">", .ptr = &do_gt },
    .{ .name = "=", .ptr = &do_eq },
    .{ .name = "mod", .ptr = &do_mod },
    .{ .name = "cond", .ptr = &do_cond },
    .{ .name = "dotimes", .ptr = &do_dotimes },
    .{ .name = "if", .ptr = &do_if },
    .{ .name = "princ", .ptr = &do_princ },
    .{ .name = "print", .ptr = &do_print },
    .{ .name = "setq", .ptr = &do_setq },
    .{ .name = "defun", .ptr = &do_defun },
    .{ .name = "length", .ptr = &do_length },
    .{ .name = "lambda", .ptr = &do_lambda },
    .{ .name = "funcall", .ptr = &do_funcall },
    .{ .name = "concatenate", .ptr = &do_concatenate }
};

const SyntaxError = error{};

const RuntimeError = error{};

const ParseIntError = std.fmt.ParseIntError;

const LispError = {
    RuntimeError,
    SyntaxError,
    OutOfMemory,
    EndOf
}




