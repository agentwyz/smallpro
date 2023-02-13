#lang racket

(provide (all-defined-out))

(define-syntax letv
  (syntax-rules ()
    [(_ () body ...)
     ()]))

(define-syntax first-val)

(define-syntax second-val)

(define *dubug* #f)

(define-syntax peek)

(define fatal)

(define fold2)

(define orf)

(define char->string string)

(define read-file)


(define new-progress)

(define hash-put!)

(define hash-get)

(define hash-put2!)

(define hash-get2)


(define predand)

(define predor)

(define set-)

(define string-join)

