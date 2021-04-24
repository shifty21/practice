# Maintain value and index in stack
# conditions 
# 1. if len(Stack) == 0 apppend -1 in result
# 2. if len(stack) >0 and stack[-1][0] > arr[i] result append stack[-1][1]
# 3. if len(Stack) <= 0 and arr[i] >= stack[-1][0] pop and check end conditions
def calculateSpan(a,n):
    stack = []
    result = []
    i = 0
    while i < n:
        if len(stack)==0:
            result.append(-1)
        elif len(stack) >0 and a[i] < stack[-1][0]:
            result.append(stack[-1][1])
        elif len(stack) >0 and a[i] >= stack[-1][0]:
            while len(stack)>0 and a[i] >= stack[-1][0]:
                del stack[-1]
            if len(stack) == 0:
                result.append(-1)
            else:
                result.append(stack[-1][1])
        stack.append((a[i],i))
        i = i +1
    for i in range(len(result)):
        result[i] = i - result[i]
    return result
