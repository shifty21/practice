#code
def Premute(S):
    result = []
    if len(S) <=1 :
        result.append(S)
        return
    output = ""
    solve(S,output,result)
    return result
    
def solve(S, output, result):
    if len(S)==0:
        result.append(output)
        return  
    op1 = output + S[0].upper()
    op2 = output + S[0].lower()
    S = S[1:]
    solve(S,op1,result)
    solve(S,op2,result)
    return
    
if __name__ == "__main__":
    result = Premute("ABC")
    print(result)
    