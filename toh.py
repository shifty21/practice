class Solution:
    def toh(self, N, fromm, to, aux):
        steps = 1
        if N == 1 :
            print("move disk %d from rod %d to rod %d" % (N,fromm,to))
            return 1
        steps += self.toh(N-1,fromm,aux,to)
        print("move disk %d from rod %d to rod %d" % (N,fromm,to))
        steps += self.toh(N-1,aux,to,fromm)
        return 

