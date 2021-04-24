
class Solution:
	def NBitBinary(self, N):
        result = []
        self.N= N
        return self.solve(0,0,"",result)
        
    def solve(self,ones,twos,output,result):
        if twos + ones == self.N:
            output
            result.append(output)
            return result
        
        if ones >=twos:
            result = self.solve(ones+1,twos,output+"1",result)
        if twos < ones:
            result = self.solve(ones,twos+1,output+"0",result)
        return result

# https://practice.geeksforgeeks.org/problems/print-n-bit-binary-numbers-having-more-1s-than-0s0252/1