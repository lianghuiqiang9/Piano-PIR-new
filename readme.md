



go run client_new/client_new.go -ip localhost:50051 -thread 1
2025/04/19 07:27:33 Server address localhost:50051, thread number 1 ignoreOffline false
2025/04/19 07:27:33 1048576 1211212
2025/04/19 07:27:33 DBSize 1048576, DBSeed 1211212, ChunkSize 2048, SetSize 512
2025/04/19 07:27:33 *******************************DB N: 1048576, Entry Size 8 Bytes, DB Size 8 MB
2025/04/19 07:27:33 totalQueryNum 14195
2025/04/19 07:27:33 k = 38
2025/04/19 07:27:33 totalQueryNum 1000
2025/04/19 07:27:33 Start fetching the whole DB
2025/04/19 07:27:33 Every Local Set Size 32 bytes
2025/04/19 07:27:33 Every Local Backup Set Size 16 bytes
2025/04/19 07:27:33 Local Set Num 77824, Local Backup Set Num 41472
2025/04/19 07:27:33 Local Set Size 32 bytes -----------------------
2025/04/19 07:27:33 Local Storage Size 3.640625 MB
2025/04/19 07:27:33 Per query communication cost 4.0078125 kb
2025/04/19 07:27:33 received chunk 0
2025/04/19 07:27:36 Finish Setup Phase, store 77824 local sets, 41472 backup sets/replacement pairs
2025/04/19 07:27:36 *******************************Local Storage Size 3.640625 MB
2025/04/19 07:27:36 Setup Phase took 3634 ms, amortized time 0.26 ms per query
2025/04/19 07:27:36 Setup Phase Comm Cost 8 MB, amortized cost 0.58 KB per query
2025/04/19 07:27:36 Num of local miss elements 0
2025/04/19 07:27:36 Making 0-th query
2025/04/19 07:27:36 Correct value [6937244748061497987] at index 459352
2025/04/19 07:27:37 Finish Online Phase with 1000 queries
2025/04/19 07:27:37 Online Phase took 477 ms, amortized time 0.48 ms
2025/04/19 07:27:37 *******************************Per query upload cost 4 kb --> KB
2025/04/19 07:27:37 *******************************Per query download cost 0.0078 kb --> KB
2025/04/19 07:27:37 End to end amortized time 0.73 ms
2025/04/19 07:27:37 End to end amortized comm cost 4.6 kb
2025/04/19 07:27:37 ---------------breakdown-------------------------
2025/04/19 07:27:37 End to end amortized time 0.73 ms
2025/04/19 07:27:37 Average Online Time 0.48 ms
2025/04/19 07:27:37 Average Network Latency 0.32 ms
2025/04/19 07:27:37 *******************************Average Server Time 0.014 ms
2025/04/19 07:27:37 *******************************Average Client Time 0.15 ms
2025/04/19 07:27:37 Average Find Hint Time 0.12 ms
2025/04/19 07:27:37 -------------------------------------------------



go run client_new/client_new.go -ip localhost:50051 -thread 1
2025/04/19 07:29:09 Server address localhost:50051, thread number 1 ignoreOffline false
2025/04/19 07:29:09 268435456 1211212
2025/04/19 07:29:09 DBSize 268435456, DBSeed 1211212, ChunkSize 32768, SetSize 8192
2025/04/19 07:29:09 *******************************DB N: 268435456, Entry Size 8 Bytes, DB Size 2048 MB
2025/04/19 07:29:09 totalQueryNum 317982
2025/04/19 07:29:09 k = 42
2025/04/19 07:29:09 totalQueryNum 1000
2025/04/19 07:29:09 Start fetching the whole DB
2025/04/19 07:29:09 Every Local Set Size 32 bytes
2025/04/19 07:29:09 Every Local Backup Set Size 16 bytes
2025/04/19 07:29:09 Local Set Num 1376256, Local Backup Set Num 933888
2025/04/19 07:29:09 Local Set Size 32 bytes -----------------------
2025/04/19 07:29:09 Local Storage Size 70.5 MB
2025/04/19 07:29:09 Per query communication cost 64.0078125 kb
2025/04/19 07:29:09 received chunk 0
2025/04/19 07:30:59 received chunk 1000
2025/04/19 07:32:45 received chunk 2000
2025/04/19 07:34:30 received chunk 3000
2025/04/19 07:36:25 received chunk 4000
2025/04/19 07:38:24 received chunk 5000
2025/04/19 07:40:30 received chunk 6000
2025/04/19 07:42:15 received chunk 7000
2025/04/19 07:44:00 received chunk 8000
2025/04/19 07:44:20 Finish Setup Phase, store 1376256 local sets, 933888 backup sets/replacement pairs
2025/04/19 07:44:20 *******************************Local Storage Size 70.5 MB
2025/04/19 07:44:20 Setup Phase took 910613 ms, amortized time 2.9 ms per query
2025/04/19 07:44:20 Setup Phase Comm Cost 2048 MB, amortized cost 6.6 KB per query
2025/04/19 07:44:20 Num of local miss elements 0
2025/04/19 07:44:20 Making 0-th query
2025/04/19 07:44:21 Correct value [13017278094231248772] at index 132348128
2025/04/19 07:46:14 Finish Online Phase with 1000 queries
2025/04/19 07:46:14 Online Phase took 114260 ms, amortized time 1.1e+02 ms
2025/04/19 07:46:14 *******************************Per query upload cost 64 kb --> KB
2025/04/19 07:46:14 *******************************Per query download cost 0.0078 kb --> KB
2025/04/19 07:46:14 End to end amortized time 1.2e+02 ms
2025/04/19 07:46:14 End to end amortized comm cost 71 kb
2025/04/19 07:46:14 ---------------breakdown-------------------------
2025/04/19 07:46:14 End to end amortized time 1.2e+02 ms
2025/04/19 07:46:14 Average Online Time 1.1e+02 ms
2025/04/19 07:46:14 Average Network Latency 0.76 ms
2025/04/19 07:46:14 *******************************Average Server Time 1.1e+02 ms
2025/04/19 07:46:14 *******************************Average Client Time 1.9 ms
2025/04/19 07:46:14 Average Find Hint Time 1.3 ms