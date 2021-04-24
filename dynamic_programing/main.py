#User function Template for python3

matrix = [[-1]*1002 for i in range(1002)]

class Solution:
    
    #Function to return max value that can be put in knapsack of capacity W.
    def knapSack(self,W, wt, val, n):
        if (n==0 or W==0 ):
            return 0
            
        if matrix[n][W] !=-1:
            return matrix[n][W]
        
        if wt[n-1] <= W:
            matrix[n][W]=max(val[n-1] + self.knapSack(W-wt[n-1],wt,val,n-1),
            self.knapSack(W,wt,val,n-1))
            return matrix[n][W]
        
        matrix[n][W]=self.knapSack(W,wt,val,n-1)    
        return matrix[n][W]

    def knapSacktopdown(self,W, wt, val, n):
        matrix = [[0]*(W+1) for i in range(n+1)]
        for i in range(n+1):
            for w in range (W+1):
                if (i==0 or w==0 ):
                    matrix[i][w] = 0
                elif wt[i-1] <= w:
                    matrix[i][w]=max(val[i-1] + matrix[i-1][w-wt[i-1]], matrix[i-1][w])
                else:
                    matrix[i][w]= matrix[i-1][w]
        return matrix[n][w]

    def equalPartition(self, N, arr):
        total_sum = 0
        for x in arr :
            total_sum = total_sum + x
        if total_sum%2!=0:
            return False
        
        return self.subsetsum(arr,total_sum/2,len(arr))
    
    
    def subsetsum(self,arr,s, n):
        if s == 0:
            return True
        if n == 0:
            return False
        
        if arr[n-1] > s:
            return self.subsetsum(arr,s,n-1)
        
        return self.subsetsum(arr,s-arr[n-1],n-1) or self.subsetsum(arr,s,n-1)
        
    def subsetsumtopdown(self,arr,s, n):
        matrix = [[False]*(s+1) for i in range(n+1)]
        for i in range(0,n+1):
            matrix[i][0] = True
        
        for i in range(1,n+1):
            for j in range(1,s+1):
                if arr[i-1] > j:
                    matrix[i][j] = False
                    continue
                matrix[i][j] = matrix[i-1][j-arr[i-1]] or matrix[i-1][j]
        return matrix[n][s] 
    
    def countsubsetsum(self, arr, s,n):
        if s == 0:
            return 1
        if n == 0:
            return 0
        
        if s > arr[n-1]:
            return self.countsubsetsum(arr,s,n-1)
        
        return self.countsubsetsum(arr,s-arr[n-1],n-1) + self.countsubsetsum(arr,s,n-1)

    def countsubsetsumtopdown(self, arr, s,n):
        mat = [[0]*(s+1) for i in range(n+1)]
        for i in range(n+1):
            mat[i][0] = 1

        for i in range(1,n+1):
            for j in range (1, s+1):
                if j > arr[n-1]:
                    mat[i][j] = mat[i-1][j]
                else:
                    mat[i][j] = mat[i-1][j] + mat[i-1][j-arr[i-1]]
        
        return mat[n][s]
    def internalminsubset(self,arr,n,min, out):
        if n == 0 :
            return sum(out)
        out1 = self.internalminsubset(self,arr,n-1,min,out.append(arr[n-1]))
        out2 = self.internalminsubset(self,arr,n-1,min,out)
        if out1 > out2 :
            diff = out1 - out2
        else :
            diff = out2 - out1
        if min > diff:
            return diff
        return min

    def minsubsetsum(self,arr,n):
        min = 0

        return self.internalminsubset(arr,n,min, [])




#{ 
#  Driver Code Starts
#Initial Template for Python 3
import atexit
import io
import sys

# Contributed by : Nagendra Jha

if __name__ == '__main__':
    test_cases = int(input())
    for cases in range(test_cases):
        n = int(input())
        W = int(input())
        val = list(map(int,input().strip().split()))
        wt = list(map(int,input().strip().split()))
        ob=Solution()
        print(ob.knapSack(W,wt,val,n))
# } Driver Code Ends